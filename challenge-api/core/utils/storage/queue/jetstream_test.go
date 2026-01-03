package queue

import (
	"challenge/core/utils/storage"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"gopkg.in/yaml.v3"
)

func findProjectRootFrom(wd string) (string, error) {
	d := wd
	for {
		if _, err := os.Stat(filepath.Join(d, "go.mod")); err == nil {
			return d, nil
		}
		parent := filepath.Dir(d)
		if parent == d {
			return "", os.ErrNotExist
		}
		d = parent
	}
}

func loadJetStreamOptionsFromSettingsYml(t *testing.T) *JetStreamOptions {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd error: %v", err)
	}
	root, err := findProjectRootFrom(wd)
	if err != nil {
		t.Fatalf("find project root error: %v", err)
	}
	settingsPath := filepath.Join(root, "config", "settings.yml")
	b, err := os.ReadFile(settingsPath)
	if err != nil {
		t.Fatalf("read %s error: %v", settingsPath, err)
	}

	var cfg struct {
		Settings struct {
			Queue struct {
				NATS *JetStreamOptions `yaml:"nats"`
			} `yaml:"queue"`
		} `yaml:"settings"`
	}
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		t.Fatalf("yaml unmarshal %s error: %v", settingsPath, err)
	}
	if cfg.Settings.Queue.NATS == nil {
		t.Fatalf("settings.queue.nats is empty in %s", settingsPath)
	}
	return cfg.Settings.Queue.NATS
}

func TestJetStream_AppendAndConsume(t *testing.T) {
	// 这是一个集成测试：需要本机/环境里已经启动 NATS JetStream。
	//
	// 使用方法（示例）：
	//    go test ./core/utils/storage/queue -run TestJetStream_AppendAndConsume -v
	opts := loadJetStreamOptionsFromSettingsYml(t)

	q, err := NewJetStream(opts)
	if err != nil {
		t.Fatalf("NewJetStream error: %v", err)
	}
	defer q.Shutdown()

	streamName := "ping"

	var (
		wg   sync.WaitGroup
		once sync.Once
	)
	wg.Add(1)

	q.Register(streamName, func(m storage.Messager) error {
		t.Logf("[consume] stream=%s values=%#v", m.GetStream(), m.GetValues())
		defer once.Do(func() { wg.Done() })
		v := m.GetValues()["hello"]
		s, _ := v.(string)
		if s != "world" {
			t.Fatalf("unexpected payload: %#v", m.GetValues())
		}
		return nil
	})

	msg := new(Message)
	msg.SetStream(streamName)
	msg.SetValues(map[string]interface{}{"hello": "world"})
	t.Logf("[produce] stream=%s values=%#v", msg.GetStream(), msg.GetValues())
	err = q.Append(msg)
	if err != nil {
		t.Fatalf("Append error: %v", err)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return
	case <-time.After(10 * time.Second):
		t.Fatalf("timeout waiting for jetstream message")
	}
}
