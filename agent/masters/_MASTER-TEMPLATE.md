# MASTER NN — <Feature Domain Title>

> Read first: `AGENTS.md`, `agent/masters/_common.md`.
> Outputs: Granular Milestone files in `agent/Milestones/NN-MM-<milestone-slug>.md` (1 to 10+ detailed milestones)

---

## 1. Executive Summary & Objective
- **Problem Statement:** <What problem does this feature solve?>
- **Target Value:** <Why is this feature being implemented?>
- **Success Criteria:** <Measurable outcomes>

---

## 2. Architectural Boundaries & Dependencies
- **Impacted Subsystems:** <List packages/directories>
- **Contract Changes:** <New APIs, schemas, DTOs, or database tables>
- **Third-Party Dependencies:** <List new libraries, if strictly mandated; otherwise state "None">

---

## 3. Security, Performance & Reliability Gates
- **Authentication & Authorization:** <Required roles, scopes, or token checks>
- **Input Sanitization:** <Validation rules, bounds checking, rate limits>
- **Resilience / Fallback:** <Timeout values, circuit breakers, error responses>
- **Performance Budget:** <Target latency, memory limits, indexing strategy>

---

## 4. Decomposed Granular Milestone Plan (DAG)
Decompose this master into **1 to 10+ sequential, highly granular milestones**. Each milestone focuses on a specific architectural slice ($\le 5$ files, 1-2 hours) and must have its own dedicated file in `agent/Milestones/NN-MM-<milestone-slug>.md`.

| Milestone ID | Title | Primary Role | Depends On | Target Files / Scope | Milestone File |
|---|---|---|---|---|---|
| `M<NN>-01` | Database Schema & DDL Migration | `architect` | None | `migrations/`, `src/models/` | `agent/Milestones/<NN>-01-schema-migration.md` |
| `M<NN>-02` | Contract & Interface Definitions | `architect` | `M<NN>-01` | `src/contracts/`, `src/dto/` | `agent/Milestones/<NN>-02-contract-definition.md` |
| `M<NN>-03` | Data Access / Repository Layer | `developer` | `M<NN>-02` | `src/repositories/` | `agent/Milestones/<NN>-03-repository-layer.md` |
| `M<NN>-04` | Core Domain & Service Logic | `developer` | `M<NN>-03` | `src/services/` | `agent/Milestones/<NN>-04-domain-services.md` |
| `M<NN>-05` | Security Guards & Middleware | `developer` | `M<NN>-04` | `src/middleware/`, `src/auth/` | `agent/Milestones/<NN>-05-auth-middleware.md` |
| `M<NN>-06` | Transport / API Endpoints | `developer` | `M<NN>-05` | `src/controllers/`, `src/routes/` | `agent/Milestones/<NN>-06-api-endpoints.md` |
| `M<NN>-07` | Edge-Case & Error Resilience | `developer` | `M<NN>-06` | `src/errors/`, `src/services/` | `agent/Milestones/<NN>-07-error-handling.md` |
| `M<NN>-08` | Unit & TDD Test Suite | `qa` | `M<NN>-07` | `tests/unit/` | `agent/Milestones/<NN>-08-unit-tests.md` |
| `M<NN>-09` | Integration & Failure Injection Tests | `qa` | `M<NN>-08` | `tests/integration/` | `agent/Milestones/<NN>-09-integration-tests.md` |
| `M<NN>-10` | CI/CD & Deployment Verification | `devops` | `M<NN>-09` | `.github/`, `Dockerfile` | `agent/Milestones/<NN>-10-deployment-ci.md` |

*(Note: Scale the number of milestones from 1 to 10+ depending on feature complexity. Never lump disparate concerns into a single massive milestone).*
