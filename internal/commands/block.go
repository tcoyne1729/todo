package commands

import (
	"fmt"
	// "time"
	//
	// "github.com/google/uuid"
	// "github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
)

type BlockCmd struct {
	ID string
}

func (b *BlockCmd) Add(store *storage.Store, blockedBy string, note string) error {
	if store.Current == "" && b.ID == "" {
		return fmt.Errorf("no task to mark as blocked")
	}
	var targetId string
	if b.ID != "" {
		targetId = b.ID
	} else {
		targetId = store.Current
	}
	blockedTask, err := store.GetTask(targetId)
	if err != nil {
		return err
	}
	fmt.Printf("blocked: %v", *blockedTask)
	// blockers := blockedTask.Blockers
	// newNote := models.Note{
	// 	TimeStamp: time.Now(),
	// 	Note:      note,
	// }
	// newBlocker := models.Blocker{
	// 	ID:        uuid.New().String(),
	// 	EnteredAt: time.Now(),
	// 	BlockedBy: blockedBy,
	// 	Notes:     []models.Note{newNote},
	// }
	// blockers = append(blockers, newBlocker)
	// // store updated task
	// if err = store.UpdateTask(blockedTask); err != nil {
	// 	return err
	// }
	// if err = store.SaveAll(); err != nil {
	// 	return err
	// }
	return nil
}
func (b *BlockCmd) AddNote(store *storage.Store, id string, note string) error {
	return nil
}
func (b *BlockCmd) Remove(store *storage.Store, id string) error {
	return nil
}
