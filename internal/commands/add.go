package commands

import (
	"fmt"
	"log"
	"time"

	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/internal/storage"
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

	newTask := models.NewTask(models.NewTaskConfig{
		Title:     c.Title,
		Body:      c.Body,
		Priority:  priority,
		EnteredAt: time.Now(),
	},
	)
	newTask.Tags = tags

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
