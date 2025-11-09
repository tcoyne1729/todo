package commands

import "github.com/tcoyne1729/todo/internal/storage"

// update a task

type UpdateCmd struct {
	ID       string `help:"ID to update"`
	Title    string `help:"new title"`
	Body     string `help:"new body"`
	Priority int    `help:"new priority (should be 1, 2 or 3)"`
}

func (u *UpdateCmd) Run(store *storage.Store) error {
	existingTask, err := store.GetTask(u.ID)
	if err != nil {
		return err
	}
	if u.Title != "" {
		existingTask.Title = u.Title
	}
	if u.Body != "" {
		existingTask.Body = u.Body
	}
	if u.Priority != 0 {
		existingTask.Priority = u.Priority
	}
	if err = store.UpdateTask(existingTask); err != nil {
		return err
	}
	return nil
}
