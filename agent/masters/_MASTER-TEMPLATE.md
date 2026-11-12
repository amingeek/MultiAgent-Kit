# MASTER NN — <Feature Domain Title>

> Read first: `AGENTS.md`, `agent/masters/_common.md`.
> Output: Milestone file `agent/Milestones/NN-<feature-slug>.md`

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

## 4. Decomposed Milestone Plan

### Milestone M<NN>A: Contract & Schema Layer
- **Role:** `architect`
- **Goal:** Define interfaces, DTOs, and migration scripts.
- **Artifacts:** `src/models/...`, `migrations/...`

### Milestone M<NN>B: Core Implementation Layer
- **Role:** `developer`
- **Goal:** Implement service logic, handlers, and persistence integration.
- **Artifacts:** `src/services/...`, `src/handlers/...`

### Milestone M<NN>C: Verification & Test Automation
- **Role:** `qa`
- **Goal:** Write unit, integration, and failure-injection tests.
- **Artifacts:** `tests/...`
