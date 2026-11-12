---
name: checkpointing
type: always
description: Interruption recovery and atomic state checkpointing across steps.
---
# Atomic Checkpointing

Prevent context loss and enable zero-friction task resumption across interruptions or agent role transitions.

## Checkpointing Procedure
1. **Atomic Step Completion:**
   - Commit changes atomically per completed step: `git add <files>` then `git commit -m "M<NN> step <X>: <description>"`.
   - Never leave uncommitted dirty work in repo after completing a step.
2. **State File Synchronization:**
   - Immediately update `agent/STATE.md`:
     - Set `Last update` to current ISO timestamp.
     - Append completed step to `Quick Tasks Log` or update `Milestones Summary`.
     - Record CLI test execution evidence under `Execution Evidence Log`.
3. **Interruption Recovery Protocol:**
   - When woken up or starting a session: inspect `agent/.state/handoff.json` first.
   - If `status == "IN_PROGRESS"`, resume from `current_step.target_files`.
   - If `status == "QA_FAILED"`, read `feedback_loop.error_logs` and resolve failure.
   - If `status == "BLOCKED"`, do not proceed autonomously; report blocker to operator.