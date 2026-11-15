# Milestone M<NN>-<MM>: <Milestone Title>

- **Source Master:** `agent/masters/NN-<slug>.md`
- **Milestone ID:** `M<NN>-<MM>` (e.g. `M01-01`, `M01-02` ... up to `M01-10+`)
- **Depends On:** <None or previous milestone ID, e.g. M<NN>-<MM-1>>
- **Status:** NOT_STARTED | IN_PROGRESS | COMPLETED | BLOCKED
- **Scope Budget:** $\le 5$ files | 1-2 engineering hours

---

## 1. Deep Technical Specification & Contract Alignment
> *Note: This section translates the macro design of the Source Master into concrete, granular technical contracts for this specific milestone.*

### 1.1 Architectural Scope & Objectives
<Detailed description of the exact component, layer, or subsystem implemented in this milestone.>

### 1.2 Data Structures, Interfaces & API Contracts
```typescript/python/go
// Explicit code signatures, data models, DTOs, or interfaces to be implemented/extended
```
- **New / Modified Entities:** `<List exact model names, tables, or DTOs>`
- **Expected Return Codes & Error Formats:** `<List error codes, e.g. ERR_INVALID_TOKEN, 401 Unauthorized>`
- **State Invariants:** `<Non-negotiable invariants, e.g. balance cannot be negative>`

---

## 2. Negative Scenarios, Edge Cases & Resilience Gates
> *Every milestone must explicitly define negative paths before implementation.*

- **Validation Failures:** `<Handling of missing fields, malformed types, bounds violations>`
- **Concurrency & Race Conditions:** `<Handling of duplicate concurrent requests, locking strategy>`
- **Timeouts & Network Partitions:** `<Circuit breaking, fallback response, retry behavior>`
- **Idempotency Guarantee:** `<How duplicate calls are safely handled>`

---

## 3. Sequential Actionable Execution Steps

### Step 1: <Short Title (e.g. Contract Definition / DDL Schema)>
- **Assigned Role:** `architect` | `developer` | `qa` | `devops` | `security`
- **Target Files:**
  - `path/to/target1`
  - `path/to/target2`
- **Detailed Technical Instructions:**
  <Provide precise, code-level instructions for the assigned role. Specify functions, classes, and types.>
- **Acceptance Checklist:**
  - [ ] Contract/code implemented without syntax or type errors
  - [ ] Verified against project coding standards (`agent/rules/01-code-quality.md`)
  - [ ] No regression introduced to existing components

### Step 2: <Short Title (e.g. Core Implementation / Business Logic)>
- **Assigned Role:** `developer`
- **Target Files:**
  - `path/to/target`
- **Detailed Technical Instructions:**
  <Specific logic implementation guidelines. Developer MUST NOT write tests.>
- **Acceptance Checklist:**
  - [ ] Logic correctly satisfies section 1 specifications
  - [ ] Negative and edge case handling matches section 2
  - [ ] Pure implementation without test modifications

### Step 3: <Short Title (e.g. Comprehensive TDD Test Suite & Verification)>
- **Assigned Role:** `qa`
- **Target Files:**
  - `tests/path/to/test`
- **Detailed Technical Instructions:**
  <Write deterministic unit, integration, and edge-case tests. QA MUST NOT touch application code.>
- **Acceptance Checklist:**
  - [ ] 100% of new unit tests pass
  - [ ] All negative scenarios from section 2 are covered
  - [ ] Overall test count $\ge$ previous baseline
  - [ ] Test command executes cleanly with exit code 0

---

## 4. TDD Verification Matrix
| Test Case ID | Target Component | Scenario | Expected Outcome | Role |
|---|---|---|---|---|
| `TC-01` | `<Component>` | Happy path execution with valid inputs | Success response (200 / valid entity) | `qa` |
| `TC-02` | `<Component>` | Missing required parameters / malformed input | Validation error / 400 Bad Request | `qa` |
| `TC-03` | `<Component>` | Concurrent identical requests | Idempotent handling / safe lock | `qa` |
| `TC-04` | `<Component>` | Downstream dependency timeout | Fallback / graceful degradation | `qa` |

---

## 5. Verification Evidence Log
- Date: `YYYY-MM-DD`
- Commit: `M<NN>-<MM> step <X>: <description>`
- CLI Verification Command: `<command>`
- Test Suite Output Summary: `<summary>`
