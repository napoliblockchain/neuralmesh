# Agent Brief — 004 AI Compute Layer

## Your Mission

Implement the AI workload execution layer for NeuralMesh.
You are building the Docker-based ONNX inference runners that do the actual "useful computation."

## Context

Read first:

- `/var/www/neuralmesh/docs/agents/AGENT-CONTEXT.md`
- `/var/www/neuralmesh/CLAUDE.md`
- `/var/www/neuralmesh/docs/004-ai-compute/README.md`
- All 6 NIP files in `docs/004-ai-compute/`

## Prerequisites

These must be completed before you start:

- NIP-023 (Docker sandboxing) — ✅ Completed
- NIP-022 (Scheduler) — ✅ Completed

## Work Location

Your work lives in two places:

1. `docker/` — Docker base images
2. `neuralmesh-client/internal/ai/` — Go packages that interact with the containers

```text
docker/
  base-images/
    onnx-cpu/
    onnx-cuda/
    onnx-rocm/
  workers/
    llm/
    image-gen/
    ocr/
    audio/

neuralmesh-client/internal/ai/
  onnx/        ← NIP-030 runtime integration
  sharding/    ← NIP-031 model sharding
  llm/         ← NIP-032 LLM pipeline
  image/       ← NIP-033 image generation
  audio/       ← NIP-034 speech/STT
  ocr/         ← NIP-035 OCR
```text

## Execution Order

**Start with (no deps between them):**

1. **NIP-030** — Build ONNX base Docker images (CPU first). Write Go package `internal/ai/onnx/` for model loading + inference execution.

**Then in parallel:**
2. **NIP-032** — LLM inference (highest priority). Depends on NIP-030.
3. **NIP-031** — Model sharding design. Depends on NIP-030.

**Then:**
4. **NIP-033** — Image generation. Depends on NIP-030.
5. **NIP-034**, **NIP-035** — Audio and OCR. Depend on NIP-030.

## Key Technical Constraints

- Docker images must work offline (weights baked in or downloaded on first build, not per-task)
- ONNX Runtime: `onnxruntime` Go bindings or invoke via subprocess
- Containers: `--network none`, `--memory 8g`, `--cpus 4.0` (configurable defaults)
- All output includes: result bytes, SHA-256 hash, execution timing, intermediate checkpoint hashes
- Models: support ONNX format only. Provide 1 example model per workload type for testing.

## Definition of Done

- ONNX CPU base image builds with `docker build`
- `go test ./internal/ai/...` passes
- LLM: can load a quantized 1B ONNX model and generate 50 tokens
- Execution timing is measured and logged per inference

## Completion

Update each NIP status and `docs/NIP-INDEX.md`.
