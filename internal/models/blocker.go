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
	BlockerNotes genericnotes.Notes[*BlockerNote]
	BlockedBy    string
}
