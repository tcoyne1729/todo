package tags

import (
	"fmt"

	"github.com/tcoyne1729/todo/internal/storage"
)

type TagListCmd struct{}

func (t *TagListCmd) Run(store *storage.Store) error {
	for _, tag := range store.Tags {
		fmt.Printf("Name: %s\nContext: %s\n\n", tag.Name, tag.Context)
	}
	return nil
}
