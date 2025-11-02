package commands

import (
	"fmt"
	"github.com/tcoyne1729/todo/internal/storage"
	"time"
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
