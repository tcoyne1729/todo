package models

import (
	"time"

	"github.com/google/uuid"
	genericnotes "github.com/tcoyne1729/todo/internal/generic_notes"
)

type Task struct {
	ID                      string                            `json:"id"`
	Title                   string                            `json:"title"`
	Body                    string                            `json:"body"`
	Tags                    []Tag                             `json:"tags"`
	Priority                int                               `json:"priority"` // 1=high,2=med,3=low
	Status                  string                            `json:"status"`   // "todo","in_progress","done","blocked"
	EstimatedTaskTimeInDays float32                           `json:"estimated_task_time_in_days"`
	EnteredAt               time.Time                         `json:"entered_at"`
	WorkLog                 *genericnotes.Notes[*WorkSession] `json:"work_log"`
	Notes                   *genericnotes.Notes[*Note]        `json:"notes"`
	Blockers                *genericnotes.Notes[*Blocker]     `json:"blocker"`
}

type NewTaskConfig struct {
	ID        string
	Title     string
	Body      string
	Priority  int
	EnteredAt time.Time
	Status    string
}

func NewTask(config NewTaskConfig) *Task {
	var id string
	if config.ID == "" {
		id = uuid.New().String()
	} else {
		id = config.ID
	}
	var status string
	if config.Status == "" {
		status = "todo"
	} else {
		status = config.Status
	}
	enteredAt := time.Now()
	if !config.EnteredAt.IsZero() {
		enteredAt = config.EnteredAt
	}
	return &Task{
		ID:        id,
		Title:     config.Title,
		Body:      config.Body,
		Status:    status,
		Priority:  config.Priority,
		EnteredAt: enteredAt,
		WorkLog:   genericnotes.NewNotes[*WorkSession](),
		Notes:     genericnotes.NewNotes[*Note](),
		Blockers:  genericnotes.NewNotes[*Blocker](),
	}
}

func (t *Task) LastUpdated() time.Time {
	return time.Time{}
}

func (t *Task) IsActive() (bool, error) {
	wl, err := t.WorkLog.GetLast()
	if err != nil {
		return false, err
	}
	active := wl.IsActive()
	return active, nil
}
