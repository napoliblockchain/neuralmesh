# Feasibility and Known Limits

## Public Position

NeuralMesh is an engineering project for verifiable distributed AI inference.
The first milestone is not a fully trustless AI blockchain.

## Known Limits

### zkML

Zero-knowledge machine learning proofs are not practical for MVP inference.
They remain future research until proof generation overhead becomes usable.

### Distributed Training

Distributed model training is out of scope for MVP.
The first target is inference: OCR, embeddings, LLM calls, and similar workloads.

### Verification

Proof of Useful Computation is probabilistic in MVP.
Redundant execution, hidden tasks, nonces, and similarity checks raise attacker cost but do not provide mathematical certainty.

### Token

Real token launch comes only after:

- working MVP,
- demonstrated compute utility,
- public benchmarks,
- security review,
- legal review.

MVP rewards are simulated.

### Model Sharding

Internet-wide sharding is suitable only for batch/async tasks in early phases.
Real-time chat should run on one capable node until latency data proves another path.

### Consumer Hardware

Hardware participation depends on task type.
OCR and embeddings can run on modest machines.
Useful LLM and image workloads need modern RAM/VRAM.

## MVP Promise

MVP should prove one narrow claim:

> A consumer node can execute a useful AI task, produce verifiable output metadata, and receive a simulated reward based on probabilistic verification.
