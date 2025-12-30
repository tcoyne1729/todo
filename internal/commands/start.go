package commands

import (
	"fmt"
	"time"

	genericnotes "github.com/tcoyne1729/todo/internal/generic_notes"
	"github.com/tcoyne1729/todo/internal/storage"
)

type StartCmd struct {
	ID string `arg:"" help:"ID of the task you are starting."`
}

func (c *StartCmd) Run(store *storage.Store) error {
	// update the work
	var fullId string
	if len(c.ID) < 36 {
		// this must be a short ID
		id, err := ShortToLongId(c.ID, store)
		if err != nil {
			return err
		}
		fullId = id
	} else {
		fullId = c.ID
	}
	curTask, err := store.GetTask(fullId)
	if err != nil {
		return err
	}
	lastWork, err := curTask.WorkLog.GetLast()
	if err != nil {
		return err
	}
	if lastWork != nil {
		if lastWork.IsActive() {
			return fmt.Errorf("this task is already started")
		}
	}
	newId, err := curTask.WorkLog.New(genericnotes.NewConfig{})
	if err != nil {
		return err
	}
	curTask.Status = "in_progress"
	newWorkLog, err := curTask.WorkLog.GetNote(newId)
	if err != nil {
		return err
	}

	store.Current = fullId
	if err := store.SaveAll(); err != nil {
		return fmt.Errorf("failed to start task: %w", err)
	}
	fmt.Printf("%s, Started task: %s\n", newWorkLog.CreateTime.Format(time.Kitchen), curTask.Title)
	return nil
}
