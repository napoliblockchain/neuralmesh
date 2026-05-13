# NeuralMesh Agent Orchestration

This folder contains everything needed to spawn AI agents to work on NeuralMesh NIPs.

## Files

| File | Purpose |
|------|---------|
| [AGENT-GUIDE.md](AGENT-GUIDE.md) | How to spawn and manage agents |
| [PARALLELISM-MAP.md](PARALLELISM-MAP.md) | Which NIPs can run in parallel — dependency graph |
| [AGENT-CONTEXT.md](AGENT-CONTEXT.md) | Compact project context for agent cold-start |
| [briefs/](briefs/) | Self-contained agent brief per macro-area |
| [codex/](codex/) | Codex-specific parallel agent briefs by discipline |

## Quick Start

To spawn an agent on a specific area, use the corresponding brief in `briefs/`.
Each brief is a self-contained prompt: paste it directly into a Claude Code agent invocation.

## Parallel Execution Strategy

The dependency graph in `PARALLELISM-MAP.md` defines which NIPs are safe to work on simultaneously.
Never assign two agents to NIPs with a direct dependency relationship.

## Agent Naming Convention

Name each agent after its NIP for traceability:

```text
agent-nip-021-p2p
agent-nip-030-onnx
agent-nip-024-database
```text
