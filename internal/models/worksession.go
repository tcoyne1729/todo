package models

import (
	genericnotes "github.com/tcoyne1729/todo/internal/generic_notes"
)

type WorkSession struct {
	genericnotes.EntryBase
	AutoClosed bool `json:"auto_closed"`
}

// func NewWorkSession(createTime time.Time) *WorkSession {
// 	if createTime.IsZero() {
// 		createTime = time.Now()
// 	}
// 	return &WorkSession{
// 		EntryBase: &genericnotes.EntryBase{
// 			CreateTime: createTime,
// 		},
// 	}
// }
