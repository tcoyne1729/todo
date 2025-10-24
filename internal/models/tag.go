package models

type Tag struct {
	Name    string `json:"name"`
	Context string `json:"context"`
	Colour  string `json:"colour,omitempty"`
}
