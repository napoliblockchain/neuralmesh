package mvp

import "testing"

func TestSchedulerAssignsSupportedNode(t *testing.T) {
	scheduler := NewScheduler([]Node{{ID: "node-a", RAMGB: 16, ComputeUnit: 2}})

	node, err := scheduler.Assign(Task{ID: "task-1", Type: TaskOCR, Payload: "hello"})
	if err != nil {
		t.Fatalf("Assign returned error: %v", err)
	}

	if node.ID != "node-a" {
		t.Fatalf("node ID = %q, want node-a", node.ID)
	}
}

func TestDemoEmbeddingDeterministic(t *testing.T) {
	a := DemoEmbedding("distributed ai")
	b := DemoEmbedding("distributed ai")

	if a != b {
		t.Fatalf("embedding not deterministic: %q != %q", a, b)
	}
}

func TestVerifyMatchingHash(t *testing.T) {
	result := Result{Output: "ok", OutputHash: "same"}

	if score := Verify(result, result); score != 1 {
		t.Fatalf("score = %v, want 1", score)
	}
}

func TestSimulateReward(t *testing.T) {
	reward := SimulateReward(10, 1, 5, 10, 500)

	if reward != 5 {
		t.Fatalf("reward = %v, want 5", reward)
	}
}
