package commands

import (
	"fmt"
	"github.com/tcoyne1729/todo/internal/storage"
	"time"
)

type StartCmd struct {
	ID string `arg:"" help:"ID of the task you are starting."`
}

func (c *StartCmd) Run(store *storage.Store) error {
	// update the work
	newTask, err := Start(c.ID, store)
	if err != nil {
		return err
	}
	// garunteed to have a worklog from the Start function as long as no err
	newWorkLog := newTask.WorkLog[len(newTask.WorkLog)-1]

	// if err := store.UpdateTask(newTask); err != nil {
	// 	return err
	// }
	err = store.UpdateTask(newTask)
	if err != nil {
		return err
	}
	// update current task
	store.Current = c.ID
	if err := store.SaveAll(); err != nil {
		return fmt.Errorf("failed to start task: %w", err)
	}
	fmt.Printf("%s, Started task: %s\n", newWorkLog.StartedAt.Format(time.Kitchen), newTask.Title)
	return nil
}
