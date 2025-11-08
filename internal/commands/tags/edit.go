package tags

import (
	"fmt"

	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
)

type TagEditCmd struct {
	OldName    string `help:"existing tag name to be updated"`
	NewName    string `help:"new tag name"`
	NewContext string `help:"new tag context"`
}

func (t *TagEditCmd) Run(store *storage.Store) error {
	existing, err := store.GetTag(t.OldName)
	if err != nil {
		return err
	}
	newName := t.NewName
	if t.NewName == "" {
		fmt.Println("no new name provided - skipping")
		newName = existing.Name
	}
	newContext := t.NewContext
	if t.NewContext == "" {
		fmt.Println("no new context provided - skipping")
		newContext = existing.Context
	}
	newTag := models.Tag{
		Name:    newName,
		Context: newContext,
	}
	err = store.EditTag(t.OldName, newTag)
	if err = store.SaveAll(); err != nil {
		return err
	}
	return nil
}
