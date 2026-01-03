package queue

import "testing"

func TestJetStreamOptions_withDefaults(t *testing.T) {
	o := (&JetStreamOptions{}).withDefaults()
	if o.URL == "" {
		t.Fatalf("expected default URL")
	}
	if o.Stream == "" {
		t.Fatalf("expected default stream")
	}
	if o.SubjectPrefix == "" {
		t.Fatalf("expected default subjectPrefix")
	}
	if o.DurablePrefix == "" {
		t.Fatalf("expected default durablePrefix")
	}
	if o.AckWaitSeconds <= 0 {
		t.Fatalf("expected default ackWaitSeconds")
	}
	if o.MaxAckPending <= 0 {
		t.Fatalf("expected default maxAckPending")
	}
	if o.MaxDeliver <= 0 {
		t.Fatalf("expected default maxDeliver")
	}
	if o.ConnectTimeoutMs <= 0 {
		t.Fatalf("expected default connectTimeoutMs")
	}
	if o.ReconnectWaitMs <= 0 {
		t.Fatalf("expected default reconnectWaitMs")
	}
	if o.MaxReconnects != -1 {
		t.Fatalf("expected default maxReconnects=-1, got %d", o.MaxReconnects)
	}
}

func TestJetStream_subjectAndDurable(t *testing.T) {
	q := &JetStream{opts: (&JetStreamOptions{SubjectPrefix: "queue.", DurablePrefix: "durable_"}).withDefaults()}
	if got := q.subject("ping"); got != "queue.ping" {
		t.Fatalf("unexpected subject: %s", got)
	}
	if got := q.durable("ping"); got != "durable_ping" {
		t.Fatalf("unexpected durable: %s", got)
	}
}
