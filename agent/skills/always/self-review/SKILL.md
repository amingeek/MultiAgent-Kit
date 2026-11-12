---
name: self-review
type: always
description: Pre-commit sanity verification to ensure safety, quality, and clean git staging.
---
# Pre-Commit Self-Review

Before staging files or updating state to `SUCCESS`, execute this mandatory sanity check:

## Verification Checklist
1. **Diff Audit:**
   - Run `git diff` on modified files.
   - Verify every line changed is directly related to the current task. No accidental reformatting, unintended edits, or debug prints (`console.log`, `fmt.Println`, `print()`).
2. **Secret & Credential Check:**
   - Confirm zero API keys, passwords, bearer tokens, or PEM secrets are present in modified files or git status.
   - Check `.env` is ignored and only `.env.example` has placeholder keys.
3. **Build & Test Regression:**
   - Run the project's build command (`make build`, `npm run build`, `go build ./...`, etc.).
   - Run the test suite: verified test pass count must be $\ge$ baseline count with 0 new failures.
4. **Contract Integrity:**
   - Confirm no public API signatures, protobuf tags, JSON keys, or DB schemas were broken or deleted without deprecation.