package commands_test

import (
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
	"testing"
	"time"
)

func TestStop(t *testing.T) {
	t.Run("stop a running task", func(t *testing.T) {
		worklog := models.WorkSession{
			StartedAt: time.Now(),
		}
		task := models.Task{
			ID:      "t1",
			Title:   "Task1",
			Status:  "in_progress",
			WorkLog: []models.WorkSession{worklog},
		}
		store := &storage.Store{
			Tasks:   []models.Task{task},
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
		t1WorkLog := gotTask.WorkLog[0]
		if t1WorkLog.EndedAt == nil {
			t.Errorf("task 1 was not ended correctly. worklog: %v", t1WorkLog)
		}
	})

}
