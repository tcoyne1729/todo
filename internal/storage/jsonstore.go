package storage

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/tcoyne1729/todo/internal/models"
	"os"
	"path/filepath"
)

type Store struct {
	Tasks   []models.Task
	Current string
	path    string
}

func NewStore(pathOverride ...string) *Store {
	var p string
	if len(pathOverride) > 0 {
		p = pathOverride[0] // Use the provided path
	} else {
		p = DefaultDir() // Use the default
	}
	s := &Store{path: p}
	s.LoadAll()
	return s
}

func (s *Store) NewID() string {
	id := uuid.New()
	return id.String()
}

func (s *Store) LoadAll() error {
	// load all Store data from disk
	taskPath := filepath.Join(s.path, "tasks.json")
	currentPath := filepath.Join(s.path, "current.json")
	allTasks, err := loadJSON[[]models.Task](taskPath)
	if err != nil {
		return err
	}
	curr, err := loadJSON[string](currentPath)
	if err != nil {
		return err
	}
	s.Tasks = allTasks
	s.Current = curr
	return nil
}

// save all the store data to disk
func (s *Store) SaveAll() error {
	taskPath := filepath.Join(s.path, "tasks.json")
	currentPath := filepath.Join(s.path, "current.json")
	if err := saveJSON(taskPath, s.Tasks); err != nil {
		return fmt.Errorf("error saving tasks: %w.", err)
	}
	if err := saveJSON(currentPath, s.Current); err != nil {
		return fmt.Errorf("error saving current task: %w.", err)
	}
	return nil
}

func (s *Store) ListTasks() []models.Task {
	return s.Tasks

}

func (s *Store) AddTask(task models.Task) error {
	s.Tasks = append(s.Tasks, task)
	return nil
}

func (s *Store) UpdateTask(taskUpdate models.Task) error {
	// remove the task and replace the new one
	for i, task := range s.Tasks {
		if task.ID == taskUpdate.ID {
			// update the task
			s.Tasks[i] = taskUpdate
			return nil
		}
	}
	return fmt.Errorf("No id found for id= %s", taskUpdate.ID)
}

// loadJSON reads JSON from a file and unmarshals it into the provided type T.
func loadJSON[T any](path string) (T, error) {
	var data T // Declare the variable to hold the unmarshaled data

	fileData, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		// If the file doesn't exist, return the zero value of T and no error
		return data, nil
	}
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(fileData, &data); err != nil {
		return data, err
	}
	return data, nil
}

// saveJSON marshals the provided value v and writes it to the file path.
func saveJSON(path string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data to JSON: %w", err)
	}
	// Use os.WriteFile with standard permissions (0644)
	return os.WriteFile(path, data, 0644)
}

func (s *Store) GetTask(id string) (models.Task, error) {
	for _, task := range s.Tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, fmt.Errorf("no task defined for id = %s", id)

}
