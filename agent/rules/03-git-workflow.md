# 03 — Git Workflow & History

1. Atomic commits: Exactly one commit per completed step (`M<NN> step <X>: <description>`).
2. Never push or merge: Only the human operator pushes to remote origins.
3. Forbidden commands: `git push --force`, `git reset --hard`, `git clean -fd`, `git checkout .`.
4. Clean staging: Use `git add <specific_file>`, never `git add .` or `git add -A`.
5. Secrets prevention: Never stage `.env`, credentials, PEM keys, tokens, or build artifacts.
