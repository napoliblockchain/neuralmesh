package main

import (
	"fmt"

	"github.com/jambtc/neuralmesh/internal/mvp"
)

func main() {
	reward := mvp.SimulateReward(10, 0.95, 8, 12, 650)
	fmt.Printf("simulated_reward_nmc=%.4f\n", reward)
}
