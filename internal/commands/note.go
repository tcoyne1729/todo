package commands

import (
	"fmt"

	genericnotes "github.com/tcoyne1729/todo/internal/generic_notes"
	"github.com/tcoyne1729/todo/internal/storage"
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
	task.Notes.New(genericnotes.NewConfig{
		Text: l.Note,
	})

	if err = store.UpdateTask(task); err != nil {
		return err
	}

	if err = store.SaveAll(); err != nil {
		return err
	}

	return nil
}
