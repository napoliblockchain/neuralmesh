package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jambtc/neuralmesh/internal/mvp"
)

func main() {
	results, err := mvp.DemoRun()
	if err != nil {
		fmt.Fprintf(os.Stderr, "demo failed: %v\n", err)
		os.Exit(1)
	}

	encoded, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "encode failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(encoded))
}
