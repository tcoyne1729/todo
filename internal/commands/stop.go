package commands

import (
	"fmt"
	"time"

	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
)

type StopCmd struct {
	ID         string    `help:"The id of the task you want to stop."`
	CloseTime  time.Time `help:"Time the task is to be closed at."`
	AutoClosed bool      `short:"a" help:"if this task is autoclosed because it overran the threshold."`
}

func (c *StopCmd) Run(store *storage.Store) error {
	task, err := Stop(c.ID, c.CloseTime, c.AutoClosed, store)
	if err != nil {
		return err
	}
	if err := store.UpdateTask(task); err != nil {
		return err
	}
	// unset current task
	store.Current = ""
	if err := store.SaveAll(); err != nil {
		return fmt.Errorf("failed to stop task: %w", err)
	}
	return nil
}

func Stop(id string, closeTime time.Time, autoClosed bool, store *storage.Store) (*models.Task, error) {
	if id == "" {
		id = store.Current
		if store.Current == "" {
			return nil, fmt.Errorf("no ID and no current item")
		}
	}
	newTask, err := store.GetTask(id)
	if err != nil {
		return nil, err
	}
	// check if the session has already been started
	lastWorkLog, err := newTask.WorkLog.GetLast()
	if err != nil {
		return nil, fmt.Errorf("task has no active work sessions: id = %s", id)
	}
	if lastWorkLog.CompleteTime != nil {
		return nil, fmt.Errorf("last task has already been ended: id = %s", id)
	}
	var endTime time.Time
	if closeTime.IsZero() {
		endTime = time.Now()
	} else {
		endTime = closeTime
	}
	lastWorkLog.SetCompleteTime(endTime)
	lastWorkLog.AutoClosed = autoClosed
	return newTask, nil
}
