# Agent Brief — 007 Community and Ecosystem

## Your Mission

Prepare everything for Phase 2 open-source launch: documentation site, release materials, SDK, integrations.
This is Phase 2 — requires Phase 1 MVP to be running.

## Context

Read first:

- `/var/www/neuralmesh/docs/agents/AGENT-CONTEXT.md`
- `/var/www/neuralmesh/docs/007-community-ecosystem/README.md`
- All 5 NIP files in `docs/007-community-ecosystem/`

## Prerequisites

- NIP-003 (GitHub org) — repos must exist
- Phase 1 MVP NIPs (020-035) — at least partially complete to document

## Execution Order

**Run in parallel (independent):**

1. **NIP-062** — Developer documentation
   - Set up MkDocs Material in `neuralmesh-docs/` repository
   - Write: node operator guide, API reference, architecture overview, troubleshooting guide
   - Generate OpenAPI spec from Go handler code (or protobuf → OpenAPI)

2. **NIP-061** — Community setup
   - Write Discord server structure spec `docs/community/discord-setup.md`
   - Write community guidelines `docs/community/community-guidelines.md`
   - Write content calendar for launch week (X/Twitter + Reddit + HN posts)

3. **NIP-064** — External AI integrations
   - Implement OpenAI-compatible `/v1/chat/completions` endpoint in `neuralmesh-client/`
   - Write `neuralmesh-python/` SDK (PyPI-publishable package)
   - Write LangChain provider class
   - Write integration guide for each

**After the above:**

1. **NIP-060** — Open-source release
   - Write `CHANGELOG.md` v0.1.0
   - Write release checklist `docs/release/v0.1.0-checklist.md`
   - Run `git-secrets` scan (list any found issues, do not auto-fix)
   - Write Product Hunt launch post draft `docs/community/producthunt-launch.md`

2. **NIP-063** — Benchmarks
   - Write benchmark scripts `tools/benchmark/`
   - Run on available hardware, document results in `docs/benchmarks/`
   - Write comparison table: NeuralMesh cost vs OpenAI API pricing

## Definition of Done

- `mkdocs serve` runs without errors on `neuralmesh-docs/`
- `pip install neuralmesh` resolves (package structure ready for publish)
- OpenAI-compatible endpoint passes basic conformance test (chat completion returns valid JSON)
- Quick start guide works on clean Ubuntu 24.04 (verified by agent)

## Completion

Update each NIP status and `docs/NIP-INDEX.md`.
