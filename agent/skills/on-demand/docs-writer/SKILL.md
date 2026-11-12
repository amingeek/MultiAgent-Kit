---
name: docs-writer
type: on-demand
description: Accurate, drift-proof documentation derived exclusively from real code.
---
# Documentation Writer & Synchronization Standard

Documentation must always reflect real, verified code. Never document aspirational features or invent interfaces.

## Documentation Workflow

### 1. Code-First Fact Gathering
- Inspect real source files, route declarations, type definitions, and CLI flags.
- Verify exact signatures, default values, error codes, and configuration options.
- Run `--help` flags or inspect API schemas to confirm actual CLI/API behaviors.

### 2. Updating Architectural & Project Maps
- When new components, contracts, or public endpoints are introduced:
  - Update `agent/info/PROJECT_MAP.md` with new modules, routes, and layout entries.
  - Update `agent/rules/10-contracts-and-conventions.md` with new data structures.
  - Keep `README.md` and user-facing docs synchronized.

### 3. Clear Structure & Style
- Use GitHub Flavored Markdown (GFM).
- Provide copy-pasteable CLI commands or code examples verified against tests.
- Include prerequisites, installation steps, configuration tables (with types and defaults), and troubleshooting notes.
- Keep explanations crisp, dense, and unambiguous. Avoid marketing fluff.