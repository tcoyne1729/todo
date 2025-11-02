package commands

import (
	"fmt"
	"time"

	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
)

type SwitchCmd struct {
	ID string `arg:"" help:"The id of the task you want to switch to."`
}

func (c *SwitchCmd) Run(store *storage.Store) error {
	if c.ID == "" {
		return fmt.Errorf("no id provided for the task you want to switch to")
	}
	if store.Current != "" {
		// stop the current task
		currentTask, err := store.GetTask(store.Current)
		if err != nil {
			return err
		}
		if err := stopIfStarted(&currentTask, store); err != nil {
			return err
		}
	}
	// get the task
	switchToTask, err := store.GetTask(c.ID)
	if err != nil {
		return err
	}
	// if the switch to task had been previously started, then stop it
	if err := stopIfStarted(&switchToTask, store); err != nil {
		return err
	}

	// start new task
	start := StartCmd{
		ID: c.ID,
	}
	if err := start.Run(store); err != nil {
		return err
	}

	return nil
}

func stopIfStarted(task *models.Task, store *storage.Store) error {

	if len(task.WorkLog) > 0 {
		needToAutoClose := false
		closeTime := time.Now()
		lastWorkLog := task.WorkLog[len(task.WorkLog)-1]
		lastStartedTime := lastWorkLog.StartedAt
		if lastWorkLog.EndedAt == nil {
			timeNow := time.Now()
			// autoclose if the task was started on the prior date
			if !(lastStartedTime.Year() == timeNow.Year() && lastStartedTime.Month() == timeNow.Month() && lastStartedTime.Day() == timeNow.Day()) {
				// auto close
				needToAutoClose = true
				closeTime = time.Date(lastStartedTime.Year(), lastStartedTime.Month(), lastStartedTime.Day(), 0, 0, 0, 0, lastWorkLog.StartedAt.Location()).AddDate(0, 0, 1)
			}
			// stop the task and start a new one
			stop := StopCmd{
				ID:         task.ID,
				CloseTime:  closeTime,
				AutoClosed: needToAutoClose,
			}
			if err := stop.Run(store); err != nil {
				return err
			}
		}
	}
	return nil
}
