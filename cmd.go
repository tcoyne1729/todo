package main

import (
	"fmt"
	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
	"log"
	"time"
)

type Cmd interface {
	Run(store *storage.Store) error
}

func pointString(taskID string, currentTask string) string {
	if taskID == currentTask {
		return "-->"
	}
	return "   " // 3 spaces
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
	priority := c.Priority
	if c.Priority == 0 {
		priority = 2 // medium
	}

	newTask := models.Task{
		ID:          store.NewID(),
		Title:       c.Title,
		Body:        c.Body,
		Tags:        tags,
		Priority:    priority,
		Status:      "todo",
		IsActive:    false,
		EnteredAt:   time.Now(),
		LastUpdated: time.Now(),
	}

	// add the tasks
	if err := store.AddTask(newTask); err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}
	log.Println("Added item...")
	if err := store.SaveAll(); err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}
	fmt.Printf("Task added: %s (%s)\n", newTask.Title, newTask.ID)
	return nil
}

func getTask(store *storage.Store, id string) (models.Task, error) {
	for _, task := range store.Tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, fmt.Errorf("no task defined for id = %s", id)

}

type StartCmd struct {
	ID string `arg:"" help:"ID of the task you are starting."`
}

func (c *StartCmd) Run(store *storage.Store) error {
	// update the work
	newTask, err := Start(c.ID, store)
	if err != nil {
		return err
	}
	// garunteed to have a worklog from the Start function as long as no err
	newWorkLog := newTask.WorkLog[len(newTask.WorkLog)-1]

	if err := store.UpdateTask(newTask); err != nil {
		return err
	}
	// update current task
	store.Current = c.ID
	if err := store.SaveAll(); err != nil {
		return fmt.Errorf("failed to start task: %w", err)
	}
	fmt.Printf("%s, Started task: %s\n", newWorkLog.StartedAt.Format(time.Kitchen), newTask.Title)
	return nil
}
func Start(id string, store *storage.Store) (models.Task, error) {
	newWorkLog := models.WorkSession{
		StartedAt: time.Now(),
	}
	newTask, err := getTask(store, id)
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

type StopCmd struct {
	ID         string    `help:"The id of the task you want to stop."`
	CloseTime  time.Time `help:"Time the task is to be closed at."`
	AutoClosed bool      `short:"a" help:"if this task is autoclosed because it overran the threshold."`
}

func (c *StopCmd) Run(store *storage.Store) error {
	task, err := Stop(c.ID, c.CloseTime, c.AutoClosed, store)
	if err != nil {
		return err
	}
	if err := store.UpdateTask(task); err != nil {
		return err
	}
	// unset current task
	store.Current = ""
	if err := store.SaveAll(); err != nil {
		return fmt.Errorf("failed to stop task: %w", err)
	}
	return nil
}

func Stop(id string, closeTime time.Time, autoClosed bool, store *storage.Store) (models.Task, error) {
	if id == "" {
		id = store.Current
		if store.Current == "" {
			return models.Task{}, fmt.Errorf("no ID and no current item")
		}
	}
	newTask, err := getTask(store, id)
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

type SwitchCmd struct {
	ID string `arg:"" help:"The id of the task you want to switch to."`
}

func (c *SwitchCmd) Run(store *storage.Store) error {
	if c.ID == "" {
		return fmt.Errorf("no id provided")
	}
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
		ID: c.ID,
	}
	if err := start.Run(store); err != nil {
		return err
	}

	return nil
}
