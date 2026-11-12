---
name: test-design
type: on-demand
description: Comprehensive test planning covering happy path, edge cases, failure injection, and concurrency.
---
# Test Design & Architecture Standard

Design deterministic, high-coverage test suites that test system behavior rather than implementation details.

## 4-Tier Test Matrix
For any new feature or contract, design tests across four critical tiers:

### 1. Happy Path Coverage
- Validate expected behavior under typical valid inputs.
- Verify return codes, payload structures, side effects, and state transitions.

### 2. Boundary & Edge Cases
- Test boundary limits: empty strings, zero length slices, max integers, special characters, unicode, leap years, timezone offsets.
- Test missing/optional parameters and malformed inputs.
- Verify clear, informative error responses rather than unhandled panics or 500 crashes.

### 3. Failure Injection & Resilience
- Inject artificial faults for external dependencies (database down, redis timeout, 3rd party API returning 503).
- Verify circuit breakers, retries with exponential backoff, and graceful fallback paths.
- Ensure no leaked connections or hanging operations when errors occur.

### 4. Concurrency & Race Conditions
- Run parallel test executions where appropriate (e.g. `go test -race`, concurrent threads).
- Test race conditions on shared states, mutexes, and cache invalidation.

## Test Implementation Principles
- **Determinism:** Never rely on wall-clock time (`time.Sleep`). Use mock clocks or event-driven synchronization.
- **Isolation:** Each test must clean up its created resources or run in isolated transactional scopes.
- **Assertion Quality:** Test domain assertions, not mock call counts.