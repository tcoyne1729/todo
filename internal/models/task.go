package models

import "time"

type Task struct {
	ID                      string        `json:"id"`
	Title                   string        `json:"title"`
	Body                    string        `json:"body"`
	Tags                    []Tag         `json:"tags"`
	Priority                int           `json:"priority"`  // 1=high,2=med,3=low
	Status                  string        `json:"status"`    // "todo","in_progress","done","blocked"
	IsActive                bool          `json:"is_active"` // current focus
	EstimatedTaskTimeInDays float32       `json:"estimated_task_time_in_days"`
	EnteredAt               time.Time     `json:"entered_at"`
	WorkLog                 []WorkSession `json:"work_log"`
	Notes                   []string      `json:"notes"`
	LastUpdated             time.Time     `json:"last_updated"`
}
