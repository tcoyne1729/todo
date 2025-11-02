package commands_test

import (
	"fmt"
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
	"testing"
	"time"
)

func TestSwitch(t *testing.T) {
	t.Run("two tasks switch one new one existing", func(t *testing.T) {
		worklog1 := models.WorkSession{
			StartedAt: time.Now(),
		}
		task1 := models.Task{
			ID:      "t1",
			Title:   "Task1",
			Status:  "in_progress",
			WorkLog: []models.WorkSession{worklog1},
		}
		task2 := models.Task{
			ID:     "t2",
			Title:  "Task2",
			Status: "todo",
		}
		store := &storage.Store{
			Tasks:   []models.Task{task1, task2},
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
		t1WorkLog := gotTask1.WorkLog[0]
		if t1WorkLog.EndedAt == nil {
			t.Errorf("task 1 was not ended correctly. worklog: %v", t1WorkLog)
		}
		gotTask2, err := store.GetTask("t2")
		if err != nil {
			t.Errorf("error loading task2 after switch")
		}
		t2WorkLog := gotTask2.WorkLog
		if len(t2WorkLog) != 1 {
			t.Fatalf("task 2 worklog entry should have len 1 but got len %d", len(t2WorkLog))
		}
		if t2WorkLog[0].EndedAt != nil {
			t.Errorf("task 2 worklog should be open but is closed")
		}

	})

	// test case where the switch to task was started the prior day but not stopped
	t.Run("task started the prior day", func(t *testing.T) {

		startTime := time.Date(2025, time.January, 15, 11, 1, 0, 0, time.Local)
		autoClosedTime := time.Date(2025, time.January, 16, 0, 0, 0, 0, time.Local)
		worklog1 := models.WorkSession{
			StartedAt: startTime,
		}
		task1 := models.Task{
			ID:      "t1",
			Title:   "Task1",
			Status:  "in_progress",
			WorkLog: []models.WorkSession{worklog1},
		}
		task2 := models.Task{
			ID:     "t2",
			Title:  "Task2",
			Status: "todo",
		}
		store := &storage.Store{
			Tasks:   []models.Task{task1, task2},
			Current: "t1",
		}
		switchCmd := commands.SwitchCmd{
			ID: "t2",
		}
		if err := switchCmd.Run(store); err != nil {
			t.Errorf("error switching: %v\n", err)
		}
		// expectations:
		// task 1 should be auto closed
		t1Output, err := store.GetTask("t1")
		fmt.Printf("t1Output: %v", t1Output)
		if err != nil {
			t.Fatal("could not get t1 task")
		}
		if len(t1Output.WorkLog) != 1 {
			t.Fatalf("worklog output shoud be 1 got %d", len(t1Output.WorkLog))
		}
		wl := t1Output.WorkLog[0]
		if !wl.AutoClosed {
			t.Errorf("worklog should have been autoclosed but got value %t", wl.AutoClosed)
		}
		if wl.EndedAt == nil {
			t.Fatal("endedat should be a *time.Time object but was nul pointer")
		}
		if !autoClosedTime.Equal(*wl.EndedAt) {
			t.Errorf("worklog should have been autoclosed at %s, got %s", autoClosedTime, wl.EndedAt)
		}

	})
}
