package queue

import (
	"challenge/core/utils/storage"
	redisqueue2 "challenge/core/utils/storage/queue/redisqueue"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"testing"
	"time"
)

func TestRedis_Append(t *testing.T) {
	type fields struct {
		ConnectOption   *redis.Options
		ConsumerOptions *redisqueue2.ConsumerOptions
		ProducerOptions *redisqueue2.ProducerOptions
		client          *redis.Client
		consumer        *redisqueue2.Consumer
		producer        *redisqueue2.Producer
	}
	type args struct {
		name    string
		message storage.Messager
	}
	client := redis.NewClient(&redis.Options{})
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"test01",
			fields{
				ConnectOption: &redis.Options{},
				ConsumerOptions: &redisqueue2.ConsumerOptions{
					VisibilityTimeout: 60 * time.Second,
					BlockingTimeout:   5 * time.Second,
					ReclaimInterval:   1 * time.Second,
					BufferSize:        100,
					Concurrency:       10,
					RedisClient:       client,
				},
				ProducerOptions: &redisqueue2.ProducerOptions{
					StreamMaxLength:      100,
					ApproximateMaxLength: false,
					RedisClient:          client,
				},
			},
			args{
				name: "test",
				message: &Message{redisqueue2.Message{
					Stream: "test",
					Values: map[string]interface{}{
						"key": "value",
					},
				}},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if r, err := NewRedis(tt.fields.ProducerOptions, tt.fields.ConsumerOptions); err != nil {
				t.Errorf("SetQueue() error = %v", err)
			} else {
				if err := r.Append(tt.args.message); (err != nil) != tt.wantErr {
					t.Errorf("SetQueue() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

// TestRedis_SendAndConsume verifies a message can be produced and consumed via Redis streams.
func TestRedis_SendAndConsume(t *testing.T) {
	addr := os.Getenv("TEST_REDIS_ADDR")
	if addr == "" {
		addr = "127.0.0.1:6379"
	}
	client := redis.NewClient(&redis.Options{Addr: addr})
	if err := client.Ping(context.Background()).Err(); err != nil {
		t.Skipf("skip: redis not available at %s: %v", addr, err)
	}

	stream := fmt.Sprintf("test_stream_%d", time.Now().UnixNano())
	producerOptions := &redisqueue2.ProducerOptions{
		StreamMaxLength:      100,
		ApproximateMaxLength: true,
		RedisClient:          client,
	}
	consumerOptions := &redisqueue2.ConsumerOptions{
		VisibilityTimeout: 30 * time.Second,
		BlockingTimeout:   2 * time.Second,
		ReclaimInterval:   1 * time.Second,
		BufferSize:        10,
		Concurrency:       1,
		RedisClient:       client,
	}

	q, err := NewRedis(producerOptions, consumerOptions)
	if err != nil {
		t.Fatalf("NewRedis error: %v", err)
	}
	defer func() {
		q.Shutdown()
		_ = client.Del(context.Background(), stream).Err()
	}()

	got := make(chan map[string]interface{}, 1)
	q.Register(stream, func(message storage.Messager) error {
		fmt.Printf("[consume] stream=%s values=%v\n", message.GetStream(), message.GetValues())
		got <- message.GetValues()
		return nil
	})

	// Start consumer
	go q.Run()

	msg := &Message{redisqueue2.Message{
		Stream: stream,
		Values: map[string]interface{}{"hello": "world"},
	}}
	fmt.Printf("[produce] stream=%s values=%v\n", msg.GetStream(), msg.GetValues())
	if err := q.Append(msg); err != nil {
		t.Fatalf("Append error: %v", err)
	}

	select {
	case vals := <-got:
		if vals["hello"] != "world" {
			t.Fatalf("unexpected payload: %+v", vals)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("timed out waiting for message consumption")
	}
}

func TestRedis_Register(t *testing.T) {
	type fields struct {
		ConnectOption   *redis.Options
		ConsumerOptions *redisqueue2.ConsumerOptions
		ProducerOptions *redisqueue2.ProducerOptions
		client          *redis.Client
		consumer        *redisqueue2.Consumer
		producer        *redisqueue2.Producer
	}
	type args struct {
		name string
		f    storage.ConsumerFunc
	}
	client := redis.NewClient(&redis.Options{})
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"test01",
			fields{
				ConnectOption: &redis.Options{},
				ConsumerOptions: &redisqueue2.ConsumerOptions{
					VisibilityTimeout: 60 * time.Second,
					BlockingTimeout:   5 * time.Second,
					ReclaimInterval:   1 * time.Second,
					BufferSize:        100,
					Concurrency:       10,
					RedisClient:       client,
				},
				ProducerOptions: &redisqueue2.ProducerOptions{
					StreamMaxLength:      100,
					ApproximateMaxLength: true,
					RedisClient:          client,
				},
			},
			args{
				name: "login_log_queue",
				f: func(message storage.Messager) error {
					fmt.Println("ok")
					fmt.Println(message.GetValues())
					return nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if r, err := NewRedis(tt.fields.ProducerOptions, tt.fields.ConsumerOptions); err != nil {
				t.Errorf("SetQueue() error = %v", err)
			} else {
				r.Register(tt.args.name, tt.args.f)
				r.Run()
			}
		})
	}
}
