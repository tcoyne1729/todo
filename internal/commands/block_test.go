package commands_test

import (
	"testing"

	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
)

func TestBlock(t *testing.T) {
	t.Run("add single block", func(t *testing.T) {
		// setup
		task1 := models.NewTask(models.NewTaskConfig{
			ID: "t1",
		})
		store := &storage.Store{
			Tasks:   []*models.Task{task1},
			Current: "",
		}
		cmd := commands.BlockCmd{ID: "t1"}
		blockPerson := "block person"
		blockNote := "block note"
		blockID, err := cmd.Add(store, blockPerson, blockNote)
		if err != nil {
			t.Fatalf("failed to add task: %v", err)
		}
		// checks
		// we expect task 1 to now have a single blocker item

		if len(task1.Blockers.Data) == 0 {
			t.Fatalf("no blocker items found, got: %d", len(task1.Blockers.Data))
		}
		gotBlock, err := task1.Blockers.GetNote(blockID)
		if err != nil {
			t.Fatalf("failed to load blocker: %v", err)
		}
		if gotBlock.Text != blockNote {
			t.Errorf("blockedby want: %s, got: %s", blockNote, gotBlock.Text)
		}
		if gotBlock.BlockedBy != blockPerson {
			t.Errorf("blockedby want: %s, got: %s", blockPerson, gotBlock.BlockedBy)
		}
	})
}
