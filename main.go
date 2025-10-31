package main

import (
	"github.com/alecthomas/kong"
	"github.com/tcoyne1729/todo/internal/storage"
)

var CLI struct {
	Add    AddCmd    `cmd:"" help:"Add a new task"`
	List   ListCmd   `cmd:"" help:"List tasks"`
	Start  StartCmd  `cmd:"" help:"Start or resume a task"`
	Stop   StopCmd   `cmd:"" help:"Stop current task"`
	Switch SwitchCmd `cmd:"" help:"Switch to another task"`
	// Done    DoneCmd    `cmd:"" help:"Mark a task done"`
	// Note    NoteCmd    `cmd:"" help:"Add a note"`
	// Comment CommentCmd `cmd:"" help:"Add a comment"`
	// Summary SummaryCmd `cmd:"" help:"Show work summary"`
	// Export  ExportCmd  `cmd:"" help:"Export data"`
}

func main() {
	ctx := kong.Parse(&CLI)
	app := storage.NewStore() // loads JSON files
	err := ctx.Run(app)
	ctx.FatalIfErrorf(err)
}
