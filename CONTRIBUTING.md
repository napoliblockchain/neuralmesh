# Contributing to NeuralMesh

## Before You Start

1. Read the [README](README.md) and [whitepaper](docs/01_whitepaper.md).
2. Check [docs/NIP-INDEX.md](docs/NIP-INDEX.md) — find an open NIP to work on.
3. Check the [feasibility review](docs/FEASIBILITY-REVIEW.md) for known limitations.

## Branch Strategy

```text
main        protected — merges via PR only, requires 1 review
develop     integration branch — merges via PR only
nip/NIP-XXX-short-name   feature branches
```text

Always branch from `develop`, not `main`.

## Workflow

```bash
git checkout develop
git pull origin develop
git checkout -b nip/NIP-030-onnx-runtime
# ... work ...
git push -u origin nip/NIP-030-onnx-runtime
# open PR into develop
```text

## Commit Messages

```text
NIP-030: implement ONNX Runtime CPU base image

Add Dockerfile for onnx-cpu base image and Go package
for model loading and inference execution.
```text

Format: `NIP-XXX: short description` (imperative, lowercase after colon).

## NIP Completion

When your PR implements a NIP:

1. Set NIP status to `✅ Completed`.
2. Add `## Verification` section to the NIP file.
3. Update `docs/NIP-INDEX.md`.
4. Reference the NIP in your PR description.

## Code Standards

- Go 1.22+
- `go fmt` and `go vet` must pass
- `golangci-lint run` must pass
- Test coverage ≥ 80% for new packages
- No global mutable state
- SQL via `sqlc` only (no ORM)
- Docker containers: `--network none` by default

## Proposing a New NIP

Open a GitHub Issue using the **NIP Proposal** template.
If accepted, create the NIP file in the correct `docs/` folder and update `docs/NIP-INDEX.md`.

## Versioning

We use [Semantic Versioning](https://semver.org/):

| Version | Meaning |
|---------|---------|
| 0.1.x | Phase 0 — Foundation |
| 0.2.x | Phase 1 — MVP |
| 0.3.x | Phase 2 — Community |
| 1.0.0 | Phase 3 — Economic Layer (token live) |
| 2.0.0 | Phase 4 — Full Decentralization |

## License

By contributing, you agree your contributions are licensed under [Apache 2.0](LICENSE).
