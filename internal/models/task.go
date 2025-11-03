package models

import "time"

type Task struct {
	ID                      string        `json:"id"`
	Title                   string        `json:"title"`
	Body                    string        `json:"body"`
	Tags                    []Tag         `json:"tags"`
	Priority                int           `json:"priority"` // 1=high,2=med,3=low
	Status                  string        `json:"status"`   // "todo","in_progress","done","blocked"
	EstimatedTaskTimeInDays float32       `json:"estimated_task_time_in_days"`
	EnteredAt               time.Time     `json:"entered_at"`
	WorkLog                 []WorkSession `json:"work_log"`
	Notes                   []Note        `json:"notes"`
	LastUpdated             time.Time     `json:"last_updated"`
}

func (t *Task) IsActive() bool {
	wl := t.WorkLog
	if len(wl) == 0 {
		return false
	}
	lastWl := wl[len(wl)-1]
	if lastWl.EndedAt == nil {
		return true
	} else {
		return false
	}
}
