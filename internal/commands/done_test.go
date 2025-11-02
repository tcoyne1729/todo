package commands_test

import (
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
	"testing"
	"time"
)

func TestDone(t *testing.T) {
	t.Run("simple case", func(t *testing.T) {
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
		cmd := commands.DoneCmd{}
		if err := cmd.Run(store); err != nil {
			t.Fatal("could not run done command")
		}
		// expect t1 now done
		updatedT1, err := store.GetTask("t1")
		if err != nil {
			t.Fatal("cant load task after done command")
		}
		if updatedT1.IsActive() {
			t.Errorf("task still active after done command")
		}
	})

	t.Run("autoclose case", func(t *testing.T) {

		worklog1 := models.WorkSession{
			StartedAt: time.Date(2025, time.April, 1, 0, 0, 0, 0, time.UTC),
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
		cmd := commands.DoneCmd{}
		if err := cmd.Run(store); err != nil {
			t.Fatal("could not run done command")
		}
		// expect t1 now done
		updatedT1, err := store.GetTask("t1")
		if err != nil {
			t.Fatal("cant load task after done command")
		}
		if updatedT1.IsActive() {
			t.Errorf("task still active after done command")
		}
		wls := updatedT1.WorkLog
		if len(wls) != 1 {
			t.Errorf("incorrect number of worklogs, expected 1 got %d", len(wls))
		}
		if !wls[0].AutoClosed {
			t.Errorf("should have been autoclosed")
		}

	})
}
