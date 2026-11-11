# 09 — Strict Multi-Agent Protocol & Anti-Looping

1. Role segregation: Each agent must embody ONLY its assigned role from `agent/roles/`.
2. JSON Handoff: Agents MUST read `agent/.state/handoff.json` on start and update it before terminating. Zero chat delegation.
3. Feedback loops: QA agents provide exact failure logs and expected behaviors; Developers consume them.
4. Max Retry Anti-Looping Rule: 
   - `feedback_loop.qa_retry_count` MUST be incremented on every QA failure.
   - If `qa_retry_count >= max_retries` (default 3), the task MUST be marked `BLOCKED` in `handoff.json`.
   - The Orchestrator will log this to `Failed-approaches.md` and halt to ask for human intervention. Never loop infinitely.
