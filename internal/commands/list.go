package commands

import (
	"fmt"
	"github.com/tcoyne1729/todo/internal/storage"
)

type ListCmd struct {
	All bool `short:"a" help:"add this flag to see all tasks. Default is false which removed done tasks."`
}

func (l *ListCmd) Run(store *storage.Store, output bool) ([]string, error) {
	allTasks := store.ListTasks()
	allIds := make([]string, 0)
	if len(allTasks) == 0 {
		fmt.Println("No tasks found.")
		return allIds, nil
	}
	currentTask := store.Current

	for _, t := range allTasks {
		if l.All || t.Status != "done" {
			allIds = append(allIds, t.ID)
		}
	}
	allShortIds, err := ShortIds(allIds)
	if err != nil {
		return allIds, err
	}
	for _, t := range allTasks {
		if l.All || t.Status != "done" {
			shortId, ok := allShortIds.GetLongToShort(t.ID)
			if !ok {
				shortId = ""
			}
			if output {
				fmt.Printf("[%s]: %s %s %s [%s]\n", shortId, PointString(t.ID, currentTask), t.Status, t.Title, t.ID)
			}
		}
	}
	return allIds, nil
}
