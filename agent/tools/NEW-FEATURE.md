# NEW FEATURE → MASTER & MILESTONE GENERATOR

> **Usage:**
> In your AI CLI, run:
> `Read AGENTS.md, then read agent/tools/NEW-FEATURE.md and follow it exactly.`
> `Feature Request: <describe feature in detail>`

---

## Agent Instructions (Architect / Orchestrator Mode)

You are acting as the **System Architect** formulating a complete, production-grade engineering plan for a new feature request. Do not write application code now; your task is planning, decomposition, and state registration.

### Step 1: Analyze & Clarify Feature Scope
1. Inspect `agent/info/PROJECT_MAP.md` and `agent/rules/10-contracts-and-conventions.md`.
2. Determine required modules, interfaces, database schemas, and endpoints.
3. Check for architectural constraints, security implications, and potential breaking changes.

### Step 2: Create Master Planning File
Determine the next available number `NN` in `agent/masters/` (e.g. `01`, `02`).
Create `agent/masters/NN-<feature-slug>.md` using this structure:
```markdown
# MASTER NN — <Feature Title>

## Objective
<Detailed description of the feature and problem solved>

## Architecture & Boundary Impact
- Target modules / packages
- New DTOs, schemas, or DB tables
- External dependencies (if any)

## Security & Reliability Gates
- Auth requirements & permission checks
- Input validation & rate limiting
- Safe fallback / degradation policy

## Sequential Milestone Breakdown
1. Milestone M<NN>A: Contract & Schema Definition (Architect)
2. Milestone M<NN>B: Business Logic Implementation (Developer)
3. Milestone M<NN>C: Comprehensive Test Suite & QA (QA Engineer)
```

### Step 3: Generate Concrete Milestone File(s)
In `agent/Milestones/`, create the corresponding actionable milestone file `NN-<feature-slug>.md`:
Follow `agent/Milestones/_MILESTONE-TEMPLATE.md`.
Each milestone must have:
- Clear list of ordered steps.
- Assigned role for each step (`architect`, `developer`, `qa`, etc.).
- Target files list.
- Concrete acceptance checklist with test criteria.

### Step 4: Register in Global Indices
1. Register milestone in `agent/Milestones/INDEX.md`.
2. Update `agent/STATE.md` under `Milestones Summary`.
3. Update `agent/.state/handoff.json`:
   - Set `active_milestone` to the new milestone ID and path.
   - Set `current_step` to step 1 (assigned to `architect` or `developer`).
   - Reset `status` to `IDLE` or `IN_PROGRESS`.
   - Reset `feedback_loop.qa_retry_count` to 0.

### Step 5: Validate & Output Next Command
Run: `python3 agent/tools/validate_state.py` to confirm state validity.
Prompt the user with the exact command to start execution:
`Read AGENTS.md, then agent/masters/EXECUTE.md and follow it exactly.`