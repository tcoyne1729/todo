package commands

import (
	"fmt"
	"github.com/tcoyne1729/todo/internal/storage"
)

type DoneCmd struct{}

func (l *DoneCmd) Run(store *storage.Store) error {
	if store.Current == "" {
		return fmt.Errorf("no current task to mark as done")
	}
	doneTask, err := store.GetTask(store.Current)
	if err != nil {
		return err
	}
	if err := StopIfStarted(&doneTask, store); err != nil {
		return err
	}
	doneTask.Status = "done"

	if err = store.UpdateTask(doneTask); err != nil {
		return err
	}
	store.Current = ""

	if err = store.SaveAll(); err != nil {
		return err
	}

	return nil
}
