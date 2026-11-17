---
name: handoff-sync
type: always
description: Multi-agent state alignment and strict JSON validation before turn end.
---
# Handoff-Sync Protocol

Ensure every handoff between agents strictly adheres to `agent/HANDOFF_SCHEMA.md` and maintains state integrity.

## Synchronization Protocol
1. **Schema Validation:**
   - Before completing your turn, ensure `agent/.state/handoff.json` contains valid JSON.
   - Run verification via CLI: `./agent/tools/mak validate`.
2. **Mandatory Keys Verification:**
   - `active_milestone`: Must include `id`, `title`, and `source_file`.
   - `current_step`: Must include `step_number`, `description`, `assigned_role`, `role_file`, and `target_files`.
   - `status`: Must be one of `IDLE`, `IN_PROGRESS`, `QA_FAILED`, `BLOCKED`, `SUCCESS`.
   - `feedback_loop`: Must have valid `qa_retry_count` ($\le$ `max_retries`).
3. **Transition Rules:**
   - Architect $\rightarrow$ Developer: status=`SUCCESS`, assign to Developer.
   - Developer $\rightarrow$ QA: status=`IN_PROGRESS`, assign to QA.
   - QA Pass $\rightarrow$ Orchestrator: status=`SUCCESS`, reset `qa_retry_count`=0.
   - QA Fail $\rightarrow$ Developer: increment `qa_retry_count`. If $\ge$ 3, status=`BLOCKED`, stop.