---
name: test-design
type: on-demand
description: TDD architecture, boundary value analysis, fault injection, and concurrency testing. Use when designing robust test suites.
---
# Advanced Test Design & Verification Architecture

Design deterministic, high-coverage test suites that test domain behavior rather than shallow mock invocations. Follow this 4-tier testing hierarchy:

---

## 1. Boundary Value Analysis (BVA) & Equivalence Partitioning
For every function parameter, entity field, and API input, generate test vectors covering all 5 boundaries:

| Input Dimension | Lower Boundary | Valid Nominal | Upper Boundary | Invalid / Malformed |
|---|---|---|---|---|
| String Fields | Empty `""` | Standard string | Max length (e.g. 255) | Exceeds max, pure whitespace, SQL/XSS injections |
| Numeric / Amounts | `0`, `-1` | Typical value | Max Int / Precision limit | Overflow, non-numeric strings, negative amounts |
| Collections / Arrays | Empty `[]` | Single item `[x]` | Batch limit (e.g. 100) | Exceeds batch limit, duplicates, `nil/null` |
| Timestamps / Dates | Epoch `0` | Current UTC | Far future / Leap year | Malformed ISO-8601, past expiration |

---

## 2. Negative Scenarios & Fault-Injection Matrix
Every module communicating across network, IPC, or storage boundaries must verify safe degradation:
- **Downstream Timeout:** Mock 3rd-party API timeout ($>5000\text{ms}$) $\rightarrow$ verify circuit breaker triggers and returns structured error.
- **Connection Refused:** Simulate DB connection pool exhaustion $\rightarrow$ verify graceful fallback without unhandled process crash.
- **Corrupt / Partial Payloads:** Feed truncated JSON or invalid protobuf $\rightarrow$ verify 400 Bad Request with actionable field validation error.

---

## 3. Concurrency, Race Hazards & Idempotency
- **Race Condition Detection:** Execute tests with race detection flags enabled (`go test -race`, Node thread concurrency).
- **Concurrent Identical Requests:** Fire 10 parallel duplicate requests with identical idempotency keys $\rightarrow$ verify exactly one succeeds, 9 return cached response or safe duplicate status.
- **Deadlock Resistance:** Verify shared state locks release cleanly even when panics or early returns occur (`defer mu.Unlock()`).

---

## 4. Determinism & Test Hygiene Rules
- **Freeze Time:** Never use arbitrary wall-clock waits (`sleep(100)`). Use injected clock interfaces or event-driven sync channels.
- **Hermetic Isolation:** Every test must run in a clean transactional boundary or ephemeral test database. Zero cross-test state leakage.
- **Assertion Granularity:** Assert exact return types, status codes, and domain invariants; never merely assert that a mock was called.