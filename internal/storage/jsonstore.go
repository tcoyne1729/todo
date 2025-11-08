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
	Tags    []models.Tag
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
	tagPath := filepath.Join(s.path, "tags.json")
	allTasks, err := loadJSON[[]models.Task](taskPath)
	if err != nil {
		return err
	}
	allTags, err := loadJSON[[]models.Tag](tagPath)
	if err != nil {
		return err
	}
	curr, err := loadJSON[string](currentPath)
	if err != nil {
		return err
	}
	s.Tasks = allTasks
	s.Current = curr
	s.Tags = allTags
	return nil
}

// save all the store data to disk
func (s *Store) SaveAll() error {
	taskPath := filepath.Join(s.path, "tasks.json")
	currentPath := filepath.Join(s.path, "current.json")
	tagPath := filepath.Join(s.path, "tags.json")
	if err := saveJSON(taskPath, s.Tasks); err != nil {
		return fmt.Errorf("error saving tasks: %w.", err)
	}
	if err := saveJSON(tagPath, s.Tags); err != nil {
		return fmt.Errorf("error saving current tags: %w", err)
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

func (s *Store) AddTag(tag models.Tag) error {
	for _, t := range s.Tags {
		if t.Name == tag.Name {
			return fmt.Errorf("tag %s already exists\n", tag.Name)
		}
	}
	s.Tags = append(s.Tags, tag)
	return nil
}

func (s *Store) DeleteTag(tagName string) error {
	inputTags := s.Tags
	inputNoTags := len(inputTags)
	indexToRemove := -1
	for i, tag := range inputTags {
		// delete the tag from list
		if tag.Name == tagName {
			indexToRemove = i
		}
	}
	if indexToRemove > -1 && indexToRemove < inputNoTags-1 {
		s.Tags = append(s.Tags[:indexToRemove], s.Tags[indexToRemove+1:]...)
		return nil
	} else if indexToRemove > -1 {
		s.Tags = s.Tags[:indexToRemove]
		return nil
	}
	return fmt.Errorf("tag %s does not exist\n", tagName)
}

func (s *Store) GetTag(name string) (models.Tag, error) {
	for _, tag := range s.Tags {
		if tag.Name == name {
			return tag, nil
		}
	}
	return models.Tag{}, fmt.Errorf("tag %s does not exist\n", name)
}

func (s *Store) EditTag(originalName string, newTag models.Tag) error {
	editIndex := -1
	for i, tag := range s.Tags {
		if tag.Name == originalName {
			editIndex = i
			// order and number of tags not changing so modify here
			s.Tags[i] = newTag
		}
	}
	if editIndex < 0 {
		return fmt.Errorf("tag %s does not exist\n", originalName)
	}
	return nil
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
