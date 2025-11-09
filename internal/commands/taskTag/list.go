package tasktag

import (
	"fmt"

	"github.com/tcoyne1729/todo/internal/storage"
)

type TagListCmd struct {
	Tag string `arg`
}

func (t *TagListCmd) Run(store *storage.Store) error {
	currentTask, err := store.GetTask(t.Tag)
	if err != nil {
		return err
	}
	currentTags := currentTask.Tags
	fmt.Println("Current tags:")
	if len(currentTags) == 0 {
		return nil
	}
	for _, tag := range currentTags {
		fmt.Printf("%s\n", tag)
	}
	return nil
}
