package commands_test

import (
	"testing"

	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
)

func TestStart(t *testing.T) {
	t.Run("start a new task", func(t *testing.T) {
		task := models.NewTask(models.NewTaskConfig{
			ID:     "t",
			Title:  "Task",
			Status: "todo",
		})
		store := &storage.Store{
			Tasks:   []*models.Task{task},
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
		gotWorkLog := gotTask.WorkLog.Data
		if len(gotWorkLog) != 1 {
			t.Fatalf("task worklog entry should have len 1 but got len %d", len(gotWorkLog))
		}
		if gotWorkLog[0].CompleteTime != nil {
			t.Errorf("task 2 worklog should be open but is closed")
		}

	})

	t.Run("store is updated after start", func(t *testing.T) {
		task := models.NewTask(models.NewTaskConfig{
			ID:     "t",
			Title:  "Task",
			Status: "todo",
		})
		store := &storage.Store{
			Tasks:   []*models.Task{task},
			Current: "t",
		}

		startCmd := commands.StartCmd{ID: "t"}
		if err := startCmd.Run(store); err != nil {
			t.Fatalf("error starting task: %v", err)
		}

		if task.Status != "in_progress" {
			t.Errorf("expected store.Tasks[0].Status = in_progress, got %q", task.Status)
		}
		if len(task.WorkLog.Data) != 1 {
			t.Errorf("expected store.Tasks[0].WorkLog length = 1, got %d", len(task.WorkLog.Data))
		}
		if task.WorkLog.Data[0].CompleteTime != nil {
			t.Errorf("expected open WorkLog, got EndedAt=%v", task.WorkLog.Data[0].CompleteTime)
		}
	})

	t.Run("store is updated after start 2", func(t *testing.T) {
		task := models.NewTask(models.NewTaskConfig{
			ID:     "t",
			Title:  "Task",
			Status: "todo",
		})
		store := &storage.Store{
			Tasks:   []*models.Task{task},
			Current: "t",
		}

		startCmd := commands.StartCmd{ID: "t"}
		if err := startCmd.Run(store); err != nil {
			t.Fatalf("error starting: %v", err)
		}

		if "todo" == task.Status {
			t.Errorf("expected task status to change from %q to something else", "todo")
		}
		if len(task.WorkLog.Data) == 0 {
			t.Errorf("expected WorkLog to be updated, got none")
		}
	})
}
