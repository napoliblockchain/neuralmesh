package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jambtc/neuralmesh/internal/mvp"
)

func main() {
	payload := strings.Join(os.Args[1:], " ")
	if payload == "" {
		payload = "neuralmesh ocr demo"
	}

	result, err := mvp.RunTask(
		mvp.Task{ID: "task-runner-demo", Type: mvp.TaskOCR, Payload: payload, CreatedAt: time.Now()},
		mvp.Node{ID: "local-node", RAMGB: 16, Reputation: 500, ComputeUnit: 8},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "task failed: %v\n", err)
		os.Exit(1)
	}

	encoded, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(encoded))
}
