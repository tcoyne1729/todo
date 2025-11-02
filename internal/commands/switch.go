package commands

import (
	"fmt"

	"github.com/tcoyne1729/todo/internal/storage"
)

type SwitchCmd struct {
	ID string `arg:"" help:"The id of the task you want to switch to."`
}

func (c *SwitchCmd) Run(store *storage.Store) error {
	if c.ID == "" {
		return fmt.Errorf("no id provided for the task you want to switch to")
	}
	if store.Current != "" {
		// stop the current task
		currentTask, err := store.GetTask(store.Current)
		if err != nil {
			return err
		}
		if err := StopIfStarted(&currentTask, store); err != nil {
			return err
		}
	}
	// get the task
	switchToTask, err := store.GetTask(c.ID)
	if err != nil {
		return err
	}
	// if the switch to task had been previously started, then stop it
	if err := StopIfStarted(&switchToTask, store); err != nil {
		return err
	}

	// start new task
	start := StartCmd{
		ID: c.ID,
	}
	if err := start.Run(store); err != nil {
		return err
	}

	return nil
}
