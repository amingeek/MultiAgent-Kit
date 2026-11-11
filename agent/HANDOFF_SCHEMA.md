# 🤝 Multi-Agent Handoff Protocol & Schema (Strict)

All agents MUST read and write to `agent/.state/handoff.json`. Unstructured plain text delegation in chat is strictly forbidden.

## Schema Structure (`agent/.state/handoff.json`)
```json
{
  "$schema": "agent/HANDOFF_SCHEMA.md",
  "project_name": "string",
  "tech_stack": ["string"],
  "active_milestone": {
    "id": "M01",
    "title": "string",
    "source_file": "agent/Milestones/01-name.md"
  },
  "current_step": {
    "step_number": 1,
    "description": "Concrete task description",
    "assigned_role": "orchestrator | architect | developer | qa | devops | security",
    "role_file": "agent/roles/02-developer.md",
    "target_files": ["src/service.go", "src/types.go"]
  },
  "status": "IDLE | IN_PROGRESS | BLOCKED | QA_FAILED | SUCCESS",
  "context_summary": "Minimal high-signal context. Max 10 lines.",
  "feedback_loop": {
    "qa_retry_count": 0,
    "max_retries": 3,
    "error_logs": "Exact compiler output or test trace",
    "failing_tests": ["TestName"],
    "required_fix": "Fix error handling without altering contract"
  },
  "acceptance_checklist": [
    {"criterion": "Endpoints return 200", "verified": false}
  ],
  "last_updated": "2026-11-11T12:00:00Z"
}
```
*Rule: If `feedback_loop.qa_retry_count` reaches `max_retries` (3), the task MUST be marked `BLOCKED` and returned to Orchestrator to prevent infinite hallucination loops.*
