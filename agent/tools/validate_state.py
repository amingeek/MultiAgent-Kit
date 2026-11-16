#!/usr/bin/env python3
"""
MultiAgent-Kit State & Integrity Validator (v5.3 Production-Ready)
Validates agent/.state/handoff.json, agent/STATE.md, rules, skills, and milestone integrity.
Zero external dependencies (pure Python 3.8+ standard library).

Usage:
  python3 agent/tools/validate_state.py              # Validate state & handoff files
  python3 agent/tools/validate_state.py --check-all  # Deep audit: rules, skills, templates, state
  python3 agent/tools/validate_state.py --summary    # Print visual execution dashboard
  python3 agent/tools/validate_state.py --touch      # Sync timestamps to current UTC
  python3 agent/tools/validate_state.py --reset-retry# Reset QA retry count to 0
"""

import json
import sys
import os
import argparse
from datetime import datetime, timezone
from pathlib import Path

VALID_STATUSES = {"IDLE", "IN_PROGRESS", "BLOCKED", "QA_FAILED", "SUCCESS"}
VALID_ROLES = {"orchestrator", "architect", "developer", "qa", "devops", "security"}

# ANSI Color codes with terminal detection
USE_COLOR = sys.stdout.isatty() and os.name != "nt"
CLR_GREEN = "\033[92m" if USE_COLOR else ""
CLR_RED = "\033[91m" if USE_COLOR else ""
CLR_YELLOW = "\033[93m" if USE_COLOR else ""
CLR_CYAN = "\033[96m" if USE_COLOR else ""
CLR_BOLD = "\033[1m" if USE_COLOR else ""
CLR_RESET = "\033[0m" if USE_COLOR else ""

def get_repo_root() -> Path:
    return Path(__file__).resolve().parent.parent.parent

def validate_state(root: Path, errors: list, warnings: list):
    handoff_path = root / "agent" / ".state" / "handoff.json"
    state_path = root / "agent" / "STATE.md"

    # 1. Validate STATE.md
    if not state_path.is_file():
        errors.append(f"Missing STATE.md at {state_path}")
    else:
        content = state_path.read_text(encoding="utf-8")
        required_headers = [
            "# Global System State",
            "## Milestones Summary",
            "## Quick Tasks Log",
            "## Execution Evidence Log"
        ]
        for h in required_headers:
            if h not in content:
                errors.append(f"STATE.md missing mandatory section header: '{h}'")

    # 2. Validate handoff.json
    if not handoff_path.is_file():
        errors.append(f"Missing handoff.json at {handoff_path}")
        return None

    try:
        data = json.loads(handoff_path.read_text(encoding="utf-8"))
    except json.JSONDecodeError as e:
        errors.append(f"handoff.json contains invalid JSON syntax: {e}")
        return None

    required_keys = [
        "project_name", "tech_stack", "active_milestone", "current_step",
        "status", "context_summary", "feedback_loop", "acceptance_checklist", "last_updated"
    ]
    for key in required_keys:
        if key not in data:
            errors.append(f"handoff.json missing required key: '{key}'")

    # Status validation
    status = data.get("status")
    if status not in VALID_STATUSES:
        errors.append(f"Invalid status '{status}'. Must be one of {sorted(VALID_STATUSES)}")

    # Active milestone validation
    milestone = data.get("active_milestone")
    if not isinstance(milestone, dict):
        errors.append("'active_milestone' must be an object")
    else:
        for m_key in ["id", "title", "source_file"]:
            if m_key not in milestone or not milestone[m_key]:
                errors.append(f"'active_milestone' missing or empty field: '{m_key}'")
        if isinstance(milestone.get("source_file"), str):
            src_path = root / milestone["source_file"]
            if not src_path.is_file():
                warnings.append(f"active_milestone.source_file '{milestone['source_file']}' does not exist on disk")

    # Current step validation
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
        role_file = step.get("role_file")
        if role_file and not (root / role_file).is_file():
            warnings.append(f"current_step.role_file '{role_file}' does not exist on disk")

    # Feedback loop validation
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
                f"qa_retry_count ({retry_count}) >= max_retries ({max_retries}), but status is '{status}' (MUST be 'BLOCKED' to prevent looping)"
            )

    return data

def validate_all_integrity(root: Path, errors: list, warnings: list):
    # Check rules 00 to 10
    rules_dir = root / "agent" / "rules"
    if not rules_dir.is_dir():
        errors.append(f"Missing rules directory at {rules_dir}")
    else:
        expected_rules = [
            "00-core-safety.md", "01-code-quality.md", "02-never-remove-features.md",
            "03-git-workflow.md", "04-testing.md", "05-security-secrets.md",
            "06-off-limits.md", "07-agent-workspace.md", "08-skills.md",
            "09-multi-agent-protocol.md", "10-contracts-and-conventions.md"
        ]
        for r in expected_rules:
            if not (rules_dir / r).is_file():
                errors.append(f"Missing rule file: agent/rules/{r}")

    # Check skills catalog
    skills_index = root / "agent" / "skills" / "INDEX.md"
    if not skills_index.is_file():
        errors.append(f"Missing skills index at {skills_index}")
    else:
        for line in skills_index.read_text(encoding="utf-8").splitlines():
            if "|" in line and "agent/skills/" in line:
                parts = [p.strip() for p in line.split("|")]
                if len(parts) >= 4 and parts[3].startswith("agent/skills/"):
                    skill_path = root / parts[3]
                    if not skill_path.is_file():
                        errors.append(f"Skill referenced in INDEX.md not found on disk: {parts[3]}")

    # Check Milestones index
    milestones_index = root / "agent" / "Milestones" / "INDEX.md"
    if not milestones_index.is_file():
        errors.append(f"Missing milestones index at {milestones_index}")

    # Check templates
    master_tpl = root / "agent" / "masters" / "_MASTER-TEMPLATE.md"
    if not master_tpl.is_file():
        warnings.append("Missing master template at agent/masters/_MASTER-TEMPLATE.md")

    milestone_tpl = root / "agent" / "Milestones" / "_MILESTONE-TEMPLATE.md"
    if not milestone_tpl.is_file():
        warnings.append("Missing milestone template at agent/Milestones/_MILESTONE-TEMPLATE.md")

def print_summary(root: Path, data: dict):
    if not data:
        print(f"{CLR_RED}Cannot show summary: handoff.json could not be loaded.{CLR_RESET}")
        return

    m = data.get("active_milestone", {})
    step = data.get("current_step", {})
    fb = data.get("feedback_loop", {})
    status = data.get("status", "UNKNOWN")

    status_color = CLR_GREEN if status == "SUCCESS" else (CLR_RED if status in ("BLOCKED", "QA_FAILED") else CLR_CYAN)

    print(f"\n{CLR_BOLD}═══════════════════════════════════════════════════════════{CLR_RESET}")
    print(f"{CLR_BOLD} 🤖 MultiAgent-Kit System Dashboard (v5.3 Production){CLR_RESET}")
    print(f"{CLR_BOLD}═══════════════════════════════════════════════════════════{CLR_RESET}")
    print(f"  {CLR_BOLD}Project Name:{CLR_RESET}      {data.get('project_name', 'N/A')}")
    print(f"  {CLR_BOLD}Tech Stack:{CLR_RESET}        {', '.join(data.get('tech_stack', [])) or 'Polyglot / Universal'}")
    print(f"  {CLR_BOLD}System Status:{CLR_RESET}     {status_color}{status}{CLR_RESET}")
    print(f"  {CLR_BOLD}Active Milestone:{CLR_RESET}  [{m.get('id', 'N/A')}] {m.get('title', 'N/A')}")
    print(f"  {CLR_BOLD}Milestone Source:{CLR_RESET}  {m.get('source_file', 'N/A')}")
    print(f"  {CLR_BOLD}Current Step:{CLR_RESET}      #{step.get('step_number', 0)}: {step.get('description', 'N/A')}")
    print(f"  {CLR_BOLD}Assigned Role:{CLR_RESET}     {CLR_YELLOW}{step.get('assigned_role', 'N/A')}{CLR_RESET} ({step.get('role_file', 'N/A')})")
    print(f"  {CLR_BOLD}QA Retry Count:{CLR_RESET}    {fb.get('qa_retry_count', 0)} / {fb.get('max_retries', 3)}")
    print(f"  {CLR_BOLD}Last Updated:{CLR_RESET}      {data.get('last_updated', 'N/A')}")
    print(f"{CLR_BOLD}═══════════════════════════════════════════════════════════{CLR_RESET}\n")

def touch_timestamps(root: Path):
    now_iso = datetime.now(timezone.utc).strftime("%Y-%m-%dT%H:%M:%SZ")
    handoff_path = root / "agent" / ".state" / "handoff.json"
    state_path = root / "agent" / "STATE.md"

    if handoff_path.is_file():
        data = json.loads(handoff_path.read_text(encoding="utf-8"))
        data["last_updated"] = now_iso
        handoff_path.write_text(json.dumps(data, indent=2) + "\n", encoding="utf-8")
        print(f"{CLR_GREEN}[OK]{CLR_RESET} Updated handoff.json last_updated to {now_iso}")

    if state_path.is_file():
        lines = state_path.read_text(encoding="utf-8").splitlines()
        new_lines = []
        for l in lines:
            if l.startswith("- Last update:"):
                new_lines.append(f"- Last update: {now_iso}")
            else:
                new_lines.append(l)
        state_path.write_text("\n".join(new_lines) + "\n", encoding="utf-8")
        print(f"{CLR_GREEN}[OK]{CLR_RESET} Updated STATE.md Last update to {now_iso}")

def reset_retry_count(root: Path):
    handoff_path = root / "agent" / ".state" / "handoff.json"
    if handoff_path.is_file():
        data = json.loads(handoff_path.read_text(encoding="utf-8"))
        if "feedback_loop" in data:
            data["feedback_loop"]["qa_retry_count"] = 0
            data["feedback_loop"]["error_logs"] = ""
            data["feedback_loop"]["failing_tests"] = []
            data["feedback_loop"]["required_fix"] = ""
        if data.get("status") in ("QA_FAILED", "BLOCKED"):
            data["status"] = "IN_PROGRESS"
        handoff_path.write_text(json.dumps(data, indent=2) + "\n", encoding="utf-8")
        print(f"{CLR_GREEN}[OK]{CLR_RESET} Successfully reset feedback_loop.qa_retry_count to 0 and set status to IN_PROGRESS.")

def main():
    parser = argparse.ArgumentParser(description="MultiAgent-Kit State and Integrity Validator")
    parser.add_argument("--check-all", "-a", action="store_true", help="Perform deep audit of all rules, skills, and templates")
    parser.add_argument("--summary", "-s", action="store_true", help="Display visual execution dashboard")
    parser.add_argument("--touch", "-t", action="store_true", help="Synchronize last_updated timestamps in state files")
    parser.add_argument("--reset-retry", "-r", action="store_true", help="Reset QA retry count to 0 and unblock")
    args = parser.parse_args()

    root = get_repo_root()

    # Prefer native compiled Go binary if present
    go_bin = root / "agent" / "tools" / "validate_state"
    if go_bin.is_file() and os.access(go_bin, os.X_OK) and os.environ.get("FORCE_PYTHON_VALIDATOR") != "1":
        import subprocess
        result = subprocess.run([str(go_bin)] + sys.argv[1:])
        sys.exit(result.returncode)

    if args.touch:
        touch_timestamps(root)
        return

    if args.reset_retry:
        reset_retry_count(root)
        return

    errors = []
    warnings = []

    data = validate_state(root, errors, warnings)

    if args.check_all:
        validate_all_integrity(root, errors, warnings)

    if args.summary:
        print_summary(root, data)

    # Output warnings if any
    for w in warnings:
        print(f"{CLR_YELLOW}[WARN]{CLR_RESET} {w}", file=sys.stderr)

    # Output errors or success
    if errors:
        print(f"\n{CLR_RED}[FAIL] State Validation Failed ({len(errors)} error(s)):{CLR_RESET}", file=sys.stderr)
        for err in errors:
            print(f"  {CLR_RED}✖{CLR_RESET} {err}", file=sys.stderr)
        sys.exit(1)
    else:
        if not args.summary:
            print(f"{CLR_GREEN}[OK]{CLR_RESET} State validation successful. handoff.json and STATE.md are valid.")
            if args.check_all:
                print(f"{CLR_GREEN}[OK]{CLR_RESET} Full kit integrity check passed: rules, skills, and templates verified.")
        sys.exit(0)

if __name__ == "__main__":
    main()
