package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	validStatuses = map[string]bool{
		"IDLE":        true,
		"IN_PROGRESS": true,
		"BLOCKED":     true,
		"QA_FAILED":   true,
		"SUCCESS":     true,
	}

	validRoles = map[string]bool{
		"orchestrator": true,
		"architect":    true,
		"developer":    true,
		"qa":           true,
		"devops":       true,
		"security":     true,
	}
)

type Milestone struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	SourceFile string `json:"source_file"`
}

type CurrentStep struct {
	StepNumber   int      `json:"step_number"`
	Description  string   `json:"description"`
	AssignedRole string   `json:"assigned_role"`
	RoleFile     string   `json:"role_file"`
	TargetFiles  []string `json:"target_files"`
}

type FeedbackLoop struct {
	QARetryCount int      `json:"qa_retry_count"`
	MaxRetries   int      `json:"max_retries"`
	ErrorLogs    string   `json:"error_logs"`
	FailingTests []string `json:"failing_tests"`
	RequiredFix  string   `json:"required_fix"`
}

type AcceptanceChecklist struct {
	Criterion string `json:"criterion"`
	Verified  bool   `json:"verified"`
}

type HandoffData struct {
	Schema              string                `json:"$schema"`
	ProjectName         string                `json:"project_name"`
	TechStack           []string              `json:"tech_stack"`
	ActiveMilestone     Milestone             `json:"active_milestone"`
	CurrentStep         CurrentStep           `json:"current_step"`
	Status              string                `json:"status"`
	ContextSummary      string                `json:"context_summary"`
	FeedbackLoop        FeedbackLoop          `json:"feedback_loop"`
	AcceptanceChecklist []AcceptanceChecklist `json:"acceptance_checklist"`
	LastUpdated         string                `json:"last_updated"`
}

type ProjectStats struct {
	MastersCount     int
	TotalMilestones  int
	CompletedCount   int
	InProgressCount  int
	BlockedCount     int
	NotStartedCount  int
	CompletionRatio  float64
}

func findRepoRoot() (string, error) {
	// First check current working directory
	cwd, err := os.Getwd()
	if err == nil {
		if isRepoRoot(cwd) {
			return cwd, nil
		}
		// Check parent directories
		dir := cwd
		for i := 0; i < 5; i++ {
			parent := filepath.Dir(dir)
			if parent == dir {
				break
			}
			if isRepoRoot(parent) {
				return parent, nil
			}
			dir = parent
		}
	}

	// Then check directory containing the executable
	exePath, err := os.Executable()
	if err == nil {
		exeDir := filepath.Dir(exePath)
		dir := exeDir
		for i := 0; i < 5; i++ {
			if isRepoRoot(dir) {
				return dir, nil
			}
			parent := filepath.Dir(dir)
			if parent == dir {
				break
			}
			dir = parent
		}
	}

	return "", fmt.Errorf("could not locate repository root (missing agent/.state/handoff.json)")
}

func isRepoRoot(dir string) bool {
	checkPath := filepath.Join(dir, "agent", ".state", "handoff.json")
	info, err := os.Stat(checkPath)
	return err == nil && !info.IsDir()
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func validateState(root string, errors *[]string, warnings *[]string) *HandoffData {
	statePath := filepath.Join(root, "agent", "STATE.md")
	handoffPath := filepath.Join(root, "agent", ".state", "handoff.json")

	// 1. Validate STATE.md
	if !fileExists(statePath) {
		*errors = append(*errors, fmt.Sprintf("Missing STATE.md at %s", statePath))
	} else {
		contentBytes, err := os.ReadFile(statePath)
		if err != nil {
			*errors = append(*errors, fmt.Sprintf("Failed to read STATE.md: %v", err))
		} else {
			content := string(contentBytes)
			requiredHeaders := []string{
				"# Global System State",
				"## Milestones Summary",
				"## Quick Tasks Log",
				"## Execution Evidence Log",
			}
			for _, h := range requiredHeaders {
				if !strings.Contains(content, h) {
					*errors = append(*errors, fmt.Sprintf("STATE.md missing mandatory section header: '%s'", h))
				}
			}
		}
	}

	// 2. Validate handoff.json
	if !fileExists(handoffPath) {
		*errors = append(*errors, fmt.Sprintf("Missing handoff.json at %s", handoffPath))
		return nil
	}

	bytes, err := os.ReadFile(handoffPath)
	if err != nil {
		*errors = append(*errors, fmt.Sprintf("Failed to read handoff.json: %v", err))
		return nil
	}

	var rawMap map[string]interface{}
	if err := json.Unmarshal(bytes, &rawMap); err != nil {
		*errors = append(*errors, fmt.Sprintf("handoff.json contains invalid JSON syntax: %v", err))
		return nil
	}

	requiredKeys := []string{
		"project_name", "tech_stack", "active_milestone", "current_step",
		"status", "context_summary", "feedback_loop", "acceptance_checklist", "last_updated",
	}
	for _, key := range requiredKeys {
		if _, ok := rawMap[key]; !ok {
			*errors = append(*errors, fmt.Sprintf("handoff.json missing required key: '%s'", key))
		}
	}

	var data HandoffData
	if err := json.Unmarshal(bytes, &data); err != nil {
		*errors = append(*errors, fmt.Sprintf("Failed to parse typed handoff.json: %v", err))
		return nil
	}

	// Validate status
	if !validStatuses[data.Status] {
		*errors = append(*errors, fmt.Sprintf("Invalid status '%s'. Must be one of IDLE, IN_PROGRESS, BLOCKED, QA_FAILED, SUCCESS", data.Status))
	}

	// Validate active_milestone
	if data.ActiveMilestone.ID == "" {
		*errors = append(*errors, "'active_milestone.id' cannot be empty")
	}
	if data.ActiveMilestone.Title == "" {
		*errors = append(*errors, "'active_milestone.title' cannot be empty")
	}
	if data.ActiveMilestone.SourceFile == "" {
		*errors = append(*errors, "'active_milestone.source_file' cannot be empty")
	} else {
		srcPath := filepath.Join(root, data.ActiveMilestone.SourceFile)
		if !fileExists(srcPath) {
			*warnings = append(*warnings, fmt.Sprintf("active_milestone.source_file '%s' does not exist on disk", data.ActiveMilestone.SourceFile))
		}
	}

	// Validate current_step
	if !validRoles[data.CurrentStep.AssignedRole] {
		*errors = append(*errors, fmt.Sprintf("Invalid assigned_role '%s'. Must be one of orchestrator, architect, developer, qa, devops, security", data.CurrentStep.AssignedRole))
	}
	if data.CurrentStep.TargetFiles == nil {
		*errors = append(*errors, "'current_step.target_files' must be an array of file paths")
	}
	if data.CurrentStep.RoleFile != "" {
		rolePath := filepath.Join(root, data.CurrentStep.RoleFile)
		if !fileExists(rolePath) {
			*warnings = append(*warnings, fmt.Sprintf("current_step.role_file '%s' does not exist on disk", data.CurrentStep.RoleFile))
		}
	}

	// Validate feedback_loop
	if data.FeedbackLoop.QARetryCount < 0 {
		*errors = append(*errors, "'feedback_loop.qa_retry_count' must be a non-negative integer")
	}
	if data.FeedbackLoop.MaxRetries <= 0 {
		*errors = append(*errors, "'feedback_loop.max_retries' must be a positive integer")
	}
	if data.FeedbackLoop.MaxRetries > 0 && data.FeedbackLoop.QARetryCount >= data.FeedbackLoop.MaxRetries && data.Status != "BLOCKED" {
		*errors = append(*errors, fmt.Sprintf("qa_retry_count (%d) >= max_retries (%d), but status is '%s' (MUST be 'BLOCKED' to prevent looping)",
			data.FeedbackLoop.QARetryCount, data.FeedbackLoop.MaxRetries, data.Status))
	}

	return &data
}

func validateAllIntegrity(root string, errors *[]string, warnings *[]string) {
	// 1. Check rules 00 to 10
	rulesDir := filepath.Join(root, "agent", "rules")
	if !dirExists(rulesDir) {
		*errors = append(*errors, fmt.Sprintf("Missing rules directory at %s", rulesDir))
	} else {
		expectedRules := []string{
			"00-core-safety.md", "01-code-quality.md", "02-never-remove-features.md",
			"03-git-workflow.md", "04-testing.md", "05-security-secrets.md",
			"06-off-limits.md", "07-agent-workspace.md", "08-skills.md",
			"09-multi-agent-protocol.md", "10-contracts-and-conventions.md",
		}
		for _, r := range expectedRules {
			path := filepath.Join(rulesDir, r)
			if !fileExists(path) {
				*errors = append(*errors, fmt.Sprintf("Missing rule file: agent/rules/%s", r))
			}
		}
	}

	// 2. Check roles 00 to 05
	rolesDir := filepath.Join(root, "agent", "roles")
	if !dirExists(rolesDir) {
		*errors = append(*errors, fmt.Sprintf("Missing roles directory at %s", rolesDir))
	} else {
		expectedRoles := []string{
			"00-orchestrator.md", "01-architect.md", "02-developer.md",
			"03-qa.md", "04-security-reviewer.md", "05-devops.md",
		}
		for _, r := range expectedRoles {
			path := filepath.Join(rolesDir, r)
			if !fileExists(path) {
				*errors = append(*errors, fmt.Sprintf("Missing role file: agent/roles/%s", r))
			}
		}
	}

	// 3. Check skills catalog
	skillsIndex := filepath.Join(root, "agent", "skills", "INDEX.md")
	if !fileExists(skillsIndex) {
		*errors = append(*errors, fmt.Sprintf("Missing skills index at %s", skillsIndex))
	} else {
		content, err := os.ReadFile(skillsIndex)
		if err == nil {
			lines := strings.Split(string(content), "\n")
			for _, line := range lines {
				if strings.Contains(line, "|") && strings.Contains(line, "agent/skills/") {
					parts := strings.Split(line, "|")
					if len(parts) >= 4 {
						skillRel := strings.TrimSpace(parts[3])
						if strings.HasPrefix(skillRel, "agent/skills/") {
							skillPath := filepath.Join(root, skillRel)
							if !fileExists(skillPath) {
								*errors = append(*errors, fmt.Sprintf("Skill referenced in INDEX.md not found on disk: %s", skillRel))
							}
						}
					}
				}
			}
		}
	}

	// 4. Check Milestones index
	milestonesIndex := filepath.Join(root, "agent", "Milestones", "INDEX.md")
	if !fileExists(milestonesIndex) {
		*errors = append(*errors, fmt.Sprintf("Missing milestones index at %s", milestonesIndex))
	}

	// 5. Check templates
	masterTpl := filepath.Join(root, "agent", "masters", "_MASTER-TEMPLATE.md")
	if !fileExists(masterTpl) {
		*warnings = append(*warnings, "Missing master template at agent/masters/_MASTER-TEMPLATE.md")
	}

	milestoneTpl := filepath.Join(root, "agent", "Milestones", "_MILESTONE-TEMPLATE.md")
	if !fileExists(milestoneTpl) {
		*warnings = append(*warnings, "Missing milestone template at agent/Milestones/_MILESTONE-TEMPLATE.md")
	}

	// 6. Check PRD.md presence
	prdPath := filepath.Join(root, "PRD.md")
	if !fileExists(prdPath) {
		*warnings = append(*warnings, "PRD.md not found at repository root (run 'mak init-prd' to create template)")
	}
}

func calculateStats(root string) ProjectStats {
	stats := ProjectStats{}

	// Count Masters (excluding templates and docs)
	mastersDir := filepath.Join(root, "agent", "masters")
	if dirExists(mastersDir) {
		entries, err := os.ReadDir(mastersDir)
		if err == nil {
			for _, e := range entries {
				name := e.Name()
				if !e.IsDir() && strings.HasSuffix(name, ".md") &&
					!strings.HasPrefix(name, "_") &&
					name != "README.md" && name != "EXECUTE.md" && name != "00-info.md" {
					stats.MastersCount++
				}
			}
		}
	}

	// Read Milestones INDEX.md
	milestonesIndex := filepath.Join(root, "agent", "Milestones", "INDEX.md")
	if fileExists(milestonesIndex) {
		content, err := os.ReadFile(milestonesIndex)
		if err == nil {
			lines := strings.Split(string(content), "\n")
			rowRegex := regexp.MustCompile(`^\|\s*(M\d+-\d+)\s*\|`)
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if rowRegex.MatchString(line) {
					stats.TotalMilestones++
					upperLine := strings.ToUpper(line)
					if strings.Contains(upperLine, "COMPLETED") || strings.Contains(upperLine, "SUCCESS") {
						stats.CompletedCount++
					} else if strings.Contains(upperLine, "IN_PROGRESS") {
						stats.InProgressCount++
					} else if strings.Contains(upperLine, "BLOCKED") {
						stats.BlockedCount++
					} else {
						stats.NotStartedCount++
					}
				}
			}
		}
	}

	if stats.TotalMilestones > 0 {
		stats.CompletionRatio = float64(stats.CompletedCount) / float64(stats.TotalMilestones) * 100.0
	}

	return stats
}

func printStats(stats ProjectStats) {
	boldColor := "\033[1m"
	resetColor := "\033[0m"
	greenColor := "\033[92m"
	cyanColor := "\033[96m"
	yellowColor := "\033[93m"
	redColor := "\033[91m"

	// Progress bar
	barLength := 24
	completedBlocks := int((stats.CompletionRatio / 100.0) * float64(barLength))
	if completedBlocks > barLength {
		completedBlocks = barLength
	}
	progressBar := strings.Repeat("█", completedBlocks) + strings.Repeat("░", barLength-completedBlocks)

	fmt.Printf("\n%s═══════════════════════════════════════════════════════════%s\n", boldColor, resetColor)
	fmt.Printf("%s 📊 MultiAgent-Kit Project Execution Metrics (v6.0)%s\n", boldColor, resetColor)
	fmt.Printf("%s═══════════════════════════════════════════════════════════%s\n", boldColor, resetColor)
	fmt.Printf("  %sDomain Masters (Epics):%s   %d\n", boldColor, resetColor, stats.MastersCount)
	fmt.Printf("  %sTotal Milestones:%s         %d\n", boldColor, resetColor, stats.TotalMilestones)
	fmt.Printf("  %sCompleted:%s                %s%d%s\n", boldColor, resetColor, greenColor, stats.CompletedCount, resetColor)
	fmt.Printf("  %sIn Progress:%s              %s%d%s\n", boldColor, resetColor, cyanColor, stats.InProgressCount, resetColor)
	fmt.Printf("  %sBlocked:%s                  %s%d%s\n", boldColor, resetColor, redColor, stats.BlockedCount, resetColor)
	fmt.Printf("  %sNot Started:%s              %s%d%s\n", boldColor, resetColor, yellowColor, stats.NotStartedCount, resetColor)
	fmt.Printf("  %sProgress Bar:%s             [%s%s%s] %.1f%%\n", boldColor, resetColor, greenColor, progressBar, resetColor, stats.CompletionRatio)
	fmt.Printf("%s═══════════════════════════════════════════════════════════%s\n\n", boldColor, resetColor)
}

func printSummary(data *HandoffData) {
	if data == nil {
		fmt.Println("Cannot show summary: handoff.json could not be loaded.")
		return
	}

	statusColor := "\033[96m" // Cyan
	if data.Status == "SUCCESS" {
		statusColor = "\033[92m" // Green
	} else if data.Status == "BLOCKED" || data.Status == "QA_FAILED" {
		statusColor = "\033[91m" // Red
	}
	resetColor := "\033[0m"
	boldColor := "\033[1m"
	yellowColor := "\033[93m"

	tech := strings.Join(data.TechStack, ", ")
	if tech == "" {
		tech = "Polyglot / Universal"
	}

	fmt.Printf("\n%s═══════════════════════════════════════════════════════════%s\n", boldColor, resetColor)
	fmt.Printf("%s 🤖 MultiAgent-Kit System Dashboard (v6.0 Go Native)%s\n", boldColor, resetColor)
	fmt.Printf("%s═══════════════════════════════════════════════════════════%s\n", boldColor, resetColor)
	fmt.Printf("  %sProject Name:%s      %s\n", boldColor, resetColor, data.ProjectName)
	fmt.Printf("  %sTech Stack:%s        %s\n", boldColor, resetColor, tech)
	fmt.Printf("  %sSystem Status:%s     %s%s%s\n", boldColor, resetColor, statusColor, data.Status, resetColor)
	fmt.Printf("  %sActive Milestone:%s  [%s] %s\n", boldColor, resetColor, data.ActiveMilestone.ID, data.ActiveMilestone.Title)
	fmt.Printf("  %sMilestone Source:%s  %s\n", boldColor, resetColor, data.ActiveMilestone.SourceFile)
	fmt.Printf("  %sCurrent Step:%s      #%d: %s\n", boldColor, resetColor, data.CurrentStep.StepNumber, data.CurrentStep.Description)
	fmt.Printf("  %sAssigned Role:%s     %s%s%s (%s)\n", boldColor, resetColor, yellowColor, data.CurrentStep.AssignedRole, resetColor, data.CurrentStep.RoleFile)
	fmt.Printf("  %sQA Retry Count:%s    %d / %d\n", boldColor, resetColor, data.FeedbackLoop.QARetryCount, data.FeedbackLoop.MaxRetries)
	fmt.Printf("  %sLast Updated:%s      %s\n", boldColor, resetColor, data.LastUpdated)
	fmt.Printf("%s═══════════════════════════════════════════════════════════%s\n\n", boldColor, resetColor)
}

func touchTimestamps(root string) error {
	nowISO := time.Now().UTC().Format(time.RFC3339)
	handoffPath := filepath.Join(root, "agent", ".state", "handoff.json")
	statePath := filepath.Join(root, "agent", "STATE.md")

	if fileExists(handoffPath) {
		bytes, err := os.ReadFile(handoffPath)
		if err == nil {
			var rawMap map[string]interface{}
			if err := json.Unmarshal(bytes, &rawMap); err == nil {
				rawMap["last_updated"] = nowISO
				outBytes, err := json.MarshalIndent(rawMap, "", "  ")
				if err == nil {
					_ = os.WriteFile(handoffPath, append(outBytes, '\n'), 0644)
					fmt.Printf("[OK] Updated handoff.json last_updated to %s\n", nowISO)
				}
			}
		}
	}

	if fileExists(statePath) {
		bytes, err := os.ReadFile(statePath)
		if err == nil {
			lines := strings.Split(string(bytes), "\n")
			var newLines []string
			for _, l := range lines {
				if strings.HasPrefix(l, "- Last update:") {
					newLines = append(newLines, fmt.Sprintf("- Last update: %s", nowISO))
				} else {
					newLines = append(newLines, l)
				}
			}
			_ = os.WriteFile(statePath, []byte(strings.Join(newLines, "\n")), 0644)
			fmt.Printf("[OK] Updated STATE.md Last update to %s\n", nowISO)
		}
	}

	return nil
}

func resetRetryCount(root string) error {
	handoffPath := filepath.Join(root, "agent", ".state", "handoff.json")
	if fileExists(handoffPath) {
		bytes, err := os.ReadFile(handoffPath)
		if err != nil {
			return err
		}
		var rawMap map[string]interface{}
		if err := json.Unmarshal(bytes, &rawMap); err != nil {
			return err
		}

		if fb, ok := rawMap["feedback_loop"].(map[string]interface{}); ok {
			fb["qa_retry_count"] = 0
			fb["error_logs"] = ""
			fb["failing_tests"] = []string{}
			fb["required_fix"] = ""
		}

		status, _ := rawMap["status"].(string)
		if status == "QA_FAILED" || status == "BLOCKED" {
			rawMap["status"] = "IN_PROGRESS"
		}

		outBytes, err := json.MarshalIndent(rawMap, "", "  ")
		if err != nil {
			return err
		}
		if err := os.WriteFile(handoffPath, append(outBytes, '\n'), 0644); err != nil {
			return err
		}
		fmt.Println("[OK] Successfully reset feedback_loop.qa_retry_count to 0 and set status to IN_PROGRESS.")
	}
	return nil
}

func initPRDTemplate(root string) error {
	prdPath := filepath.Join(root, "PRD.md")
	if fileExists(prdPath) {
		return fmt.Errorf("PRD.md already exists at %s", prdPath)
	}

	defaultPRD := `# Product Requirements Document (PRD)

> **Instructions for the Operator:**
> Fill in the sections below to describe your product idea, architecture expectations, and feature domains.
> When ready, run in your AI CLI or terminal:
> ` + "`" + `Read BOOTSTRAP.md and follow it exactly.` + "`" + `
> MultiAgent-Kit will bootstrap the project structure and generate all Domain Masters (` + "`" + `agent/masters/NN-*.md` + "`" + `).

---

## 1. Executive Summary & Vision
- **Product Name:** <Enter product name>
- **One-Sentence Description:** <High-level elevator pitch>
- **Target Audience / End Users:** <Who uses this product?>
- **Core Value Proposition:** <Why does this product exist?>

---

## 2. Technical Stack & Architecture Preferences
- **Preferred Languages:** <e.g., Go, TypeScript, Rust, Python, etc.>
- **Frameworks & Libraries:** <e.g., Gin, FastAPI, React, Next.js, etc.>
- **Database & Persistence:** <e.g., PostgreSQL, SQLite, Redis, MongoDB>
- **Communication Protocols:** <e.g., REST, GraphQL, gRPC, WebSocket>
- **Infrastructure / Deployment:** <e.g., Docker, Kubernetes, Serverless, Bare Metal>

---

## 3. High-Level Feature Domains (Masters / Epics)
*List all major domains of your application. Each domain will be translated into a dedicated Master Epic (` + "`" + `agent/masters/01-*.md` + "`" + `, ` + "`" + `02-*.md` + "`" + `, etc.) during bootstrap. There is no limit on the number of domains!*

### Domain 1: <e.g. Identity & Access Management (IAM)>
- **Summary:** User registration, authentication, authorization, JWT sessions, MFA.
- **Key Capabilities:**
  - Secure signup & login
  - Token refresh and revocation
  - Role-based access control (RBAC)

### Domain 2: <e.g. Core Business Domain / Service Logic>
- **Summary:** Primary application logic and workflows.
- **Key Capabilities:**
  - CRUD operations on core entities
  - Business rule validations
  - Transaction integrity

### Domain 3: <e.g. Data Ingestion & Storage / API Gateway>
- **Summary:** External integrations, background workers, or event-driven pipelines.
- **Key Capabilities:**
  - Async jobs / queues
  - Rate limiting & caching
  - Webhooks & external notifications

---

## 4. Non-Functional Requirements & Security
- **Performance Budget:** <e.g. p99 response time < 100ms, concurrency targets>
- **Security & Compliance:** <e.g. OWASP Top 10 compliance, encryption at rest and in transit>
- **Availability & Resilience:** <e.g. 99.9% uptime, graceful degradation on downstream failure>
- **Test Coverage Target:** <e.g. >= 80% coverage on domain logic with automated TDD suites>

---

## 5. Scope Boundaries & Constraints
- **In-Scope (v1.0 Launch):**
  - Features strictly required for MVP / v1.0
- **Out-of-Scope (Future Iterations):**
  - Features intentionally deferred to later versions
- **Off-Limits Directories / Legacy Code:**
  - Any files or folders AI agents must never modify
`
	if err := os.WriteFile(prdPath, []byte(defaultPRD), 0644); err != nil {
		return err
	}

	fmt.Printf("[OK] Initialized PRD.md template at %s\n", prdPath)
	return nil
}

func main() {
	var checkAll bool
	var summary bool
	var touch bool
	var resetRetry bool
	var showStats bool
	var initPRD bool

	flag.BoolVar(&checkAll, "check-all", false, "Perform deep audit of all rules, skills, templates, PRD, and state")
	flag.BoolVar(&checkAll, "a", false, "Alias for --check-all")
	flag.BoolVar(&summary, "summary", false, "Display visual execution dashboard")
	flag.BoolVar(&summary, "s", false, "Alias for --summary")
	flag.BoolVar(&touch, "touch", false, "Synchronize last_updated timestamps in state files")
	flag.BoolVar(&touch, "t", false, "Alias for --touch")
	flag.BoolVar(&resetRetry, "reset-retry", false, "Reset QA retry count to 0 and unblock")
	flag.BoolVar(&resetRetry, "r", false, "Alias for --reset-retry")
	flag.BoolVar(&showStats, "stats", false, "Display project execution stats, milestone counts, and progress bar")
	flag.BoolVar(&initPRD, "init-prd", false, "Create default PRD.md template if not present")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "MultiAgent-Kit CLI & State Validator (v6.0 Go Native)\n\nUsage:\n")
		fmt.Fprintf(os.Stderr, "  mak [command | flag]\n\nCommands:\n")
		fmt.Fprintf(os.Stderr, "  validate              # Validate state & handoff files (default)\n")
		fmt.Fprintf(os.Stderr, "  check-all, -a         # Deep audit: rules, roles, skills, templates, PRD, state\n")
		fmt.Fprintf(os.Stderr, "  summary, status, -s   # Print visual execution dashboard\n")
		fmt.Fprintf(os.Stderr, "  stats                 # Print milestone breakdown & progress bar\n")
		fmt.Fprintf(os.Stderr, "  touch, -t             # Sync timestamps to current UTC\n")
		fmt.Fprintf(os.Stderr, "  reset-retry, -r       # Reset QA retry count to 0 and unblock\n")
		fmt.Fprintf(os.Stderr, "  init-prd              # Initialize PRD.md template\n\nFlags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Also support positional subcommands like `mak summary`, `mak stats`, etc.
	posArgs := flag.Args()
	if len(posArgs) > 0 {
		switch strings.ToLower(posArgs[0]) {
		case "check-all", "check", "audit":
			checkAll = true
		case "summary", "status", "dashboard":
			summary = true
		case "stats", "metrics", "progress":
			showStats = true
		case "touch", "sync":
			touch = true
		case "reset-retry", "reset", "unblock":
			resetRetry = true
		case "init-prd", "prd":
			initPRD = true
		case "validate":
			// Default validation mode
		case "help":
			flag.Usage()
			return
		default:
			fmt.Fprintf(os.Stderr, "Unknown subcommand: %s (run 'mak help' for usage)\n", posArgs[0])
			os.Exit(1)
		}
	}

	root, err := findRepoRoot()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[FAIL] %v\n", err)
		os.Exit(1)
	}

	if initPRD {
		if err := initPRDTemplate(root); err != nil {
			fmt.Fprintf(os.Stderr, "[FAIL] %v\n", err)
			os.Exit(1)
		}
		return
	}

	if touch {
		if err := touchTimestamps(root); err != nil {
			fmt.Fprintf(os.Stderr, "[FAIL] Touch failed: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if resetRetry {
		if err := resetRetryCount(root); err != nil {
			fmt.Fprintf(os.Stderr, "[FAIL] Reset retry failed: %v\n", err)
			os.Exit(1)
		}
		return
	}

	var errors []string
	var warnings []string

	data := validateState(root, &errors, &warnings)

	if checkAll {
		validateAllIntegrity(root, &errors, &warnings)
	}

	if summary {
		printSummary(data)
	}

	if showStats {
		stats := calculateStats(root)
		printStats(stats)
	}

	for _, w := range warnings {
		fmt.Fprintf(os.Stderr, "\033[93m[WARN]\033[0m %s\n", w)
	}

	if len(errors) > 0 {
		fmt.Fprintf(os.Stderr, "\n\033[91m[FAIL] State Validation Failed (%d error(s)):\033[0m\n", len(errors))
		for _, e := range errors {
			fmt.Fprintf(os.Stderr, "  \033[91m✖\033[0m %s\n", e)
		}
		os.Exit(1)
	}

	if !summary && !showStats {
		fmt.Println("\033[92m[OK]\033[0m State validation successful. handoff.json and STATE.md are valid.")
		if checkAll {
			fmt.Println("\033[92m[OK]\033[0m Full kit integrity check passed: rules, roles, skills, and templates verified.")
		}
	}
}
