package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jambtc/neuralmesh/internal/mvp"
)

func main() {
	scheduler := mvp.NewScheduler([]mvp.Node{
		{ID: "node-a", RAMGB: 16, ComputeUnit: 8},
		{ID: "node-b", RAMGB: 8, ComputeUnit: 4},
	})

	node, err := scheduler.Assign(mvp.Task{ID: "schedule-demo", Type: mvp.TaskEmbedding, Payload: "hello", CreatedAt: time.Now()})
	if err != nil {
		fmt.Fprintf(os.Stderr, "schedule failed: %v\n", err)
		os.Exit(1)
	}

	encoded, _ := json.MarshalIndent(node, "", "  ")
	fmt.Println(string(encoded))
}
