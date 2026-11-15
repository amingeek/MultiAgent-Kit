#!/usr/bin/env python3
"""
MultiAgent-Kit State Validator (v5.3 Production-Ready)
Validates agent/.state/handoff.json and agent/STATE.md against HANDOFF_SCHEMA.md.
Exit code 0 on success, non-zero on validation failure.
"""

import json
import sys
import os
from pathlib import Path

VALID_STATUSES = {"IDLE", "IN_PROGRESS", "BLOCKED", "QA_FAILED", "SUCCESS"}
VALID_ROLES = {"orchestrator", "architect", "developer", "qa", "devops", "security"}

def validate():
    root = Path(__file__).resolve().parent.parent.parent
    handoff_path = root / "agent" / ".state" / "handoff.json"
    state_path = root / "agent" / "STATE.md"
    errors = []

    # 1. Validate STATE.md
    if not state_path.is_file():
        errors.append(f"Missing STATE.md at {state_path}")
    else:
        content = state_path.read_text(encoding="utf-8")
        if "# Global System State" not in content:
            errors.append("STATE.md missing '# Global System State' header")

    # 2. Validate handoff.json
    if not handoff_path.is_file():
        errors.append(f"Missing handoff.json at {handoff_path}")
        report_and_exit(errors)

    try:
        data = json.loads(handoff_path.read_text(encoding="utf-8"))
    except json.JSONDecodeError as e:
        errors.append(f"handoff.json contains invalid JSON: {e}")
        report_and_exit(errors)

    # Required top-level keys
    required_keys = [
        "project_name", "tech_stack", "active_milestone", "current_step",
        "status", "context_summary", "feedback_loop", "acceptance_checklist", "last_updated"
    ]
    for key in required_keys:
        if key not in data:
            errors.append(f"handoff.json missing required key: '{key}'")

    # Validate status enum
    status = data.get("status")
    if status not in VALID_STATUSES:
        errors.append(f"Invalid status '{status}'. Must be one of {sorted(VALID_STATUSES)}")

    # Validate active_milestone
    milestone = data.get("active_milestone")
    if not isinstance(milestone, dict):
        errors.append("'active_milestone' must be an object")
    else:
        for m_key in ["id", "title", "source_file"]:
            if m_key not in milestone or not milestone[m_key]:
                errors.append(f"'active_milestone' missing or empty field: '{m_key}'")

    # Validate current_step
    step = data.get("current_step")
    if not isinstance(step, dict):
        errors.append("'current_step' must be an object")
    else:
        if "step_number" not in step or not isinstance(step["step_number"], int):
            errors.append("'current_step.step_number' must be an integer")
        role = step.get("assigned_role")
        if role not in VALID_ROLES:
            errors.append(f"Invalid assigned_role '{role}'. Must be one of {sorted(VALID_ROLES)}")
        if not isinstance(step.get("target_files"), list):
            errors.append("'current_step.target_files' must be a list of file paths")

    # Validate feedback_loop
    feedback = data.get("feedback_loop")
    if not isinstance(feedback, dict):
        errors.append("'feedback_loop' must be an object")
    else:
        retry_count = feedback.get("qa_retry_count", 0)
        max_retries = feedback.get("max_retries", 3)
        if not isinstance(retry_count, int) or retry_count < 0:
            errors.append("'feedback_loop.qa_retry_count' must be a non-negative integer")
        if not isinstance(max_retries, int) or max_retries <= 0:
            errors.append("'feedback_loop.max_retries' must be a positive integer")
        if retry_count >= max_retries and status != "BLOCKED":
            errors.append(
                f"qa_retry_count ({retry_count}) >= max_retries ({max_retries}), but status is '{status}' (MUST be 'BLOCKED')"
            )

    report_and_exit(errors)

def report_and_exit(errors):
    if errors:
        print("[FAIL] State Validation Failed with errors:", file=sys.stderr)
        for err in errors:
            print(f"  - {err}", file=sys.stderr)
        sys.exit(1)
    else:
        print("[OK] State validation successful. handoff.json and STATE.md are valid.")
        sys.exit(0)

if __name__ == "__main__":
    validate()
