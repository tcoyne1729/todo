package tasktag

import "github.com/tcoyne1729/todo/internal/storage"

type RemoveTagsCmd struct {
	Tags []string `help:"tags to remove from task"`
}

func (t *RemoveTagsCmd) Run(store *storage.Store) error {
	return nil
}
