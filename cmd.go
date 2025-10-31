package main

import (
	"fmt"
	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
	"log"
	"time"
)

func pointString(taskId string, currentTask string) string {
	pad := "   " // 3 spaces
	if taskId != currentTask {
		return pad
	} else {
		return "-->"
	}
}

type ListCmd struct{}

func (l *ListCmd) Run(store *storage.Store) error {
	allTasks := store.ListTasks()
	if len(allTasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}
	currentTask := store.Current

	for i, t := range allTasks {
		fmt.Printf("%d: %s [%s] %s %s\n", i, pointString(t.ID, currentTask), t.Status, t.Title, t.ID)
	}
	return nil
}

type AddCmd struct {
	Title string `arg:"" help:"Title of new task."`
	// optional
	Body     string `help:"Body of new task. This is a longer description."`
	Priority int    `short:"p" help:"The priority of the tasks, 1 meaning high, 2 meaning medium and 3 meaning low."`
	Tag      string `short:"t" help:"Tags to add to the task."`
}

func (c *AddCmd) Run(store *storage.Store) error {
	var tags = []models.Tag{}
	if c.Tag != "" {
		tags = append(tags, models.Tag{
			Name: c.Tag,
		})
	}

	newTask := models.Task{
		ID:          store.NewID(),
		Title:       c.Title,
		Body:        c.Body,
		Tags:        tags,
		Priority:    c.Priority,
		Status:      "todo",
		IsActive:    false,
		EnteredAt:   time.Now(),
		LastUpdated: time.Now(),
	}

	// add the tasks
	if err := store.AddTask(newTask); err != nil {
		return fmt.Errorf("Failed to add task: %w.", err)
	}
	log.Println("Added item..")
	if err := store.SaveAll(); err != nil {
		return fmt.Errorf("Failed to add task: %w.", err)
	}
	fmt.Printf("Task added: %s (%s)", newTask.Title, newTask.ID)
	return nil
}

func getTask(store *storage.Store, id string) (models.Task, error) {
	taskUndefined := true
	var newTask models.Task
	for _, task := range store.Tasks {
		if task.ID == id {
			newTask = task
			taskUndefined = false
		}
	}
	if taskUndefined {
		return models.Task{}, fmt.Errorf("No task defined for id = %s", id)
	}
	return newTask, nil

}

type StartCmd struct {
	ID string `arg:"" help:"ID of the task you are starting."`
	// summary
	Comment string `help:"any additional information."`
}

func (c *StartCmd) Run(store *storage.Store) error {
	// update the work
	newWorkLog := models.WorkSession{
		StartedAt: time.Now(),
	}
	newTask, err := getTask(store, c.ID)
	if err != nil {
		return err
	}
	if newTask.Status == "done" {
		return fmt.Errorf("Can't work on a completed task.")
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
			return fmt.Errorf("Task already started. id = %s", c.ID)
		}
	}
	newTask.WorkLog = append(newTask.WorkLog, newWorkLog)
	if err := store.UpdateTask(newTask); err != nil {
		return err
	}
	// update current task
	store.Current = c.ID
	if err := store.SaveAll(); err != nil {
		return fmt.Errorf("Failed to start task: %w.", err)
	}
	fmt.Printf("%s, Started task: %s", newWorkLog.StartedAt, newTask.Title)
	return nil
}

type StopCmd struct {
	ID         string    `arg:"" help:"The id of the task you want to stop."`
	CloseTime  time.Time `help:"Time the task is to be closed at."`
	AutoClosed bool      `short:"a" help:"if this task is autoclosed because it overran the threshold."`
}

func (c *StopCmd) Run(store *storage.Store) error {
	newTask, err := getTask(store, c.ID)
	if err != nil {
		return err
	}
	// check if the session has already been started
	noWorklogs := len(newTask.WorkLog)
	lastWorkLog := newTask.WorkLog[noWorklogs-1]
	if lastWorkLog.EndedAt != nil {
		return fmt.Errorf("The task has not been started, so cannot be ended. id = %s", c.ID)
	}
	var endTime time.Time
	if c.CloseTime.IsZero() {
		endTime = time.Now()
	} else {
		endTime = c.CloseTime
	}
	newWorkLog := models.WorkSession{
		StartedAt:  lastWorkLog.StartedAt,
		EndedAt:    &endTime,
		AutoClosed: c.AutoClosed,
	}
	newTask.WorkLog[noWorklogs-1] = newWorkLog
	// unset current task
	store.Current = ""
	if err := store.SaveAll(); err != nil {
		return fmt.Errorf("Failed to stop task: %w.", err)
	}
	return nil
}

type SwitchCmd struct {
	ID string `arg:"" help:"The id of the task you want to switch to."`
	// optional
	Comment string `help:"any additional information."`
}

func (c *SwitchCmd) Run(store *storage.Store) error {
	// get the task
	switchToTask, err := getTask(store, c.ID)
	if err != nil {
		return err
	}
	// if the task had been previously started, then stop it
	if len(switchToTask.WorkLog) > 0 {
		needToAutoClose := false
		closeTime := time.Now()
		lastWorkLog := switchToTask.WorkLog[len(switchToTask.WorkLog)-1]
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
				ID:         c.ID,
				CloseTime:  closeTime,
				AutoClosed: needToAutoClose,
			}
			if err := stop.Run(store); err != nil {
				return err
			}
		}
	}
	// start new task
	start := StartCmd{
		ID:      c.ID,
		Comment: c.Comment,
	}
	if err := start.Run(store); err != nil {
		return err
	}

	return nil
}
