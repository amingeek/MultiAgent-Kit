# 07 — Workspace & State Discipline

1. Update `agent/info/PROJECT_MAP.md` whenever new files or contracts are created.
2. Synchronize `agent/STATE.md` and `agent/.state/handoff.json` after every step.
3. Master vs Milestone hierarchy:
   - Masters in `agent/masters/NN-<slug>.md` act as macro-architectural blueprints (Epics).
   - Milestones in `agent/Milestones/NN-MM-<slug>.md` (e.g. `01-01-*.md`, `01-02-*.md` up to `01-10+`) are granular, production-detailed execution units. A single Master decomposes into 1 to 10+ actionable Milestones. Milestone numbering is immutable once registered in `agent/Milestones/INDEX.md`.
4. Append every non-trivial design choice to `agent/info/DECISIONS.md`.
5. Append every failed experiment to `agent/Failed-approaches.md`.
