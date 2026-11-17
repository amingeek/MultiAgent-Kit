---
name: performance-investigation
type: on-demand
description: Quantitative profiling, bottleneck analysis, and algorithmic optimization. Use when investigating latency, memory leaks, or throughput constraints.
---
# Quantitative Performance Profiling & Optimization

Never optimize without empirical measurements. Follow this 4-stage profiling and benchmarking protocol:

```
[1. Baseline Metrics] ──► [2. Profile Hotspots] ──► [3. Formulate Optimization] ──► [4. Benchmark Verification]
```

---

## 1. Baseline Quantification
Measure exact baseline latency, memory allocations, and throughput under identical conditions:
- **Latency Distribution:** Capture p50, p95, p99 (never optimize for averages alone).
- **Resource Footprint:** Resident Set Size (RSS memory), CPU utilization %, Garbage Collection pause times.
- **Record Baseline:** Save numbers in `agent/STATE.md` before applying code modifications.

---

## 2. Profiling Hotspots & Root Causes

### A. CPU & Memory Profiling
- **Go Profiling:**
  ```bash
  go test -bench=. -benchmem -cpuprofile=cpu.pprof -memprofile=mem.pprof
  go tool pprof -top -cum cpu.pprof | head -n 20
  ```
- **Inspect Allocations:** Target functions allocating large temporary slices in tight loops (`-benchmem`). Prefer object reuse or pre-allocated capacities (`make([]T, 0, capacity)`).

### B. Database & I/O Bottlenecks
- **Query Plan Inspection:** Execute `EXPLAIN (ANALYZE, BUFFERS) <query>`:
  - ❌ **Sequential Scan (`Seq Scan`):** Indicates missing index or non-sargable query predicates.
  - ✅ **Index Only Scan (`Index Only Scan`):** Desired state; satisfies query entirely from B-Tree index.
- **N+1 Query Detection:** Inspect logs for repeated queries in loops. Replace with batch eager loading (`WHERE id IN (...)` or `JOIN`).

### C. Resource Leak Detection
- Verify all open resources are guaranteed to close (`defer resp.Body.Close()`, `defer rows.Close()`).
- Verify bounded cache sizes: always enforce LRU eviction or TTL on in-memory caches.

---

## 3. Targeted Optimization & Benchmark Verification
1. Apply the single targeted architectural or algorithmic optimization.
2. Re-run benchmark under identical test load:
   ```bash
   go test -bench=<BenchmarkTarget> -benchmem -count=5
   ```
3. **Verification Criterion:**
   - Demonstrable improvement in p95/p99 latency or memory allocations ($>20\%$).
   - 100% of functional tests pass with zero behavioral regressions.
4. Output comparison summary table in `agent/STATE.md`.