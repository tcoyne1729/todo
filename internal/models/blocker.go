package models

import (
	// "time"

	genericnotes "github.com/tcoyne1729/todo/internal/generic_notes"
)

type BlockerNote struct {
	genericnotes.EntryBase
}

type Blocker struct {
	genericnotes.EntryBase
	BlockerNotes []*BlockerNote
	BlockedBy    string
}

// func MyTest() error {
// 	notes := genericnotes.NewNotes[*BlockerNote]()
// 	blockerNote := &BlockerNote{
// 		EntryBase: &genericnotes.EntryBase{
// 			ID:         "1",
// 			CreateTime: time.Now(),
// 			Text:       "Need to get a response from Bob",
// 		},
// 	}
//
// 	if err := notes.Insert(blockerNote); err != nil {
// 		return err
// 	}
// 	return nil
// }
