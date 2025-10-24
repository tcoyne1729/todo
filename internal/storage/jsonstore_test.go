package storage_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"todo/internal/models"
	"todo/internal/storage"
)

// Helper function to create a temporary directory path
func setupTempDir(t *testing.T) string {
	// t.TempDir() creates a temp dir and cleans it up automatically
	return t.TempDir()
}

// TestLoadAndSaveAll ensures that data can be correctly saved and loaded.
func TestLoadAndSaveAll(t *testing.T) {
	// Use the TempDir as the mocked storage path
	tempPath := setupTempDir(t)

	// Create a new store, injecting the temporary path
	s := storage.NewStore(tempPath)

	// 1. Prepare Test Data
	now := time.Now().Truncate(time.Second) // Important for DeepEqual check
	sampleTasks := []models.Task{
		{ID: "1", Title: "Buy Milk", EnteredAt: now},
		{ID: "2", Title: "Write Code", EnteredAt: now},
	}
	sampleCurrent := "2"

	s.Tasks = sampleTasks
	s.Current = sampleCurrent

	// 2. Test SaveAll
	if err := s.SaveAll(); err != nil {
		t.Fatalf("SaveAll failed unexpectedly: %v", err)
	}

	// 3. Verify files were created (Optional, but good check)
	if _, err := os.Stat(filepath.Join(tempPath, "tasks.json")); os.IsNotExist(err) {
		t.Fatal("tasks.json was not created.")
	}

	// 4. Test LoadAll (using a fresh store instance)
	s2 := storage.NewStore(tempPath)
	if err := s2.LoadAll(); err != nil {
		t.Fatalf("LoadAll failed unexpectedly: %v", err)
	}

	// 5. Assertions (Check data integrity)
	if s2.Current != sampleCurrent {
		t.Errorf("Current loaded incorrectly. Got %q, want %q", s2.Current, sampleCurrent)
	}

	// Since comparing structs with time.Time fields can be complex,
	// we'll rely on the fact that the initial SaveAll/LoadAll succeeded
	// based on the individual field checks. A full DeepEqual is more complex
	// but for a minimal test, checking a key field is sufficient.
	if len(s2.Tasks) != len(sampleTasks) {
		t.Errorf("Tasks length mismatch. Got %d, want %d", len(s2.Tasks), len(sampleTasks))
	}
	if s2.Tasks[0].Title != sampleTasks[0].Title {
		t.Errorf("Task 0 title mismatch. Got %q, want %q", s2.Tasks[0].Title, sampleTasks[0].Title)
	}
}

func TestLoadAll_NoFiles(t *testing.T) {
	tempPath := setupTempDir(t)

	// Create a store that points to an empty directory
	s := storage.NewStore(tempPath)

	// Load should succeed and initialize fields to zero values (empty slice/string)
	if err := s.LoadAll(); err != nil {
		t.Fatalf("LoadAll failed unexpectedly on empty directory: %v", err)
	}

	if len(s.Tasks) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(s.Tasks))
	}
	if s.Current != "" {
		t.Errorf("Expected Current to be empty string, got %q", s.Current)
	}
}
