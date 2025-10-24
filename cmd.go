package main

import (
	"fmt"
	"log"
	"time"
	"todo/internal/models"
	"todo/internal/storage"
)

type ListCmd struct{}

func (l *ListCmd) Run(store *storage.Store) error {
	allTasks := store.ListTasks()
	if len(allTasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	for i, t := range allTasks {
		fmt.Printf("%d: [%s] %s\n", i, t.Status, t.Title)
	}
	return nil
}

type AddCmd struct {
	Title string `arg:"" help:"Title of new task."`
	// optional
	Body     string `help:"Body of new task. This is a longer description."`
	Area     string `short:"a" help:"Area (field) of new task."`
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
		Area:        c.Area,
		Tags:        tags,
		Priority:    c.Priority,
		Status:      "todo",
		IsActive:    false,
		EnteredAt:   time.Now(),
		StartAt:     nil,
		EndAt:       nil,
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
