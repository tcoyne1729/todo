package commands

import (
	"fmt"

	// "github.com/tcoyne1729/todo/internal/commands/update"
	"github.com/tcoyne1729/todo/internal/storage"
)

type TaskIDContext struct {
	TaskID string `arg:"" help:"the id of the task we want to manipulate the tags for"`
}
type TagCmd struct {
	// ID     TaskIDContext        `arg:""`
	Tag string `arg`
	// Add    update.AddTagsCmd    `cmd:"" help:"add tasks (space separated) to a task"`
	// Remove update.RemoveTagsCmd `cmd:"" help:"remove tasks (space separated) from a task"`
}

func (t *TagCmd) Run(store *storage.Store) error {
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
