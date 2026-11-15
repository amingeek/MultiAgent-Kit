# Agent Kit Specification

- Kit Version: 5.3 (Production-Ready Multi-Agent Edition)
- Architecture: Supervisor-Worker with Strictly Typed Contract Handoff & Hierarchical Milestone DAG
- Paradigms: Contract-First, Test-Driven Development (TDD), Context Isolation, Zero-Hallucination, Anti-Looping, Granular Milestone Decomposition
- Supported Environments: Universal (Go, Python, TypeScript/Node, Rust, C#, PHP, Docker, etc.)
- Enhancements in 5.3:
  - **Hierarchical Epic-to-Granular-Milestones Architecture:** 1 Master acts as an architectural blueprint decomposing into 1 to 10+ granular, isolated Milestones (`NN-01` to `NN-10+`).
  - **Deep Technical Milestone Specifications:** Exhaustive contracts, type signatures, invariants, negative scenarios, and edge-case resilience gates in `_MILESTONE-TEMPLATE.md`.
  - **Enhanced DAG Milestone Generation:** Upgraded `NEW-FEATURE.md` and `_MASTER-TEMPLATE.md` to formulate multi-milestone execution graphs.
- Enhancements in 5.2:
  - Automated state validation tool (`agent/tools/validate_state.py`).
  - Fully articulated, battle-tested on-demand and always skills.
  - Complete domain feature and milestone decomposition workflows (`NEW-FEATURE.md`, `_MILESTONE-TEMPLATE.md`).
  - High-precision project reconnaissance engine (`00-info.md`).
  - Subagent-aware orchestration engine with clean context segregation.
