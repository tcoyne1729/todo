package commands_test

import (
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
	"testing"
)

func TestStart(t *testing.T) {
	t.Run("start a new task", func(t *testing.T) {
		task := models.Task{
			ID:     "t",
			Title:  "Task",
			Status: "todo",
		}
		store := &storage.Store{
			Tasks:   []models.Task{task},
			Current: "t",
		}
		startCmd := commands.StartCmd{
			ID: "t",
		}
		if err := startCmd.Run(store); err != nil {
			t.Errorf("error starting: %v", err)
		}
		// expectations:
		// task should now be in_progress and have an open worklog
		gotTask, err := store.GetTask("t")
		if err != nil {
			t.Errorf("error loading task after start")
		}
		gotWorkLog := gotTask.WorkLog
		if len(gotWorkLog) != 1 {
			t.Errorf("task worklog entry should have len 1 but got len %d", len(gotWorkLog))
		}
		if gotWorkLog[0].EndedAt != nil {
			t.Errorf("task 2 worklog should be open but is closed")
		}

	})

	t.Run("store is updated after start", func(t *testing.T) {
		task := models.Task{
			ID:     "t",
			Title:  "Task",
			Status: "todo",
		}
		store := &storage.Store{
			Tasks:   []models.Task{task},
			Current: "t",
		}

		startCmd := commands.StartCmd{ID: "t"}
		if err := startCmd.Run(store); err != nil {
			t.Fatalf("error starting task: %v", err)
		}

		// Tasks should be mutated
		updated := store.Tasks[0]
		if updated.Status != "in_progress" {
			t.Errorf("expected store.Tasks[0].Status = in_progress, got %q", updated.Status)
		}
		if len(updated.WorkLog) != 1 {
			t.Errorf("expected store.Tasks[0].WorkLog length = 1, got %d", len(updated.WorkLog))
		}
		if updated.WorkLog[0].EndedAt != nil {
			t.Errorf("expected open WorkLog, got EndedAt=%v", updated.WorkLog[0].EndedAt)
		}
	})

	t.Run("store is updated after start 2", func(t *testing.T) {
		task := models.Task{
			ID:     "t",
			Title:  "Task",
			Status: "todo",
		}
		store := &storage.Store{
			Tasks:   []models.Task{task},
			Current: "t",
		}

		before := store.Tasks[0] // capture the original
		startCmd := commands.StartCmd{ID: "t"}
		if err := startCmd.Run(store); err != nil {
			t.Fatalf("error starting: %v", err)
		}
		after := store.Tasks[0]

		if before.Status == after.Status {
			t.Errorf("expected task status to change from %q to something else", before.Status)
		}
		if len(after.WorkLog) == 0 {
			t.Errorf("expected WorkLog to be updated, got none")
		}
	})
}
