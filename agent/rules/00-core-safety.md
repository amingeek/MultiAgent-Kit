# 00 — Core Safety (Highest Authority)

1. Never break the build, test suite, or existing runtime behavior.
2. Read real code before editing. Never guess filenames, functions, types, or signatures.
3. Make the smallest viable additive change. No drive-by refactorings or cosmetic cleanups.
4. Additive contracts only: Never rename or delete public API endpoints, protobuf fields, JSON keys, DB columns, or environment variables.
5. If breaking change is required: STOP immediately, mark milestone BLOCKED in `STATE.md`, report to operator.
6. New cross-cutting concerns (logging, auth) must have safe defaults and fail-safe designs.
7. Never ask operator interactive questions during execution. Choose safest additive path, record in `agent/info/DECISIONS.md`.
8. Tooling isolation: `agent/` folder must NEVER be imported by production code or shipped in Docker runtime images.
9. High-Risk Gate: Under no circumstances execute destructive commands autonomously:
   - `rm -rf`, `DROP TABLE`, destructive migrations, `git push --force`, `git reset --hard`, production deploys.
   - Mark as `BLOCKED-ON-HUMAN` with exact manual command provided.
