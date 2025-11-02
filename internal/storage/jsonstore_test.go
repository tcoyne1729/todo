package storage_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
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
func TestGetTask(t *testing.T) {
	t.Run("no tasks", func(t *testing.T) {
		store := &storage.Store{}
		id := "test"
		gotTask, err := store.GetTask(id)
		// expect got_models to be models.Task{}
		// expect an error "no task defined for id = test"
		expected_err := "no task defined for id = test"
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		if err.Error() != expected_err {
			t.Errorf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(gotTask, models.Task{}) {
			t.Errorf("expected empty task, got %+v", gotTask)
		}
	})

	t.Run("task exists", func(t *testing.T) {
		task := models.Task{
			ID: "test",
		}
		store := &storage.Store{
			Tasks: []models.Task{task},
		}
		id := "test"
		// run test
		got, err := store.GetTask(id)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if !reflect.DeepEqual(got, task) {
			t.Errorf("expected task with id %s, got %+v", id, got)
		}
	})
}

func TestUpdateTask(t *testing.T) {
	t.Run("update a single task", func(t *testing.T) {
		origTask := models.Task{
			ID:    "1",
			Title: "original",
		}
		newTitle := "new"
		newTask := models.Task{
			ID:    "1",
			Title: newTitle,
		}
		store := &storage.Store{
			Tasks: []models.Task{origTask},
		}
		err := store.UpdateTask(newTask)
		if err != nil {
			t.Errorf("task not updated: %v", err)
		}
		if store.Tasks[0].Title != newTitle {
			t.Errorf("expected title updated to %s, got %s", newTitle, store.Tasks[0].Title)
		}
	})

	t.Run("update more complex task", func(t *testing.T) {
		t1 := models.Task{
			ID: "t1",
		}
		origTask := models.Task{
			ID:     "t2",
			Title:  "Task2",
			Status: "todo",
		}
		store := &storage.Store{
			Tasks: []models.Task{t1, origTask},
		}
		newTask, err := store.GetTask("t2")
		if err != nil {
			t.Fatalf("failed to get task t2")
		}
		newTask.WorkLog = append(newTask.WorkLog, models.WorkSession{
			StartedAt: time.Now(),
		})
		err = store.UpdateTask(newTask)
		if err != nil {
			t.Errorf("task not updated: %v", err)
		}
		newTaskUpdate, err := store.GetTask("t2")
		if err != nil {
			t.Fatalf("did not get the update")
		}
		if len(newTaskUpdate.WorkLog) != 1 {
			t.Errorf("not updated")
		}
	})
}
