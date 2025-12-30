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
		cmd := commands.BlockCmd{ID: task1.ID}
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

	t.Run("list when no blockers", func(t *testing.T) {
		// setup
		task1 := models.NewTask(models.NewTaskConfig{
			ID: "t1",
		})
		store := &storage.Store{
			Tasks:   []*models.Task{task1},
			Current: "",
		}
		expectedNoBlockers := 0
		cmd := commands.BlockCmd{ID: task1.ID}
		blockers, err := cmd.List(store, true)
		if err != nil {
			t.Fatalf("error listing blockers: %v", err)
		}
		if len(blockers) != expectedNoBlockers {
			t.Errorf("got %d tasks but expected %d", len(blockers), expectedNoBlockers)
		}

	})

	t.Run("Add blocker and mark as complete", func(t *testing.T) {

		// setup
		task1 := models.NewTask(models.NewTaskConfig{
			ID: "t1",
		})
		store := &storage.Store{
			Tasks:   []*models.Task{task1},
			Current: "",
		}
		cmd := commands.BlockCmd{ID: task1.ID}
		blockPerson := "block person"
		blockNote := "block note"
		blockID, err := cmd.Add(store, blockPerson, blockNote)
		if err != nil {
			t.Fatalf("failed to add task: %v", err)
		}
		if err = cmd.Complete(store, blockID); err != nil {
			t.Fatalf("failed to complete task: %v", err)
		}
		// check
		gotBlock, err := task1.Blockers.GetNote(blockID)
		if err != nil {
			t.Fatalf("failed to load blocker: %v", err)
		}
		if gotBlock.CompleteTime == nil {
			t.Errorf("blocker should be complete, got: %v", gotBlock)
		}
	})
	t.Run("add two blockers, mark one as complete and list", func(t *testing.T) {
		// setup
		task1 := models.NewTask(models.NewTaskConfig{
			ID: "t1",
		})
		store := &storage.Store{
			Tasks:   []*models.Task{task1},
			Current: "",
		}
		// add blocker 1
		cmd := commands.BlockCmd{ID: task1.ID}
		blockPerson := "block person1"
		blockNote := "block note1"
		blockID1, err := cmd.Add(store, blockPerson, blockNote)
		if err != nil {
			t.Fatalf("failed to add task: %v", err)
		}
		// add blocker 2
		// add blocker 1
		blockPerson2 := "block person2"
		blockNote2 := "block note2"
		blockID2, err := cmd.Add(store, blockPerson2, blockNote2)
		if err != nil {
			t.Fatalf("failed to add task: %v", err)
		}
		// complete blocker 1
		if err = cmd.Complete(store, blockID1); err != nil {
			t.Fatalf("failed to complete task: %v", err)
		}
		// expect to get all not just active
		gotBlocks, err := task1.Blockers.ListAll()
		if len(gotBlocks) != 2 {
			t.Errorf("got the wrong number of blockers. want 2 got %d", len(gotBlocks))
		}
		gotB1, err := task1.Blockers.GetNote(blockID1)
		if err != nil {
			t.Fatalf("failed to load blocker 1: %v", err)
		}
		gotB2, err := task1.Blockers.GetNote(blockID2)
		if err != nil {
			t.Fatalf("failed to load blocker 2: %v", err)
		}
		if gotB1.CompleteTime == nil {
			t.Errorf("blocker 1 should be complete, got %v", gotB1)
		}
		if gotB2.CompleteTime != nil {
			t.Errorf("blocker 2 should be uncomplete, got %v", gotB1)
		}
		gotActiveBlocks, err := task1.Blockers.ListActive()
		if err != nil {
			t.Fatalf("could not list active blockers: %v", err)
		}
		if len(gotActiveBlocks) != 1 {
			t.Errorf("expected 1 active task, got %d", len(gotActiveBlocks))
		}
	})
	t.Run("add blocker and add note", func(t *testing.T) {

		// setup
		task1 := models.NewTask(models.NewTaskConfig{
			ID: "t1",
		})
		store := &storage.Store{
			Tasks:   []*models.Task{task1},
			Current: "",
		}
		cmd := commands.BlockCmd{ID: task1.ID}
		blockPerson := "block person"
		blockNote := "block note"
		blockID, err := cmd.Add(store, blockPerson, blockNote)
		if err != nil {
			t.Fatalf("failed to add task: %v", err)
		}
		commentToAdd := "comment text"
		commentID, err := cmd.AddNote(store, blockID, commentToAdd)
		if err != nil {
			t.Fatalf("failed to add comment: %v", err)
		}
		// check
		gotBlock, err := task1.Blockers.GetNote(blockID)
		if err != nil {
			t.Fatalf("failed to load block task: %v", err)
		}
		gotComment, err := gotBlock.BlockerNotes.GetNote(commentID)
		if err != nil {
			t.Fatalf("failed to load comment: %v", err)
		}
		if gotComment.Text != commentToAdd {
			t.Errorf("comment not added correctly")
		}
	})

}
