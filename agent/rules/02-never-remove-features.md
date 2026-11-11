# 02 — Never Remove Features

1. Do not delete, disable, or comment out existing endpoints, handlers, CLI flags, or data schemas.
2. Never delete or weaken a test to make CI pass. A failing test indicates defect in new implementation.
3. Deleting files is only allowed if created within the current active milestone.
4. Feature flags may conceal new behavior; they must never disable existing functionality by default.
5. The passed test count must always be >= the baseline test count with 0 new failures.
