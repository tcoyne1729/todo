package tags

import (
	"fmt"

	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/output"
	"github.com/tcoyne1729/todo/internal/storage"
)

type TagAddCmd struct {
	Name    string `help:"name of the tag."`
	Context string `help:"background and context to attach to the tag. This can be helpful to generate LLM output."`
}

func (t *TagAddCmd) Run(store *storage.Store) error {
	printer := output.NewPrinter(true)
	printer.Println(output.Red(), "Hello!!")
	if t.Name == "" {
		return fmt.Errorf("empty string is an invalid tag name")
	}
	newTag := models.Tag{
		Name:    t.Name,
		Context: t.Context,
	}
	if err := store.AddTag(newTag); err != nil {
		return err
	}
	if err := store.SaveAll(); err != nil {
		return err
	}
	return nil
}
