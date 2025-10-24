package models

import "time"

type WorkSession struct {
	StartedAt  time.Time  `json:"started_at"`
	EndedAt    *time.Time `json:"ended_at"`
	Summary    string     `json:"summary"`
	AutoClosed bool       `json:"auto_closed"`
}
