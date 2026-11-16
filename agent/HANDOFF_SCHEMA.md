# 🤝 Multi-Agent Handoff Protocol & Schema (Strict)

All agents MUST read and write to `agent/.state/handoff.json`. Unstructured plain text delegation in chat is strictly forbidden.

## Schema Structure (`agent/.state/handoff.json`)
```json
{
  "$schema": "agent/HANDOFF_SCHEMA.md",
  "project_name": "string",
  "tech_stack": ["string"],
  "active_milestone": {
    "id": "M00 | M<NN>-<MM>",
    "title": "string",
    "source_file": "agent/masters/00-info.md | agent/Milestones/<NN>-<MM>-<slug>.md"
  },
  "current_step": {
    "step_number": 0,
    "description": "Concrete task description",
    "assigned_role": "orchestrator | architect | developer | qa | devops | security",
    "role_file": "agent/roles/<NN>-<role>.md",
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
  "last_updated": "2026-11-15T11:25:00Z"
}
```

## Field Specifications & Enums
- **`status`**:
  - `IDLE`: Ready to begin next step or waiting for operator assignment.
  - `IN_PROGRESS`: Actively being executed by assigned worker role.
  - `QA_FAILED`: Test suite failed; assigned back to Developer with error logs.
  - `BLOCKED`: Retries exceeded `max_retries` (3) or manual human decision required.
  - `SUCCESS`: Current step verified passing; ready for Orchestrator to advance.
- **`assigned_role`**:
  - `orchestrator`, `architect`, `developer`, `qa`, `devops`, `security`.
- **`feedback_loop.qa_retry_count`**:
  - Incremented by 1 whenever QA tests fail.
  - If `qa_retry_count >= max_retries`, task MUST transition to `BLOCKED` to prevent infinite hallucination loops.
  - Reset to 0 when starting a new step or after human intervention unblocks the task.
