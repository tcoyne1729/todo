package commands

import (
	"fmt"
	"time"

	genericnotes "github.com/tcoyne1729/todo/internal/generic_notes"
	"github.com/tcoyne1729/todo/internal/models"
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
	// store data
	err = store.SaveAll()
	if err != nil {
		return "", err
	}
	return newId, nil
}

func (b *BlockCmd) AddNote(store *storage.Store, blockID string, note string) (string, error) {
	task, err := store.GetTask(b.ID)
	if err != nil {
		return "", err
	}
	blocker, err := task.Blockers.GetNote(blockID)
	if err != nil {
		return "", err
	}
	commentID, err := blocker.BlockerNotes.New(genericnotes.NewConfig{Text: note})
	if err != nil {
		return "", err
	}
	// store
	if err = store.SaveAll(); err != nil {
		return "", err
	}
	return commentID, nil
}

func (b *BlockCmd) Complete(store *storage.Store, blockID string) error {
	task, err := store.GetTask(b.ID)
	if err != nil {
		return err
	}
	blocker, err := task.Blockers.GetNote(blockID)
	if err != nil {
		return err
	}
	blocker.SetCompleteTime(time.Now())
	// store
	if err = store.SaveAll(); err != nil {
		return err
	}
	return nil
}

func (b *BlockCmd) List(store *storage.Store, all bool) ([]*models.Blocker, error) {
	task, err := store.GetTask(b.ID)
	if err != nil {
		return nil, err
	}
	if all {
		return task.Blockers.ListAll()
	} else {
		return task.Blockers.ListActive()
	}
}
