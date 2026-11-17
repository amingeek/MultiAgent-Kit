---
name: checkpointing
type: always
description: Atomic state checkpointing and seamless interruption recovery across agent turns.
---
# Atomic Checkpointing & State Recovery

Ensure zero loss of progress and deterministic task resumption across network drops, session resets, or subagent handoffs.

## Checkpoint Protocol
1. **Atomic Step Commits:**
   - Commit exactly once per verified step: `git add <target_files> agent/ && git commit -m "M<NN>-<MM> step <X>: <description>"`
   - Never leave dirty uncommitted files after a step transition.
2. **Synchronous State Persistence:**
   - Immediately update `agent/STATE.md` with timestamp, completed task log, and CLI test evidence.
   - Synchronize timestamps via Go CLI: `./agent/tools/mak touch`.
3. **Session Resumption Heuristic:**
   - On session startup, read `agent/.state/handoff.json` first:
     - If `status == "IN_PROGRESS"`: Resume work on `current_step.target_files`.
     - If `status == "QA_FAILED"`: Inspect `feedback_loop.error_logs`, reproduce, and fix.
     - If `status == "BLOCKED"`: Halt immediately and notify human operator. Do not loop.
     - If `status == "SUCCESS"` / `"IDLE"`: Advance to lowest unchecked step in active milestone.