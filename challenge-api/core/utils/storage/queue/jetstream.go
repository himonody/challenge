package queue

import (
	"challenge/core/utils/storage"
	json "github.com/json-iterator/go"
	"github.com/nats-io/nats.go"
	"sync"
	"sync/atomic"
	"time"
)

// JetStreamOptions 是 NATS JetStream 队列适配器的配置。
//
// 设计约定：
// - 业务侧的 "stream"（即 message.GetStream()）会被映射为 NATS Subject：SubjectPrefix + stream
// - 每个业务 stream 使用一个 Durable Consumer：DurablePrefix + stream
type JetStreamOptions struct {
	URL              string `yaml:"url" json:"url"`
	Stream           string `yaml:"stream" json:"stream"`
	SubjectPrefix    string `yaml:"subjectPrefix" json:"subjectPrefix"`
	DurablePrefix    string `yaml:"durablePrefix" json:"durablePrefix"`
	AckWaitSeconds   int    `yaml:"ackWaitSeconds" json:"ackWaitSeconds"`
	MaxAckPending    int    `yaml:"maxAckPending" json:"maxAckPending"`
	MaxDeliver       int64  `yaml:"maxDeliver" json:"maxDeliver"`
	ConnectTimeoutMs int    `yaml:"connectTimeoutMs" json:"connectTimeoutMs"`
	ReconnectWaitMs  int    `yaml:"reconnectWaitMs" json:"reconnectWaitMs"`
	MaxReconnects    int    `yaml:"maxReconnects" json:"maxReconnects"`
}

// withDefaults 返回 JetStreamOptions 的默认值。
func (e *JetStreamOptions) withDefaults() *JetStreamOptions {
	if e == nil {
		return &JetStreamOptions{}
	}
	o := *e
	if o.URL == "" {
		o.URL = nats.DefaultURL
	}
	if o.Stream == "" {
		o.Stream = "CHALLENGE_QUEUE"
	}
	if o.SubjectPrefix == "" {
		o.SubjectPrefix = "queue."
	}
	if o.DurablePrefix == "" {
		o.DurablePrefix = "durable_"
	}
	if o.AckWaitSeconds <= 0 {
		o.AckWaitSeconds = 60
	}
	if o.MaxAckPending <= 0 {
		o.MaxAckPending = 1024
	}
	if o.MaxDeliver <= 0 {
		o.MaxDeliver = 10
	}
	if o.ConnectTimeoutMs <= 0 {
		o.ConnectTimeoutMs = 2000
	}
	if o.ReconnectWaitMs <= 0 {
		o.ReconnectWaitMs = 500
	}
	if o.MaxReconnects == 0 {
		o.MaxReconnects = -1
	}
	return &o
}

// NewJetStream 创建一个新的 JetStream 队列适配器。
func NewJetStream(opts *JetStreamOptions) (*JetStream, error) {
	o := opts.withDefaults()

	// 1) 建立到 NATS Server 的连接（含重连策略）。
	nc, err := nats.Connect(
		o.URL,
		nats.Timeout(time.Duration(o.ConnectTimeoutMs)*time.Millisecond),
		nats.ReconnectWait(time.Duration(o.ReconnectWaitMs)*time.Millisecond),
		nats.MaxReconnects(o.MaxReconnects),
	)
	if err != nil {
		return nil, err
	}

	// 2) 获取 JetStream 上下文。
	js, err := nc.JetStream()
	if err != nil {
		nc.Close()
		return nil, err
	}

	// 3) 确保 Stream 存在。
	//    我们用 Subjects: SubjectPrefix + "*" 来覆盖所有业务 stream。
	subjectWildcard := o.SubjectPrefix + "*"
	_, err = js.StreamInfo(o.Stream)
	if err == nats.ErrStreamNotFound {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     o.Stream,
			Subjects: []string{subjectWildcard},
			Storage:  nats.FileStorage,
		})
	}
	if err != nil {
		nc.Close()
		return nil, err
	}

	q := &JetStream{
		opts:  o,
		nc:    nc,
		js:    js,
		subs:  make(map[string]*nats.Subscription),
		wait:  sync.WaitGroup{},
		mutex: sync.Mutex{},
	}
	return q, nil
}

// JetStream 实现了 storage.AdapterQueue，用 JetStream 来承载消息队列。
//
// Register(name, handler)：订阅 subject=SubjectPrefix+name 的消息并消费。
// Append(message)：向 subject=SubjectPrefix+message.Stream 发布消息。
type JetStream struct {
	opts *JetStreamOptions
	nc   *nats.Conn
	js   nats.JetStreamContext

	subs    map[string]*nats.Subscription
	mutex   sync.Mutex
	wait    sync.WaitGroup
	started uint32
	stopped uint32
}

// String 返回 JetStream 的字符串表示。
func (*JetStream) String() string {
	return "nats_jetstream"
}

// subject 返回给定 stream 的 subject。
func (q *JetStream) subject(stream string) string {
	return q.opts.SubjectPrefix + stream
}

// durable 返回给定 stream 的 durable 名称。
func (q *JetStream) durable(stream string) string {
	return q.opts.DurablePrefix + stream
}

// Append 向给定 stream 发布消息。
func (q *JetStream) Append(message storage.Messager) error {
	// 生产消息：把 values JSON 序列化后 publish 到对应 subject。
	b, err := json.Marshal(message.GetValues())
	if err != nil {
		return err
	}
	_, err = q.js.Publish(q.subject(message.GetStream()), b)
	return err
}

// Register 注册一个新的消费者。
func (q *JetStream) Register(name string, f storage.ConsumerFunc) {
	subject := q.subject(name)
	durable := q.durable(name)

	// 保证同一个 name 只注册一次
	q.mutex.Lock()
	if _, ok := q.subs[name]; ok {
		q.mutex.Unlock()
		return
	}
	q.mutex.Unlock()

	// 1) 确保 Durable Consumer 存在（不存在则创建）。
	//    注意：如果 durable 已存在但属于 pull consumer，那么用 Subscribe 绑定会报错。
	//    这里检测到 pull consumer 后会删除并重建为 push consumer。
	info, err := q.js.ConsumerInfo(q.opts.Stream, durable)
	if err != nil && err != nats.ErrConsumerNotFound {
		panic(err)
	}
	if err == nil {
		// DeliverSubject 为空通常意味着 pull consumer
		if info != nil && info.Config.DeliverSubject == "" {
			if err := q.js.DeleteConsumer(q.opts.Stream, durable); err != nil {
				panic(err)
			}
			info = nil
			err = nats.ErrConsumerNotFound
		}
	}
	if err == nats.ErrConsumerNotFound {
		_, err = q.js.AddConsumer(q.opts.Stream, &nats.ConsumerConfig{
			Durable: durable,
			// DeliverSubject 非空意味着这是 push consumer，后续才能用 Subscribe 进行绑定。
			DeliverSubject: nats.NewInbox(),
			FilterSubject:  subject,
			// AckExplicit：业务 handler 返回 nil 才 Ack；否则 Nak 触发重投递。
			AckPolicy:     nats.AckExplicitPolicy,
			AckWait:       time.Duration(q.opts.AckWaitSeconds) * time.Second,
			MaxAckPending: q.opts.MaxAckPending,
			MaxDeliver:    int(q.opts.MaxDeliver),
			DeliverPolicy: nats.DeliverAllPolicy,
			ReplayPolicy:  nats.ReplayInstantPolicy,
		})
		if err != nil {
			panic(err)
		}
	}

	// 2) 订阅 subject 并处理消息。
	//    - 反序列化失败：Term（终止，不再重投递）
	//    - handler 返回 error：Nak（可重投递）
	//    - handler 成功：Ack
	sub, err := q.js.Subscribe(subject, func(msg *nats.Msg) {
		data := make(map[string]interface{})
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			_ = msg.Term()
			return
		}
		m := new(Message)
		m.SetStream(name)
		m.SetValues(data)
		if err := f(m); err != nil {
			_ = msg.Nak()
			return
		}
		_ = msg.Ack()
	},
		nats.Durable(durable),
		nats.ManualAck(),
		nats.AckExplicit(),
	)
	if err != nil {
		panic(err)
	}

	q.mutex.Lock()
	q.subs[name] = sub
	q.mutex.Unlock()
}

func (q *JetStream) Run() {
	// 与现有队列实现保持一致：Run() 阻塞，直到 Shutdown()。
	if atomic.CompareAndSwapUint32(&q.started, 0, 1) {
		q.wait.Add(1)
	}
	q.wait.Wait()
}

func (q *JetStream) Shutdown() {
	if !atomic.CompareAndSwapUint32(&q.stopped, 0, 1) {
		return
	}

	// 释放订阅与连接。
	q.mutex.Lock()
	for _, sub := range q.subs {
		_ = sub.Unsubscribe()
	}
	q.subs = map[string]*nats.Subscription{}
	q.mutex.Unlock()

	if q.nc != nil {
		q.nc.Drain()
		q.nc.Close()
	}

	if atomic.LoadUint32(&q.started) == 1 {
		q.wait.Done()
	}
}
