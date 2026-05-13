# NIP-024 — Database Setup (PostgreSQL + Redis)

- Status: 📋 Open
- Area: MVP Core / Data
- Phase: 1
- Priority: High
- Date: 2026-05-11

## Context

Persistent and ephemeral data stores are required for:

- Node state and identity (PostgreSQL)
- Task queue and caching (Redis)
- Reward history and reputation (PostgreSQL)

## Proposal

Set up and schema both data stores:

### PostgreSQL — Persistent Store

Tables required:

- `nodes` — node identity, public key, hardware profile
- `tasks` — task history, status, payload hash, result hash
- `rewards` — reward ledger per task per node
- `reputation_log` — reputation change events

### Redis — Ephemeral Store

Keys required:

- `task_queue` — sorted set (priority queue)
- `node_status:{id}` — hash (current state, last seen)
- `peer_list` — set of known peers
- `rate_limit:{id}` — counter for rate limiting

### Sub-tasks

- [ ] NIP-024a: Write PostgreSQL schema SQL (all tables, indexes, constraints)
- [ ] NIP-024b: Write Redis key schema documentation
- [ ] NIP-024c: Implement migration tooling for PostgreSQL schema
- [ ] NIP-024d: Implement data access layer (repository pattern)
- [ ] NIP-024e: Configure connection pooling (pgx pool for PostgreSQL)
- [ ] NIP-024f: Configure Redis client with retry and timeout
- [ ] NIP-024g: Write backup and restore procedure for PostgreSQL

## Acceptance Criteria

- PostgreSQL schema deploys cleanly with no manual intervention.
- All tables have primary keys, foreign keys, and appropriate indexes.
- Redis key schema is documented with TTL values.
- Data access layer abstracts raw SQL from business logic.
- Connection pool is configurable via environment variables.

## Impact

- New: `neuralmesh-client/internal/store/`
- New: `neuralmesh-client/migrations/`

## Implementation Notes

- PostgreSQL 16+; Redis 7+.
- ORM: avoid — use `sqlc` for type-safe SQL generation.
- Redis client: `go-redis/v9`.
- TTL for task_queue entries: 24h (expired tasks auto-removed).
