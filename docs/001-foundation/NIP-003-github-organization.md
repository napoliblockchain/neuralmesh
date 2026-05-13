# NIP-003 — GitHub Organization Setup

- Status: 📋 Open
- Area: Foundation / DevOps
- Phase: 0
- Priority: High
- Date: 2026-05-11

## Context

No GitHub organization exists. The project needs a public presence for credibility and open-source collaboration.

## Proposal

Create and configure the NeuralMesh GitHub organization:

### Sub-tasks

- [x] NIP-003a: GitHub organization `napoliblockchain` — already exists
- [ ] NIP-003b: Create repository `neuralmesh` under `napoliblockchain`
- [ ] NIP-003c: Create repository `neuralmesh-client` (desktop client) — or monorepo
- [ ] NIP-003d: Create repository `neuralmesh-docs` (documentation site) — or monorepo
- [ ] NIP-003e: Configure branch protection rules on `main`
- [x] NIP-003f: `CONTRIBUTING.md` and `CODE_OF_CONDUCT.md` — already present
- [x] NIP-003g: Issue templates (bug, feature, NIP proposal) — already present
- [x] NIP-003h: GitHub Actions CI skeleton — already present (ci.yml, docs.yml, release.yml)

## Acceptance Criteria

- Organization exists and is public.
- At least `neuralmesh-core` and `neuralmesh-docs` repos created.
- Branch protection on `main` requires PR + 1 review.
- `CONTRIBUTING.md` and `CODE_OF_CONDUCT.md` present in `neuralmesh-core`.
- Issue templates operational.

## Impact

- External: GitHub organization and public repositories
- New: `CONTRIBUTING.md`, `CODE_OF_CONDUCT.md`, `.github/ISSUE_TEMPLATE/`

## Implementation Notes

- Organization: `napoliblockchain` (existing, confirmed 2026-05-13)
- Main repo: `napoliblockchain/neuralmesh`
- License: GPL-3.0 (as declared in README)
- Use GitHub Actions for CI (free tier sufficient for Phase 0).
