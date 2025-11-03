package models

import (
	"time"
)

type Note struct {
	TimeStamp time.Time
	Note      string
}
