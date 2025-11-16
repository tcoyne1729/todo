package commands_test

import (
	"testing"

	"github.com/tcoyne1729/todo/internal/commands"
	genericnotes "github.com/tcoyne1729/todo/internal/generic_notes"
	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
)

func TestSwitch(t *testing.T) {
	t.Run("two tasks switch one new one existing", func(t *testing.T) {
		task1 := models.NewTask(models.NewTaskConfig{
			ID:     "t1",
			Title:  "Task1",
			Status: "in_progress",
		})
		task1.WorkLog.New(genericnotes.NewConfig{})
		task2 := models.NewTask(models.NewTaskConfig{
			ID:     "t2",
			Title:  "Task2",
			Status: "todo",
		})
		store := &storage.Store{
			Tasks:   []*models.Task{task1, task2},
			Current: "t1",
		}
		switchCmd := commands.SwitchCmd{
			ID: "t2",
		}
		if err := switchCmd.Run(store); err != nil {
			t.Errorf("error switching: %v", err)
		}
		// expectations:
		// task1 should have a closed worklog
		// task2 should now be in_progress and have an open worklog
		gotTask1, err := store.GetTask("t1")
		if err != nil {
			t.Errorf("error loading task1 after switch")
		}
		t1WorkLog, err := gotTask1.WorkLog.GetLast()
		if err != nil {
			t.Fatalf("error getting task: %v", err)
		}
		if t1WorkLog.CompleteTime == nil {
			t.Errorf("task 1 was not ended correctly. worklog: %v", t1WorkLog)
		}
		gotTask2, err := store.GetTask("t2")
		if err != nil {
			t.Fatalf("error loading task2 after switch")
		}
		t2WorkLog := gotTask2.WorkLog.Data
		if len(t2WorkLog) != 1 {
			t.Fatalf("task 2 worklog entry should have len 1 but got len %d", len(t2WorkLog))
		}
		if t2WorkLog[0].CompleteTime != nil {
			t.Errorf("task 2 worklog should be open but is closed")
		}

	})

}
