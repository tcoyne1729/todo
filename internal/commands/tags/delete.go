package tags

import (
	"fmt"

	"github.com/tcoyne1729/todo/internal/storage"
)

type TagDeleteCmd struct {
	Name string `help:"name of tag to delete"`
}

func (c *TagDeleteCmd) Run(store *storage.Store) error {
	if c.Name == "" {
		return fmt.Errorf("tag name cannot be empty")
	}
	store.DeleteTag(c.Name)
	if err := store.SaveAll(); err != nil {
		return err
	}
	return nil
}
