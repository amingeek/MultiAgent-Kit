package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestRepoRootAndFileExists(t *testing.T) {
	root, err := findRepoRoot()
	if err != nil {
		t.Fatalf("Expected to find repo root, got error: %v", err)
	}

	if !isRepoRoot(root) {
		t.Fatalf("Expected %s to be recognized as repo root", root)
	}

	stateFile := filepath.Join(root, "agent", "STATE.md")
	if !fileExists(stateFile) {
		t.Errorf("Expected STATE.md to exist at %s", stateFile)
	}

	rulesDir := filepath.Join(root, "agent", "rules")
	if !dirExists(rulesDir) {
		t.Errorf("Expected rules dir to exist at %s", rulesDir)
	}
}

func TestValidateStateCurrentWorkspace(t *testing.T) {
	root, err := findRepoRoot()
	if err != nil {
		t.Fatalf("Failed to locate repo root: %v", err)
	}

	var errors []string
	var warnings []string

	data := validateState(root, &errors, &warnings)
	if len(errors) > 0 {
		t.Errorf("Expected 0 errors in current workspace state, got %d: %v", len(errors), errors)
	}
	if data == nil {
		t.Fatal("Expected parsed HandoffData, got nil")
	}

	if data.Status != "IDLE" && data.Status != "IN_PROGRESS" && data.Status != "SUCCESS" {
		t.Errorf("Unexpected status: %s", data.Status)
	}

	if !validRoles[data.CurrentStep.AssignedRole] {
		t.Errorf("Invalid assigned role: %s", data.CurrentStep.AssignedRole)
	}
}

func TestValidateAllIntegrity(t *testing.T) {
	root, err := findRepoRoot()
	if err != nil {
		t.Fatalf("Failed to locate repo root: %v", err)
	}

	var errors []string
	var warnings []string

	validateAllIntegrity(root, &errors, &warnings)
	if len(errors) > 0 {
		t.Errorf("Expected 0 integrity errors, got %d: %v", len(errors), errors)
	}
}

func TestCalculateStats(t *testing.T) {
	root, err := findRepoRoot()
	if err != nil {
		t.Fatalf("Failed to locate repo root: %v", err)
	}

	stats := calculateStats(root)
	if stats.CompletionRatio < 0.0 || stats.CompletionRatio > 100.0 {
		t.Errorf("Invalid completion ratio: %f", stats.CompletionRatio)
	}
}

func TestTouchTimestamps(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "mak-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	agentStateDir := filepath.Join(tempDir, "agent", ".state")
	if err := os.MkdirAll(agentStateDir, 0755); err != nil {
		t.Fatalf("Failed to create agent/.state: %v", err)
	}

	handoffPath := filepath.Join(agentStateDir, "handoff.json")
	initialHandoff := map[string]interface{}{
		"last_updated": "2020-01-01T00:00:00Z",
	}
	bytes, _ := json.Marshal(initialHandoff)
	_ = os.WriteFile(handoffPath, bytes, 0644)

	statePath := filepath.Join(tempDir, "agent", "STATE.md")
	_ = os.WriteFile(statePath, []byte("# Global System State\n- Last update: 2020-01-01T00:00:00Z\n"), 0644)

	if err := touchTimestamps(tempDir); err != nil {
		t.Fatalf("touchTimestamps failed: %v", err)
	}

	newBytes, _ := os.ReadFile(handoffPath)
	var updatedMap map[string]interface{}
	_ = json.Unmarshal(newBytes, &updatedMap)
	if updatedMap["last_updated"] == "2020-01-01T00:00:00Z" {
		t.Errorf("Expected last_updated to be updated, but was unchanged")
	}
}

func TestResetRetryCount(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "mak-reset-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	agentStateDir := filepath.Join(tempDir, "agent", ".state")
	if err := os.MkdirAll(agentStateDir, 0755); err != nil {
		t.Fatalf("Failed to create agent/.state: %v", err)
	}

	handoffPath := filepath.Join(agentStateDir, "handoff.json")
	initialHandoff := map[string]interface{}{
		"status": "QA_FAILED",
		"feedback_loop": map[string]interface{}{
			"qa_retry_count": 3,
			"error_logs":     "error occurred",
		},
	}
	bytes, _ := json.Marshal(initialHandoff)
	_ = os.WriteFile(handoffPath, bytes, 0644)

	if err := resetRetryCount(tempDir); err != nil {
		t.Fatalf("resetRetryCount failed: %v", err)
	}

	newBytes, _ := os.ReadFile(handoffPath)
	var updatedMap map[string]interface{}
	_ = json.Unmarshal(newBytes, &updatedMap)
	if updatedMap["status"] != "IN_PROGRESS" {
		t.Errorf("Expected status to become IN_PROGRESS, got %s", updatedMap["status"])
	}
	fb := updatedMap["feedback_loop"].(map[string]interface{})
	if fb["qa_retry_count"].(float64) != 0 {
		t.Errorf("Expected qa_retry_count to be 0, got %v", fb["qa_retry_count"])
	}
}
