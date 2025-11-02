package commands

import (
	"fmt"
	"github.com/tcoyne1729/todo/internal/storage"
)

type ListCmd struct{}

func (l *ListCmd) Run(store *storage.Store) error {
	allTasks := store.ListTasks()
	if len(allTasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}
	currentTask := store.Current

	for i, t := range allTasks {
		fmt.Printf("%d: %s [%s] %s %s\n", i, PointString(t.ID, currentTask), t.Status, t.Title, t.ID)
	}
	return nil
}
