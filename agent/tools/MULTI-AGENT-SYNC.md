# MULTI-AGENT SYNC & STATE VERIFICATION

> **Usage:**
> In your AI CLI or terminal, run:
> `Read AGENTS.md, then read agent/tools/MULTI-AGENT-SYNC.md and follow it exactly.`
> Or run directly: `./agent/tools/mak validate`

---

## Purpose
Ensures that the multi-agent system state is completely synchronized, valid, and free of contradictions between:
- `agent/.state/handoff.json` (machine state)
- `agent/STATE.md` (human / agent readable log)
- Active milestone files in `agent/Milestones/`
- Current git workspace status

---

## 4-Stage Synchronization Procedure

### Stage 1: Automated Schema Validation
Run the validator tool:
```bash
./agent/tools/mak check-all
```
If errors are reported, fix `agent/.state/handoff.json` or `agent/STATE.md` immediately until the validator outputs:
`[OK] State validation successful.`

### Stage 2: Cross-File Consistency Audit
Verify alignment across all state files:
1. **Milestone Alignment:** Check that `active_milestone.source_file` in `handoff.json` matches the currently active milestone in `agent/STATE.md` and exists in `agent/Milestones/`.
2. **Role Alignment:** Verify that `current_step.assigned_role` corresponds to `current_step.role_file` in `agent/roles/`.
3. **Checklist Parity:** Check that unchecked items in `handoff.json` match unchecked items in the milestone file.

### Stage 3: Anti-Loop & Failure Gate Check
1. Inspect `feedback_loop.qa_retry_count`.
2. If `qa_retry_count >= max_retries` (default 3):
   - Status MUST be set to `BLOCKED`.
   - Record exact error in `agent/Failed-approaches.md`.
   - Halt execution and request operator intervention. Never clear retries manually to loop again.

### Stage 4: Git Cleanliness Check
1. Run `git status -s`.
2. If completed files are uncommitted, stage and commit with proper convention (`M<NN> step <X>: <description>`).
3. Ensure no temporary artifacts or secret files are untracked.