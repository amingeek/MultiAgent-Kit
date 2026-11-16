# Repository Architectural Map (PROJECT_MAP)

> **Generated & Maintained By:** `agent/masters/00-info.md` (Project Reconnaissance)
> **Usage:** Read this map to understand codebase boundaries, build commands, and execution flows without full repository re-indexing.

---

## 1. Quick Start & Toolchain
- **Project Name:** Universal MultiAgent-Kit Project
- **Detected Tech Stack:** Universal / Polyglot
- **Build Command:** `<e.g. make build, npm run build, cargo build>`
- **Test Command:** `<e.g. pytest -v, go test ./..., npm test, cargo test>`
- **Lint Command:** `<e.g. ruff check, eslint ., golangci-lint run>`
- **Container Runtime:** `<Docker / docker compose / none>`

---

## 2. Directory Layout & Layering
| Directory | Responsibility | Architectural Layer |
|---|---|---|
| `agent/` | Multi-agent engineering tooling & state | Tooling / Meta |
| `src/` or `app/` | Core application source code | Application |
| `tests/` | Unit, integration, and E2E test suites | Verification |

---

## 3. Key Request Flows
1. **Flow A (Reconnaissance):** `BOOTSTRAP.md` -> `00-info.md` -> `PROJECT_MAP.md` & `10-contracts-and-conventions.md`.
2. **Flow B (Feature Implementation):** `NEW-FEATURE.md` -> Master Blueprint -> Granular Milestones -> `EXECUTE.md`.

---

## 4. Discovered Contracts & Data Models
- **Core Entities:** *(Populated upon project initialization)*
- **Primary Routes / APIs:** *(Populated upon project initialization)*
- **Configuration & Secrets:** `.env.example` (Never commit `.env`)