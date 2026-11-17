# NEW FEATURE → MASTER & GRANULAR MILESTONES GENERATOR

> **Usage:**
> In your AI CLI, run:
> `Read AGENTS.md, then read agent/tools/NEW-FEATURE.md and follow it exactly.`
> `Feature Request: <describe feature in detail>`

---

## Agent Instructions (Architect / Orchestrator Mode)

You are acting as the **System Architect** formulating a complete, production-grade engineering plan for a new feature request. Do not write application code now; your task is macro-architecture planning, granular decomposition, and state registration.

### Core Philosophy: Master vs. Milestone Hierarchy
- **Master (`agent/masters/NN-<slug>.md`):** High-level Architectural Epic Blueprint. Defines problem statement, subsystem impacts, overarching security/performance constraints, and the milestone DAG.
- **Milestones (`agent/Milestones/NN-MM-<slug>.md`):** Granular, production-ready execution units. **A single Master decomposes into 1 to 10+ sequential Milestones.** Each milestone translates the macro architecture into exhaustive code contracts, edge-case handling, and step-by-step role tasks ($\le 5$ files, 1-2 hours).

---

### Step 1: Analyze & Clarify Feature Scope
1. Inspect `agent/info/PROJECT_MAP.md` and `agent/rules/10-contracts-and-conventions.md`.
2. Determine required modules, interfaces, database schemas, and endpoints.
3. Check for architectural constraints, security implications, and potential breaking changes.

---

### Step 2: Create Master Planning File (Epic Blueprint)
Determine the next available number `NN` in `agent/masters/` (e.g. `01`, `02`).
Create `agent/masters/NN-<feature-slug>.md` following `agent/masters/_MASTER-TEMPLATE.md`.
Ensure section 4 establishes the complete **Decomposed Granular Milestone Plan (DAG)**, breaking down the feature into **1 to 10+ sequential milestones** (`M<NN>-01`, `M<NN>-02`, ... up to `M<NN>-10+`).

---

### Step 3: Generate Granular Milestone Files
For **EVERY** milestone defined in the Master's DAG:
Create a dedicated, comprehensive file in `agent/Milestones/NN-MM-<milestone-slug>.md` (e.g. `01-01-schema-migration.md`, `01-02-api-contracts.md`, ... `01-10-e2e-qa.md`).
Follow `agent/Milestones/_MILESTONE-TEMPLATE.md` strictly.

Every milestone file MUST include:
1. **Source Master Reference:** Exact path to `agent/masters/NN-<feature-slug>.md`.
2. **Deep Technical Specifications:** Concrete interfaces, data models, method signatures, status codes, and invariants.
3. **Negative Scenarios & Edge Cases:** Validation errors, race conditions, timeout fallbacks, idempotency rules.
4. **Actionable Steps:** Role-assigned steps (`architect`, `developer`, `qa`, `devops`, `security`) with explicit `Target Files`.
5. **TDD Verification Matrix:** Specific test cases and acceptance checklist.

---

### Step 4: Register in Global Indices
1. Register **all** generated milestones in `agent/Milestones/INDEX.md` in sequential order:
   ```markdown
   | ID | File | Title | Source Master | Depends On | Status |
   | M01-01 | agent/Milestones/01-01-schema.md | Schema & DDL | agent/masters/01-auth.md | None | NOT_STARTED |
   | M01-02 | agent/Milestones/01-02-contracts.md | Contract Definitions | agent/masters/01-auth.md | M01-01 | NOT_STARTED |
   ```
2. Update `agent/STATE.md` under `## Milestones Summary` referencing the master and its milestones.
3. Update `agent/.state/handoff.json`:
   - Set `active_milestone` to the first milestone (`M<NN>-01`).
   - Set `current_step` to step 1 of milestone `M<NN>-01`.
   - Set `status` to `IDLE` (or `IN_PROGRESS`).
   - Reset `feedback_loop.qa_retry_count` to 0.

---

### Step 5: Validate & Output Next Command
Run: `./agent/tools/mak validate` to confirm state validity.
Prompt the user with the exact command to start execution:
```bash
Read AGENTS.md, then agent/masters/EXECUTE.md and follow it exactly.
```