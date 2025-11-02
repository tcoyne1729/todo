package commands

import (
	"github.com/tcoyne1729/todo/internal/storage"
)

type Cmd interface {
	Run(store *storage.Store) error
}
