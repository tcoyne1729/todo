package commands

import (
	"fmt"
	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
	"time"
)

type NoteCmd struct {
	Note string `arg:"" help:"Note to add to the current task"`
	ID   string `help:"If you want to add a note to a non-current task, add the ID here"`
}

func (l *NoteCmd) Run(store *storage.Store) error {
	var current string
	if (store.Current == "") && (l.ID == "") {
		return fmt.Errorf("no current task to add note to")
	} else if l.ID == "" {
		current = store.Current
	} else {
		current = l.ID
	}
	task, err := store.GetTask(current)
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
