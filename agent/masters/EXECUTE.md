# 🚀 EXECUTE — Multi-Agent Orchestrator Execution Engine

> Usage: `Read AGENTS.md, then agent/masters/EXECUTE.md and follow it exactly.`
> Modes:
> - Step-by-Step (Default): Orchestrator delegates via handoff.json, then stops.
> - Autonomous: `MODE: RUN ALL MILESTONES, FULLY AUTONOMOUS.` Orchestrator cycles roles continuously.

## Execution Cycle
1. Read `agent/STATE.md` and find lowest uncompleted milestone.
2. Identify required role (Architect, Developer, QA, DevOps).
3. Update `agent/.state/handoff.json`. Reset `qa_retry_count` if new step.
4. Autonomous mode: Invoke role internally, respecting boundaries.
5. If QA fails, return to Developer. Enforce `max_retries` (3) rule from HANDOFF_SCHEMA.md. If blocked, halt.
6. When QA passes (`status: SUCCESS`), tick checklist, commit `M<NN> step <X>`, move to next step.
