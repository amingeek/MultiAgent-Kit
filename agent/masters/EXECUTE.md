# 🚀 EXECUTE — Multi-Agent Orchestration & Execution Engine

> **Usage:**
> In your AI CLI, run:
> `Read AGENTS.md, then agent/masters/EXECUTE.md and follow it exactly.`
>
> **Execution Modes:**
> - **Autonomous Mode:** `MODE: FULLY AUTONOMOUS MULTI-AGENT EXECUTION.`
>   (Orchestrator cycles roles or dispatches subagents continuously until the milestone is completed or blocked).
> - **Step-by-Step Mode (Default):**
>   (Orchestrator inspects progress, prepares the next step in `handoff.json`, and stops to prompt the operator).

---

## The Supervisor-Worker Execution Cycle

```
[Orchestrator]
      │
      ▼ (Inspects milestone & updates handoff.json)
[Architect] (If contract/schema needed) ──► handoff.json (status: SUCCESS)
      │
      ▼
[Developer] (Implements pure logic) ──► handoff.json (status: IN_PROGRESS)
      │
      ▼
[QA Engineer] (Writes & runs tests)
      ├── Pass: status=SUCCESS ──► [Orchestrator] (Commits & advances step)
      └── Fail: qa_retry_count++ ──►
             ├── If < 3: status=QA_FAILED ──► [Developer] (Fixes bug)
             └── If >= 3: status=BLOCKED ──► [Halt & Human Intervention]
```

---

## Detailed Execution Steps

### Step 1: Milestone State Assessment
1. Read `agent/STATE.md` and `agent/Milestones/INDEX.md` to identify the lowest uncompleted milestone.
2. Read the active milestone file in `agent/Milestones/` (e.g. `01-01-schema.md`).
3. Find the lowest unchecked step: `- [ ] Step X: <description>`.
4. If all steps in the milestone are checked:
   - Mark milestone completed in `Milestones/INDEX.md` and `STATE.md`.
   - Look for the next milestone, or terminate with `ALL MILESTONES COMPLETE`.

### Step 2: Role Assignment & Handoff Preparation
1. Determine the appropriate role for the step:
   - Interface / Schema / DDL $\rightarrow$ `agent/roles/01-architect.md`
   - Business Logic / Code implementation $\rightarrow$ `agent/roles/02-developer.md`
   - Test suites / Verification $\rightarrow$ `agent/roles/03-qa.md`
   - Security Audit $\rightarrow$ `agent/roles/04-security-reviewer.md`
   - Docker / CI/CD $\rightarrow$ `agent/roles/05-devops.md`
2. Update `agent/.state/handoff.json`:
   - `current_step.step_number`: Current step index.
   - `current_step.description`: Concrete actionable instructions.
   - `current_step.assigned_role`: Target role.
   - `current_step.role_file`: Path to the role specification.
   - `current_step.target_files`: Explicit list of files to read/edit.
   - `status`: `IN_PROGRESS` (or `IDLE`).
   - If starting a brand new step: reset `feedback_loop.qa_retry_count` to 0.

### Step 3: Worker Execution
- **In Autonomous Mode with Subagents:**
  Spawn or invoke the subagent for `current_step.assigned_role` with a clean context, passing `agent/.state/handoff.json`. Wait for worker completion.
- **In Autonomous Mode (Single Agent):**
  Adopt the assigned role's persona strictly adhering to its boundaries (e.g. Developer never edits tests; QA never edits application logic). Execute the task and update `handoff.json`.
- **In Step-by-Step Mode:**
  Output the prompt for the operator to invoke the worker:
  `Next role: <role>. Run: Read AGENTS.md, then <role_file> and agent/.state/handoff.json.` Stop turn.

### Step 4: QA & Anti-Looping Feedback Loop
When the worker is `QA`:
1. Execute the project test suite using the verified test command from `agent/info/PROJECT_MAP.md`.
2. **If ALL tests pass:**
   - Set `status` to `SUCCESS`.
   - Reset `feedback_loop.qa_retry_count` to 0.
   - Check the step box in the milestone file: `- [x] Step X: ...`.
   - Update `agent/STATE.md` with test evidence.
   - Stage and commit atomically: `git add <target_files> agent/ && git commit -m "M<NN>-<MM> step <X>: <description>"`.
   - Advance to next step.
3. **If ANY test fails:**
   - Increment `feedback_loop.qa_retry_count` by 1.
   - If `qa_retry_count >= max_retries` (3):
     - Set `status` to `BLOCKED`.
     - Log failure trace into `agent/Failed-approaches.md`.
     - **HALT IMMEDIATELY**. Do not loop further. Report blocker to operator.
   - If `qa_retry_count < max_retries`:
     - Set `status` to `QA_FAILED`.
     - Populate `feedback_loop.error_logs` with exact failure output and `failing_tests`.
     - Assign back to `roles/02-developer.md`.

### Step 5: Post-Step Health Check
Run `./agent/tools/validate_state` (or `python3 agent/tools/validate_state.py`) to ensure state validity after every transition.
