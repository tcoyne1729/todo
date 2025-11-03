package commands

import (
	"fmt"
	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
	"time"
)

type NoteCmd struct {
	Note string `arg:"" help:"Note to add to the current task"`
}

func (l *NoteCmd) Run(store *storage.Store) error {
	if store.Current == "" {
		return fmt.Errorf("no current task to add note to")
	}
	task, err := store.GetTask(store.Current)
	if err != nil {
		return err
	}

	newNote := models.Note{
		TimeStamp: time.Now(),
		Note:      l.Note,
	}
	task.Notes = append(task.Notes, newNote)

	if err = store.UpdateTask(task); err != nil {
		return err
	}

	if err = store.SaveAll(); err != nil {
		return err
	}

	return nil
}
