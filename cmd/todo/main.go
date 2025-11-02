package main

import (
	"github.com/alecthomas/kong"
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/internal/storage"
)

var CLI struct {
	Add    commands.AddCmd    `cmd:"" help:"Add a new task"`
	List   commands.ListCmd   `cmd:"" help:"List tasks"`
	Start  commands.StartCmd  `cmd:"" help:"Start or resume a task"`
	Stop   commands.StopCmd   `cmd:"" help:"Stop current task"`
	Switch commands.SwitchCmd `cmd:"" help:"Switch to another task"`
	// Done    DoneCmd    `cmd:"" help:"Mark a task done"`
	// Note    NoteCmd    `cmd:"" help:"Add a note"`
	// Summary SummaryCmd `cmd:"" help:"Show work summary"`
	// Export  ExportCmd  `cmd:"" help:"Export data"`
}

func main() {
	ctx := kong.Parse(&CLI)
	app := storage.NewStore() // loads JSON files
	err := ctx.Run(app)
	ctx.FatalIfErrorf(err)
}
