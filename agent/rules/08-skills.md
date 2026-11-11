# 08 — Skills System Discipline

1. `always` skills: Max 4-5 skills. Must be concise (<40 lines) to avoid context overhead. Loaded automatically.
2. `on-demand` skills: Loaded ONLY when explicitly invoked via `USE SKILLS: <name>` or listed in a milestone.
3. Priority hierarchy: Core Safety Rules > Human Instructions > Skills. Skills cannot override rules.
4. Missing skills: If a requested skill does not exist, log `skill not found: <name>` and proceed cleanly.
