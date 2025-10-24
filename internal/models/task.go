package models

import "time"

type Task struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Body        string        `json:"body"`
	Area        string        `json:"area"` // project/category
	Tags        []Tag         `json:"tags"`
	Priority    int           `json:"priority"`  // 1=high,2=med,3=low
	Status      string        `json:"status"`    // "todo","in_progress","done","archived"
	IsActive    bool          `json:"is_active"` // current focus
	EnteredAt   time.Time     `json:"entered_at"`
	StartAt     *time.Time    `json:"start_at"` // first start
	EndAt       *time.Time    `json:"end_at"`   // completed time
	WorkLog     []WorkSession `json:"work_log"`
	Notes       []string      `json:"notes"`
	Comments    []Comment     `json:"comments"`
	LastUpdated time.Time     `json:"last_updated"`
}
