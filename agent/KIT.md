# Agent Kit Specification

- Kit Version: 6.0 (Production-Ready Go-Native & PRD-Driven Multi-Agent Edition)
- Architecture: Supervisor-Worker with Strictly Typed Contract Handoff & Hierarchical Milestone DAG
- Paradigms: PRD-Driven Engineering, Contract-First, Test-Driven Development (TDD), Strict SRP Role Segregation, Context Isolation, Anti-Looping Circuit Breaker, Granular Milestone Decomposition
- Tooling Architecture: Pure Go-Native Compiled Binary (`agent/tools/mak` & `agent/tools/validate_state`), Zero Python Dependencies
- Supported Target Stacks: Universal (Go, TypeScript/Node, Rust, Python, C#, PHP, Java, Docker, Kubernetes, etc.)
- Enhancements in v6.0:
  - **PRD-Driven End-to-End Workflow:** Product concept, architecture preferences, and feature domains captured in `PRD.md`.
  - **Dynamic Multi-Master Generation:** `BOOTSTRAP.md` automatically ingests `PRD.md` and generates all Domain Masters (`agent/masters/NN-*.md`) with zero limits on quantity.
  - **Deep Technical Milestone Decomposition Engine (`PLAN-MILESTONES.md`):** Decomposes each Master into unlimited granular, production-ready Milestones (`agent/Milestones/NN-MM-*.md`) with exhaustive contracts, types, negative scenarios, and TDD verification matrices.
  - **Autonomous Product Execution Engine (`EXECUTE.md`):** Single-command full lifecycle execution from initial milestone to complete product delivery.
  - **Single-Binary Go-Native CLI (`agent/tools/mak`):** Native compiled Go tool providing subcommands for validation (`mak validate`), deep system audits (`mak check-all`), execution dashboards (`mak summary`), milestone progress metrics (`mak stats`), and template initialization (`mak init-prd`).
  - **Eliminated Python Artifacts:** Complete removal of Python scripts and wrappers in favor of standalone, high-performance Go binaries.
  - **Cleaned Up Workspace Clutter:** Removed `.idea/` configs, duplicate `.keep` files, and outdated migration prompts.
