package commands

import (
	"fmt"

	"github.com/tcoyne1729/todo/internal/storage"
)

type StatusCmd struct{}

func (s *StatusCmd) Run(store *storage.Store) error {
	if store.Current == "" {
		return fmt.Errorf("No task selected")
	}
	task, err := store.GetTask(store.Current)
	if err != nil {
		return err
	}
	fmt.Printf("Current Task\n\nTitle: %s\nID: %s\nStatus: %s\nBody: %s\n", task.Title, task.ID, task.Status, task.Body)
	notes := task.Notes
	activeNotes, err := notes.ListAll()
	if err != nil {
		return err
	}
	if len(activeNotes) > 0 {
		fmt.Println("\nNotes:")
		for _, note := range activeNotes {
			fmt.Printf("%s: %s\n", note.CreateTime.Format("2006-01-02 15:04:05 MST"), note.Text)
		}
	}
	return nil
}
