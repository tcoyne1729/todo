package genericnotes

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/google/uuid"
)

// we use some form of notes in several places
// capture the logic here

type NoteEntry interface {
	GetID() string
	GetCreationTime() time.Time
	SetCompleteTime(t time.Time)
	IsActive() bool
	SetId(id string) string
	SetCreationTime(createTime time.Time)
	SetText(text string)
}

type Notes[T NoteEntry] struct {
	Data []T
}

func NewNotes[T NoteEntry]() *Notes[T] {
	return &Notes[T]{
		Data: make([]T, 0),
	}
}

// custom JSON Marshaler and unmarshaler so we can control the JSON respresentation

// Custom wrapper. Call the standard json.Marshal but only on the Data object to prevent
// everything being wrapped in a Data: [] in the JSON output
func (n *Notes[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Data)
}

// Custom wrapper. Call the standard json.Unmarshal but put the output into the Data object
func (n *Notes[T]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &n.Data)
}

// Common functions
type NewConfig struct {
	ID         string
	CreateTime time.Time
	Text       string
}

func (n *Notes[T]) New(config NewConfig) error {
	var zero T
	t := reflect.TypeOf(zero)
	if t.Kind() != reflect.Pointer {
		return fmt.Errorf("generic type T must be a pointer to a struct")
	}

	newValue := reflect.New(t.Elem())
	// convert back to an interface type T
	newEntry, ok := newValue.Interface().(T)
	if !ok {
		fmt.Errorf("reflection failed to assert type %T to %T", newValue.Interface(), zero)
	}

	newEntry.SetId(config.ID)
	newEntry.SetCreationTime(config.CreateTime)
	newEntry.SetText(config.Text)

	n.Data = append(n.Data, newEntry)

	return nil
}

func (n *Notes[T]) GetLast() (T, error) {
	var output T
	if len(n.Data) == 0 {
		return output, nil
	}
	output = n.Data[0]
	for i, note := range n.Data {
		if output.GetCreationTime().Before(note.GetCreationTime()) {
			output = n.Data[i]
		}
	}
	return output, nil
}

func (n *Notes[T]) Insert(entry T) error {
	n.Data = append(n.Data, entry)
	return nil
}

func (n *Notes[T]) MarkComplete(id string) error {
	for i, e := range n.Data {
		if e.GetID() == id {
			n.Data[i].SetCompleteTime(time.Now())
			return nil
		}
	}
	return fmt.Errorf("no entry found for id: %s", id)
}

func (n *Notes[T]) Delete(id string) error {
	delID := -1
	for i, e := range n.Data {
		if e.GetID() == id {
			delID = i
			break
		}
	}
	switch {
	case delID < 0:
		return fmt.Errorf("no entry found for id: %s", id)
	case delID == len(n.Data):
		n.Data = n.Data[delID-1:]
	default:
		n.Data = append(n.Data[:delID], n.Data[delID+1:]...)
	}
	return nil
}

func (n *Notes[T]) GetNote(id string) (*T, error) {
	for i, note := range n.Data {
		if note.GetID() == id {
			return &n.Data[i], nil
		}
	}
	return nil, fmt.Errorf("no note with id %s", id)
}

func (n *Notes[T]) ListAll() ([]T, error) { return n.Data, nil }
func (n *Notes[T]) ListActive() ([]T, error) {
	output := []T{}
	for _, e := range n.Data {
		if e.IsActive() {
			output = append(output, e)
		}
	}
	return output, nil
}

type EntryBase struct {
	ID           string     `json:"id"`
	CreateTime   time.Time  `json:"create_time"`
	CompleteTime *time.Time `json:"complete_time"`
	Text         string     `json:"text"`
}

func (e *EntryBase) GetID() string               { return e.ID }
func (e *EntryBase) SetCompleteTime(t time.Time) { e.CompleteTime = &t }
func (e *EntryBase) IsActive() bool {
	if e == nil {
		return false
	}
	return e.CompleteTime == nil
}
func (e *EntryBase) SetId(id string) string {
	if id == "" {
		id = uuid.New().String()
	}
	e.ID = id
	return id
}
func (e *EntryBase) SetCreationTime(createTime time.Time) {
	if createTime.IsZero() {
		createTime = time.Now()
	}
	e.CreateTime = createTime
}
func (e *EntryBase) SetText(text string) {
	e.Text = text
}
func (e *EntryBase) GetCreationTime() time.Time {
	return e.CreateTime
}
