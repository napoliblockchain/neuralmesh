package main

import (
	"encoding/json"
	"fmt"

	"github.com/jambtc/neuralmesh/internal/mvp"
)

func main() {
	node := mvp.Node{
		ID:          "local-node",
		CPUCores:    8,
		RAMGB:       16,
		GPUVRAMGB:   8,
		Reputation:  500,
		ComputeUnit: 8,
	}

	encoded, _ := json.MarshalIndent(node, "", "  ")
	fmt.Println(string(encoded))
}
