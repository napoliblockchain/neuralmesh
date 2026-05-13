package mvp

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

type TaskType string

const (
	TaskOCR       TaskType = "ocr"
	TaskEmbedding TaskType = "embedding"
)

type Node struct {
	ID          string
	CPUCores    int
	RAMGB       int
	GPUVRAMGB   int
	Reputation  float64
	ComputeUnit float64
}

type Task struct {
	ID        string
	Type      TaskType
	Payload   string
	CreatedAt time.Time
}

type Result struct {
	TaskID            string
	NodeID            string
	Output            string
	OutputHash        string
	VerificationScore float64
	Reward            float64
}

type Scheduler struct {
	nodes []Node
}

func NewScheduler(nodes []Node) *Scheduler {
	cp := append([]Node(nil), nodes...)
	sort.Slice(cp, func(i, j int) bool {
		return cp[i].ComputeUnit > cp[j].ComputeUnit
	})
	return &Scheduler{nodes: cp}
}

func (s *Scheduler) Assign(task Task) (Node, error) {
	if len(s.nodes) == 0 {
		return Node{}, errors.New("no nodes available")
	}

	for _, node := range s.nodes {
		if supports(node, task.Type) {
			return node, nil
		}
	}

	return Node{}, fmt.Errorf("no node supports task type %q", task.Type)
}

func supports(node Node, taskType TaskType) bool {
	switch taskType {
	case TaskOCR:
		return node.RAMGB >= 8
	case TaskEmbedding:
		return node.RAMGB >= 8
	default:
		return false
	}
}

func RunTask(task Task, node Node) (Result, error) {
	var output string

	switch task.Type {
	case TaskOCR:
		output = DemoOCR(task.Payload)
	case TaskEmbedding:
		output = DemoEmbedding(task.Payload)
	default:
		return Result{}, fmt.Errorf("unsupported task type %q", task.Type)
	}

	hash := sha256.Sum256([]byte(task.ID + ":" + node.ID + ":" + output))

	return Result{
		TaskID:     task.ID,
		NodeID:     node.ID,
		Output:     output,
		OutputHash: hex.EncodeToString(hash[:]),
	}, nil
}

func DemoOCR(input string) string {
	clean := strings.Join(strings.Fields(input), " ")
	return strings.ToUpper(clean)
}

func DemoEmbedding(input string) string {
	tokens := strings.Fields(strings.ToLower(input))
	vector := make([]float64, 8)
	for _, token := range tokens {
		sum := sha256.Sum256([]byte(token))
		for i := range vector {
			vector[i] += float64(sum[i]) / 255.0
		}
	}

	if len(tokens) > 0 {
		for i := range vector {
			vector[i] = math.Round((vector[i]/float64(len(tokens)))*1000) / 1000
		}
	}

	parts := make([]string, len(vector))
	for i, value := range vector {
		parts[i] = fmt.Sprintf("%.3f", value)
	}

	return "[" + strings.Join(parts, ",") + "]"
}

func Verify(primary, redundant Result) float64 {
	if primary.OutputHash == redundant.OutputHash {
		return 1
	}
	if primary.Output == redundant.Output {
		return 0.95
	}
	return 0.25
}

func SimulateReward(baseReward, verificationScore, computeUnits, totalWorkerWeight, reputation float64) float64 {
	if totalWorkerWeight <= 0 {
		return 0
	}

	reputationFactor := math.Min(1.5, reputation/500)
	workerWeight := computeUnits * reputationFactor
	return baseReward * verificationScore * (workerWeight / totalWorkerWeight)
}

func DemoRun() ([]Result, error) {
	nodes := []Node{
		{ID: "node-a", CPUCores: 8, RAMGB: 16, GPUVRAMGB: 8, Reputation: 650, ComputeUnit: 8},
		{ID: "node-b", CPUCores: 4, RAMGB: 8, GPUVRAMGB: 0, Reputation: 500, ComputeUnit: 4},
	}

	tasks := []Task{
		{ID: "task-ocr-001", Type: TaskOCR, Payload: "neuralmesh useful computation", CreatedAt: time.Now()},
		{ID: "task-emb-001", Type: TaskEmbedding, Payload: "distributed ai inference", CreatedAt: time.Now()},
	}

	scheduler := NewScheduler(nodes)
	results := make([]Result, 0, len(tasks))
	totalWorkerWeight := 0.0
	for _, node := range nodes {
		totalWorkerWeight += node.ComputeUnit * math.Min(1.5, node.Reputation/500)
	}

	for _, task := range tasks {
		node, err := scheduler.Assign(task)
		if err != nil {
			return nil, err
		}

		primary, err := RunTask(task, node)
		if err != nil {
			return nil, err
		}
		redundant, err := RunTask(task, node)
		if err != nil {
			return nil, err
		}

		primary.VerificationScore = Verify(primary, redundant)
		primary.Reward = SimulateReward(10, primary.VerificationScore, node.ComputeUnit, totalWorkerWeight, node.Reputation)
		results = append(results, primary)
	}

	return results, nil
}
