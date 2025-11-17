package commands

import (
	"fmt"
	// "time"
	//
	// "github.com/google/uuid"
	// "github.com/tcoyne1729/todo/internal/models"
	genericnotes "github.com/tcoyne1729/todo/internal/generic_notes"
	"github.com/tcoyne1729/todo/internal/storage"
)

type BlockCmd struct {
	ID string
}

func (b *BlockCmd) Add(store *storage.Store, blockedBy string, note string) (string, error) {
	if store.Current == "" && b.ID == "" {
		return "", fmt.Errorf("no task to mark as blocked")
	}
	var targetId string
	if b.ID != "" {
		targetId = b.ID
	} else {
		targetId = store.Current
	}
	blockedTask, err := store.GetTask(targetId)
	if err != nil {
		return "", err
	}
	newId, err := blockedTask.Blockers.New(genericnotes.NewConfig{
		Text: note,
	})
	if err != nil {
		return "", err
	}
	block, err := blockedTask.Blockers.GetNote(newId)
	if err != nil {
		return "", err
	}
	block.BlockedBy = blockedBy
	return newId, nil
}
func (b *BlockCmd) AddNote(store *storage.Store, id string, note string) error {
	return nil
}
func (b *BlockCmd) Remove(store *storage.Store, id string) error {
	return nil
}
func (b *BlockCmd) List() {}
