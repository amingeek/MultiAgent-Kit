# REFACTORING & DECOMPOSITION PROMPT: UPGRADE ARCHITECTURE TO GRANULAR MILESTONES (v5.3)

> **Instructions for the Operator:**
> Copy and paste the entire prompt below into your target agent's terminal or chat interface (Claude Code, Antigravity, Cursor, etc.).

---

```markdown
Read AGENTS.md, then read this instruction prompt completely and follow it with absolute priority.

## Mission Objective
You are operating as the **Principal Software Architect & System Orchestrator**. 
Your current planning contains a shallow 1-to-1 mapping between Masters and Milestones (10 Masters mapped to only 10 Milestones). This violates the production-grade v5.3 architecture of MultiAgent-Kit.

Before ANY code implementation begins, you must perform a comprehensive architectural refactoring to convert each Master into a macro **Epic Architectural Blueprint** and decompose it into **granular, highly detailed Milestones** (each Master decomposing into multiple granular Milestones, up to 5-10 Milestones per domain depending on complexity).

---

## Non-Negotiable Architectural Principles
1. **Master = Macro Architectural Blueprint (Epic):**
   - Located in `agent/masters/NN-<feature-slug>.md`.
   - Never contains direct step-by-step code instructions.
   - Defines high-level problem statements, architectural boundaries, subsystem impacts, security/resilience gates, and an explicit **Milestone Decomposition DAG**.
2. **Milestones = Granular Execution Units:**
   - Located in `agent/Milestones/NN-MM-<milestone-slug>.md` (e.g., `01-01-schema-migration.md`, `01-02-contract-definition.md`, etc.).
   - Each milestone is an isolated, independently testable slice affecting $\le 5$ files and requiring 1–2 hours of focused work.
   - Must contain deep, production-ready technical contracts (exact types, fields, signatures, error codes, state invariants).
   - Must define explicit negative scenarios, race condition handling, timeout fallbacks, and idempotency guarantees.
   - Must contain role-segregated actionable steps (`architect`, `developer`, `qa`, `devops`, `security`).
   - Must contain a concrete TDD Verification Matrix.

---

## Step-by-Step Refactoring Plan

### Phase 1: Audit & Refactor the 10 Masters
For each master in `agent/masters/` (`01-*.md` through `10-*.md`):
1. Review the existing master content.
2. Restructure it strictly according to `agent/masters/_MASTER-TEMPLATE.md`:
   - **Section 1:** Executive Summary, Problem Statement, Success Criteria.
   - **Section 2:** Architectural Boundaries, Impacted Subsystems, Contract Changes.
   - **Section 3:** Security, Performance, and Reliability Gates.
   - **Section 4:** **Decomposed Granular Milestone Plan (DAG)**:
     - Decompose the master into 3 to 10+ granular milestones (`M<NN>-01`, `M<NN>-02`, ...).
     - Provide a clear decomposition table specifying: `Milestone ID`, `Title`, `Primary Role`, `Depends On`, `Target Files / Scope`, and `Milestone File`.

### Phase 2: Refactor & Expand `agent/Milestones/`
1. Remove or replace the previous shallow 1-to-1 milestone files.
2. For EVERY milestone defined across all 10 Masters' DAG tables, generate a dedicated file:
   - Path convention: `agent/Milestones/NN-MM-<milestone-slug>.md` (e.g., `01-01-*.md`, `01-02-*.md`).
   - Follow `agent/Milestones/_MILESTONE-TEMPLATE.md` strictly.
3. Every milestone file MUST be fully fleshed out with deep technical detail:
   - **Deep Technical Specification:** Exact code signatures, types, interfaces, DTOs, tables, and error codes.
   - **Negative Scenarios & Edge Cases:** Validation errors, race conditions, duplicate requests, timeouts.
   - **Sequential Actionable Steps:** Role-assigned steps with explicit target files and checklists.
   - **TDD Verification Matrix:** Specific test cases (`TC-01`, `TC-02`, etc.) with expected outcomes.

### Phase 3: Synchronize Registries & Execution State
1. **Rebuild `agent/Milestones/INDEX.md`:**
   - Register all decomposed milestones in sequential order with their IDs, file paths, titles, source masters, dependencies (`Depends On`), and initial status (`NOT_STARTED`).
2. **Update `agent/STATE.md`:**
   - Update `Last update` timestamp.
   - Under `## Milestones Summary`, document the new hierarchical structure and total milestone count across the 10 Masters.
   - Add an entry in `## Quick Tasks Log` recording the v5.3 architectural refactoring.
3. **Initialize `agent/.state/handoff.json`:**
   - Set `active_milestone`:
     - `id`: `M01-01`
     - `title`: Title of the first milestone.
     - `source_file`: `agent/Milestones/01-01-<slug>.md`
   - Set `current_step` to Step 1 of milestone `M01-01`.
   - Set `status`: `IDLE` (or `IN_PROGRESS`).
   - Reset `feedback_loop.qa_retry_count` to 0.

### Phase 4: Deterministic Validation
Run the state validator CLI:
```bash
python3 agent/tools/validate_state.py
```
Ensure it returns exit code 0: `[OK] State validation successful. handoff.json and STATE.md are valid.`

---

## Final Output Report
Once complete, output a structured summary table showing:
- Each Master (`01` through `10`) and how many Granular Milestones it was decomposed into.
- Total count of newly generated milestone files.
- Confirmation that `validate_state.py` passes cleanly.
- Ready command for starting development: `Read AGENTS.md, then agent/masters/EXECUTE.md and follow it exactly.`
```
