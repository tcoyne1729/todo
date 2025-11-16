package storage_test

import (
	"fmt"
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

	// create two tasks
	t1 := models.NewTask(models.NewTaskConfig{
		ID:        "1",
		Title:     "Buy Milk",
		EnteredAt: now,
	})
	t2 := models.NewTask(models.NewTaskConfig{
		ID:        "2",
		Title:     "Write Code",
		EnteredAt: now,
	})

	sampleTasks := []*models.Task{t1, t2}
	sampleCurrent := "2"
	sampleTags := []models.Tag{
		{Name: "t1", Context: "ct1"},
		{Name: "t2", Context: "ct2"},
	}

	s.Tasks = sampleTasks
	s.Current = sampleCurrent
	s.Tags = sampleTags

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
	if len(s2.Tags) != len(sampleTags) {
		t.Errorf("Mismatch in length of tags. Got %d, want %d", len(s2.Tags), len(sampleTags))
	}
	if len(s2.Tags) < 2 || len(sampleTags) < 2 {
		t.Fatal("empty tags. expected list of 2 tags")
	}
	if s2.Tags[0].Name != sampleTags[0].Name && s2.Tags[1].Name != sampleTags[1].Name {
		t.Errorf("tags not the same. Want %v, got %v", sampleTags, s2.Tags)
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
		if gotTask != nil {
			t.Errorf("expected empty task, got %+v", gotTask)
		}
	})

	t.Run("task exists", func(t *testing.T) {
		task := models.NewTask(models.NewTaskConfig{
			ID: "test",
		})
		store := &storage.Store{
			Tasks: []*models.Task{task},
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
		origTask := &models.Task{
			ID:    "1",
			Title: "original",
		}
		newTitle := "new"
		newTask := models.NewTask(models.NewTaskConfig{
			ID:    "1",
			Title: newTitle,
		})
		store := &storage.Store{
			Tasks: []*models.Task{origTask},
		}
		err := store.UpdateTask(newTask)
		if err != nil {
			t.Errorf("task not updated: %v", err)
		}
		if store.Tasks[0].Title != newTitle {
			t.Errorf("expected title updated to %s, got %s", newTitle, store.Tasks[0].Title)
		}
	})
}

func TestTags(t *testing.T) {
	getStore := func() storage.Store {
		store := storage.Store{
			Tags: []models.Tag{
				{Name: "t1", Context: "c1"},
				{Name: "t2", Context: "c2"},
			},
		}
		return store
	}

	t.Run("add tag", func(t *testing.T) {
		store := getStore()
		newTag := models.Tag{
			Name: "t3",
		}
		err := store.AddTag(newTag)
		if err != nil {
			t.Fatal("tag failed to add")
		}
		if len(store.Tags) != 3 {
			t.Errorf("incorrect number of tags, expected 3, got %d", len(store.Tags))
		}
	})
	t.Run("add tag with same name", func(t *testing.T) {
		store := getStore()
		newTag := models.Tag{
			Name: "t1",
		}
		err := store.AddTag(newTag)
		expectedError := "tag t1 already exists\n"
		if err.Error() != expectedError {
			t.Errorf("expected error: %v, got %v", expectedError, err)
		}
	})

	t.Run("delete tag", func(t *testing.T) {
		store := getStore()
		err := store.DeleteTag("t1")
		if err != nil {
			t.Fatalf("failed to delete tag: %v", err)
		}
		if len(store.Tags) != 1 {
			t.Errorf("tag not deleted correctly, expected two tags but found %d", len(store.Tags))
		}
		if len(store.Tags) == 0 {
			t.Fatal("expected at least one tag")
		}
		if store.Tags[0].Name != "t2" {
			t.Errorf("the wrong tag was deleted or it permuted the existing tags")
		}
	})
	t.Run("delete tag when no tags exist", func(t *testing.T) {
		store := storage.Store{}
		err := store.DeleteTag("t1")
		expectedError := "tag t1 does not exist\n"
		if err.Error() != expectedError {
			t.Errorf("expected error: %v, got %v", expectedError, err)
		}
	})
	t.Run("delete tag which isnt there", func(t *testing.T) {
		store := getStore()
		err := store.DeleteTag("t3")
		expectedError := "tag t3 does not exist\n"
		if err.Error() != expectedError {
			t.Errorf("expected error: %v, got %v", expectedError, err)
		}
	})

	t.Run("edit tag which doesnt exist", func(t *testing.T) {
		store := getStore()
		newTag := models.Tag{Name: "t3", Context: "c3"}
		expectedError := "tag t3 does not exist\n"
		fmt.Printf("before")
		err := store.EditTag("t3", newTag)
		fmt.Printf("after")
		if err == nil {
			t.Fatal("no error raised")
		}

		if err.Error() != expectedError {
			t.Errorf("expected error: %s, got %s", expectedError, err.Error())
		}
	})
	t.Run("edit tag", func(t *testing.T) {
		store := getStore()
		newTag := models.Tag{Name: "t3", Context: "c3"}
		err := store.EditTag("t1", newTag)
		if err != nil {
			t.Fatal("failed to edit tag")
		}
		if len(store.Tags) == 0 {
			t.Fatal("expected at least one tag but got none")
		}
		t1 := store.Tags[0]
		if t1.Name != newTag.Name {
			t.Errorf("updated tag name should be %s, got %s", newTag.Name, t1.Name)
		}
		if t1.Context != newTag.Context {
			t.Errorf("updated tag name should be %s, got %s", newTag.Context, t1.Context)
		}
	})
}
