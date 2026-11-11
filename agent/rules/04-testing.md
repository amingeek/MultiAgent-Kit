# 04 — Testing & TDD Standard

1. Baseline recording: Run full test suite BEFORE touching code. Record counts in `STATE.md`.
2. Test behavior, not implementation details: Cover happy path, edge cases, error conditions, and concurrency.
3. Deterministic execution: No arbitrary `sleep()`, mock external networks, freeze time where necessary.
4. Failure-injection: Every external service dependency must have a test verifying safe degradation when down.
5. Evidence requirement: Checkboxes in milestones are only ticked after verified CLI test command execution output.
