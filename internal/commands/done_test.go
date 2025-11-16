package commands_test

import (
	"testing"
	"time"

	"github.com/tcoyne1729/todo/internal/commands"
	genericnotes "github.com/tcoyne1729/todo/internal/generic_notes"
	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
)

func TestDone(t *testing.T) {
	t.Run("simple case", func(t *testing.T) {
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
		cmd := commands.DoneCmd{}
		if err := cmd.Run(store); err != nil {
			t.Fatalf("could not run done command, err: %v", err)
		}
		// expect t1 now done
		updatedT1, err := store.GetTask("t1")
		if err != nil {
			t.Fatal("cant load task after done command")
		}
		if active, err := updatedT1.IsActive(); active || err != nil {
			t.Errorf("task still active after done command")
		}
	})

	t.Run("autoclose case", func(t *testing.T) {
		task1 := models.NewTask(models.NewTaskConfig{
			ID:     "t1",
			Title:  "Task1",
			Status: "in_progress",
		})
		task1.WorkLog.New(genericnotes.NewConfig{CreateTime: time.Date(2025, time.April, 1, 0, 0, 0, 0, time.UTC)})
		task2 := models.NewTask(models.NewTaskConfig{
			ID:     "t2",
			Title:  "Task2",
			Status: "todo",
		})
		store := &storage.Store{
			Tasks:   []*models.Task{task1, task2},
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
		if active, err := updatedT1.IsActive(); active || err != nil {
			t.Errorf("task still active after done command")
		}
		wls := updatedT1.WorkLog
		if len(wls.Data) != 1 {
			t.Errorf("incorrect number of worklogs, expected 1 got %d", len(wls.Data))
		}
		if !wls.Data[0].AutoClosed {
			t.Errorf("should have been autoclosed")
		}

	})
}
