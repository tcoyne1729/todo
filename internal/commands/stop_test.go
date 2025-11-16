package commands_test

import (
	"testing"

	"github.com/tcoyne1729/todo/internal/commands"
	genericnotes "github.com/tcoyne1729/todo/internal/generic_notes"
	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
)

func TestStop(t *testing.T) {
	t.Run("stop a running task", func(t *testing.T) {
		task := models.NewTask(models.NewTaskConfig{
			ID:     "t1",
			Title:  "Task1",
			Status: "in_progress",
		})
		task.WorkLog.New(genericnotes.NewConfig{})
		store := &storage.Store{
			Tasks:   []*models.Task{task},
			Current: "t1",
		}
		stopCmd := commands.StopCmd{
			ID: "t1",
		}
		if err := stopCmd.Run(store); err != nil {
			t.Errorf("error stopping: %v", err)
		}
		// expectations:
		// task1 should have a closed worklog
		// task2 should now be in_progress and have an open worklog
		gotTask, err := store.GetTask("t1")
		if err != nil {
			t.Errorf("error loading task after stop")
		}
		t1WorkLog, err := gotTask.WorkLog.GetLast()
		if err != nil {
			t.Fatal("could not get last work log")
		}
		if t1WorkLog.CompleteTime == nil {
			t.Errorf("task 1 was not ended correctly. worklog: %v", t1WorkLog)
		}
	})

}
