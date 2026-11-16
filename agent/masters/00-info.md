# MASTER 00 — Universal Project Reconnaissance & Architectural Mapping

> **Usage:**
> In your AI CLI, run:
> `Read AGENTS.md, then read agent/masters/00-info.md and follow it exactly.`

---

## Role & Goal
You are acting as the **Principal Systems Architect**. Your mission is to perform comprehensive, non-destructive reconnaissance on this repository, establish baseline execution commands, map all architectural boundaries, and generate two critical foundation documents:
1. `agent/info/PROJECT_MAP.md`
2. `agent/rules/10-contracts-and-conventions.md`

---

## Step-by-Step Reconnaissance Protocol

### Phase 1: Tech Stack & Manifest Detection
Inspect repository root and configuration files to identify:
- **Language & Runtime:** (e.g. Go, Python 3.x, TypeScript/Node, Rust, etc.)
- **Dependency Manifests:** Check for `package.json`, `go.mod`, `Cargo.toml`, `pyproject.toml`, `requirements.txt`, `composer.json`, etc.
- **Build / Test / Lint Toolchains:** Identify exact commands:
  - Build command: (e.g. `npm run build`, `go build ./...`, `cargo build`)
  - Test command: (e.g. `npm test`, `pytest -v`, `go test ./...`, `cargo test`)
  - Lint / Format command: (e.g. `eslint .`, `golangci-lint run`, `ruff check`)
- **Docker & Orchestration:** Check for `Dockerfile`, `docker-compose.yml`, Helm charts, Kubernetes manifests.

### Phase 2: Directory Architecture & Flow Mapping
Perform a targeted file exploration:
1. Identify primary application entry points (`main.go`, `src/index.ts`, `app/main.py`, etc.).
2. Map directory roles (e.g. `api/`, `controllers/`, `services/`, `models/`, `pkg/`, `tests/`).
3. Trace 1-2 core execution flows (e.g. HTTP Request $\rightarrow$ Middleware $\rightarrow$ Controller $\rightarrow$ Service $\rightarrow$ DB).

### Phase 3: Contract Discovery
Extract existing immutable interfaces to prevent accidental breaking changes:
- Public REST routes / GraphQL schemas / gRPC protobuf definitions.
- Core domain entities / DTOs / database tables.
- Environment variables and configuration keys (inspect `.env.example` or config files).

### Phase 4: Generate `agent/info/PROJECT_MAP.md`
Write the comprehensive project map using this standard structure:
```markdown
# Repository Architectural Map (PROJECT_MAP)

## 1. Quick Start & Toolchain
- Tech Stack: <Detected languages & frameworks>
- Build Command: `<command>`
- Test Command: `<command>`
- Lint Command: `<command>`
- Container Runtime: `<docker / compose if present>`

## 2. Directory Layout & Layering
| Directory | Responsibility | Layer |
|---|---|---|
| `src/controllers/` | HTTP request handling & status codes | Transport |
| `src/services/` | Business logic & domain operations | Domain |
| `src/repositories/`| Database access & queries | Persistence |

## 3. Key Request Flows
1. **Flow A (e.g. User Authentication):** `POST /api/auth` -> `AuthController` -> `AuthService` -> `UserRepo` -> JWT response.

## 4. Discovered Contracts & Data Models
- Entities: `<list primary models/tables>`
- Core Routes: `<list key endpoints>`
```

### Phase 5: Populate `agent/rules/10-contracts-and-conventions.md`
Fill in:
- Discovered Immutable Contracts (APIs, schemas, DB tables).
- Repository Conventions (naming convention, error handling idiom, logging standard).

### Phase 6: Synchronize State & Validate
1. Update `agent/.state/handoff.json`:
   - Set `project_name` to detected project name.
   - Set `tech_stack` array to detected technologies.
   - Set `status` to `SUCCESS`.
   - Set `context_summary` to "Reconnaissance complete. Project map and conventions generated."
2. Update `agent/STATE.md`:
   - Update `Last update`, `Active Mode: Idle`, and baseline test suite results.
3. Run state validator:
   ```bash
   ./agent/tools/validate_state
   # (or: python3 agent/tools/validate_state.py)
   ```
4. Output completion summary and recommend next step:
   `Workspace initialized. To plan a feature, run agent/tools/NEW-FEATURE.md.`
