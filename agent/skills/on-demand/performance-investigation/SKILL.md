---
name: performance-investigation
type: on-demand
description: Baseline measurement, bottleneck identification, and targeted optimization.
---
# Performance Investigation Protocol

Never optimize prematurely. Always measure first, identify the true constraint, and verify improvements with benchmarks.

## Investigation Methodology

### 1. Establish Baselines
- Measure baseline latency (p50, p95, p99), throughput (RPS), and resource consumption (CPU, Memory, IOPS).
- Record exact measurement conditions to guarantee repeatability.

### 2. Profile & Identify Bottlenecks
- **CPU Profiling:** Locate hot functions and unnecessary allocations (e.g. `pprof`, `py-spy`, Chrome DevTools, flame graphs).
- **Database & I/O:**
  - Inspect query execution plans (`EXPLAIN ANALYZE`). Check for missing indexes or sequential table scans.
  - Detect N+1 query patterns. Batch database requests where applicable.
- **Memory & Leaks:** Identify unbounded caches, unclosed file descriptors, goroutine/thread leaks, or circular references.

### 3. Formulate Optimization Hypothesis
- Identify the single dominant bottleneck.
- State expected outcome: "Switching from algorithm $O(N^2)$ to $O(N \log N)$ or indexing column `user_id` will reduce p99 latency by >50%."

### 4. Implement & Benchmark
- Implement the targeted optimization without changing existing API contracts.
- Re-run benchmark tests under identical load conditions.
- Verify:
  - Latency/throughput improvement confirmed.
  - Zero regression in accuracy or test suite.
- Document metrics before and after in `agent/STATE.md` and `agent/info/DECISIONS.md`.