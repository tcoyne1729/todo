package commands

import (
	"fmt"
	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
	"log"
	"time"
)

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
