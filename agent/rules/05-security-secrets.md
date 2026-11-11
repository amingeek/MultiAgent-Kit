# 05 — Security & Secrets

1. Zero credential leaks: Never commit passwords, tokens, API keys. Use `.env.example` with dummy values.
2. Input sanitization: Prevent SQL injection, Command injection, Path traversal. Use parameterized queries.
3. Server-side authorization: Verify identities and role permissions on backend for every protected resource.
4. Least privilege: Service ports, container users, and DB users must run with minimal permissions.
