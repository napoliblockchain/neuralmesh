# Agent Brief — 003 MVP Core Infrastructure

## Your Mission

Implement the core Go daemon for NeuralMesh Phase 1 MVP.
You are building the foundational infrastructure that everything else plugs into.

## Context

Read first:

- `/var/www/neuralmesh/docs/agents/AGENT-CONTEXT.md`
- `/var/www/neuralmesh/CLAUDE.md`
- `/var/www/neuralmesh/docs/003-mvp-core/README.md`
- All 6 NIP files in `docs/003-mvp-core/`

## Prerequisites

Before writing code, verify these are available or documented:

- `docs/002-protocol-design/NIP-010-pouc-specification.md` — status: at least `🔄 In Progress`
- `docs/002-protocol-design/NIP-011-task-lifecycle.md` — status: at least `🔄 In Progress`

## Repository Setup

Create `neuralmesh-client/` in the project root with:

```text
neuralmesh-client/
  cmd/
    neuralmesh-node/   ← main entry point
  internal/
    identity/          ← NIP-020 (node keypair, hardware detection)
    network/           ← NIP-021 (libp2p P2P layer)
    scheduler/         ← NIP-022 (task queue + dispatcher)
    sandbox/           ← NIP-023 (Docker execution)
    store/             ← NIP-024 (PostgreSQL + Redis)
    rewards/           ← NIP-025 (simulation engine)
  migrations/          ← SQL migration files
  go.mod
  go.sum
  Makefile
```text

## Execution Order

**Wave 1 (independent — implement simultaneously if multiple agents):**

- NIP-024: Write SQL schema and data access layer (`internal/store/`)
- NIP-023: Write Docker sandbox manager (`internal/sandbox/`)
- NIP-021: Write P2P networking layer (`internal/network/`) using libp2p-go

**Wave 2 (after Wave 1):**

- NIP-022: Write task scheduler (`internal/scheduler/`)
- NIP-020: Wire everything together in daemon (`cmd/neuralmesh-node/`)

**Wave 3:**

- NIP-025: Write reward simulation tool (`tools/simulate/`)

## Key Technical Constraints

- Go 1.22+, no CGO unless strictly necessary
- Docker SDK: `github.com/docker/docker`
- P2P: `github.com/libp2p/go-libp2p`
- DB: `github.com/jackc/pgx/v5` + `github.com/redis/go-redis/v9`
- SQL generation: `sqlc`
- Containers: no network (`--network none`), no host mounts except `/tmp/task-{id}/`

## Definition of Done

- `go build ./...` succeeds with no errors
- `go test ./...` passes with coverage ≥ 80% per package
- `neuralmesh-node --help` outputs usage
- Docker sandbox creates, runs a hello-world container, and cleans up
- P2P node discovers at least one peer when pointed at bootstrap address

## Completion

Update each NIP status and `docs/NIP-INDEX.md`.
