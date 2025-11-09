package tasktag

import "github.com/tcoyne1729/todo/internal/storage"

type AddTagsCmd struct {
	Tags []string `help:"tags to remove from task"`
}

func (t *AddTagsCmd) Run(store *storage.Store) error {
	return nil
}
