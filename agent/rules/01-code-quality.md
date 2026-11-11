# 01 — Code Quality & Architecture

1. Follow repository patterns: match naming, directory structure, layering, and idiom of adjacent files.
2. Explicit error handling: Never ignore errors, never swallow exceptions silently, never return mock success.
3. Resource hygiene: Always release resources (DB connections, file handles, mutex locks, channels, timers).
4. Boundary validation: Sanitize and validate all external inputs (HTTP, RPC, CLI, files) at entry points.
5. Concurrency safety: Prevent race conditions, deadlocks, and goroutine/thread leaks. Use timeouts and contexts.
6. Zero dead code: No commented-out blocks, no unused imports, no temporary debug prints.
7. Dependency control: Do not add third-party dependencies unless strictly mandated.
