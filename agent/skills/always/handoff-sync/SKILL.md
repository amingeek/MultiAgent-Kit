---
name: handoff-sync
type: always
description: Multi-agent typed JSON contract alignment and state-machine transition validation.
---
# Handoff-Sync State Machine Protocol

Guarantee strict typed communication across isolated agent roles adhering to `agent/HANDOFF_SCHEMA.md`.

## State-Machine Transition Rules
- **Architect $\rightarrow$ Developer:**
  `status = SUCCESS` $\rightarrow$ `current_step.assigned_role = developer`. Passes clean interfaces and DDL.
- **Developer $\rightarrow$ QA:**
  `status = IN_PROGRESS` $\rightarrow$ `current_step.assigned_role = qa`. Passes pure implementation without tests.
- **QA Pass $\rightarrow$ Orchestrator:**
  `status = SUCCESS` $\rightarrow$ `feedback_loop.qa_retry_count = 0`. Orchestrator marks step complete and commits.
- **QA Fail $\rightarrow$ Developer (Feedback Loop):**
  Increment `qa_retry_count++`.
  - If `qa_retry_count < max_retries (3)`: `status = QA_FAILED`, populate `error_logs` and assign to `developer`.
  - If `qa_retry_count >= max_retries (3)`: `status = BLOCKED`, log to `Failed-approaches.md` and halt for human help.

## Pre-Handoff Integrity Check
Run the Go CLI validator before concluding any role turn:
```bash
./agent/tools/mak validate
```
Verify exit code is 0. If validation fails, correct `agent/.state/handoff.json` before yielding control.