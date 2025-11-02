package commands

import (
	"fmt"
	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
	"time"
)

func Stop(id string, closeTime time.Time, autoClosed bool, store *storage.Store) (models.Task, error) {
	if id == "" {
		id = store.Current
		if store.Current == "" {
			return models.Task{}, fmt.Errorf("no ID and no current item")
		}
	}
	newTask, err := store.GetTask(id)
	if err != nil {
		return models.Task{}, err
	}
	// check if the session has already been started
	noWorklogs := len(newTask.WorkLog)
	if noWorklogs == 0 {
		return models.Task{}, fmt.Errorf("task has no active work sessions: id = %s", id)
	}
	lastWorkLog := newTask.WorkLog[noWorklogs-1]
	if lastWorkLog.EndedAt != nil {
		return models.Task{}, fmt.Errorf("last task has already been ended: id = %s", id)
	}
	var endTime time.Time
	if closeTime.IsZero() {
		endTime = time.Now()
	} else {
		endTime = closeTime
	}
	newWorkLog := models.WorkSession{
		StartedAt:  lastWorkLog.StartedAt,
		EndedAt:    &endTime,
		AutoClosed: autoClosed,
	}
	newTask.WorkLog[noWorklogs-1] = newWorkLog
	return newTask, nil
}

func Start(id string, store *storage.Store) (models.Task, error) {
	newWorkLog := models.WorkSession{
		StartedAt: time.Now(),
	}
	newTask, err := store.GetTask(id)
	if err != nil {
		return models.Task{}, err
	}
	if newTask.Status == "done" {
		return models.Task{}, fmt.Errorf("can't work on a completed task")
	}
	if newTask.Status == "todo" {
		// must be in progress
		newTask.Status = "in_progress"
	}
	if !newTask.IsActive {
		newTask.IsActive = true
	}
	// need to deal with the case where the worksession
	// is open then we need to return an error saying the
	// session is already active
	noWorkLogs := len(newTask.WorkLog)
	if noWorkLogs > 0 {
		lastWorkLog := newTask.WorkLog[len(newTask.WorkLog)-1]
		if lastWorkLog.EndedAt == nil {
			// task already started
			return models.Task{}, fmt.Errorf("task already started. id = %s", id)
		}
	}
	newTask.WorkLog = append(newTask.WorkLog, newWorkLog)
	return newTask, nil
}

func PointString(taskID string, currentTask string) string {
	if taskID == currentTask {
		return "-->"
	}
	return "   " // 3 spaces
}
