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
	Done   commands.DoneCmd   `cmd:"" help:"Mark a task done"`
	Note   commands.NoteCmd   `cmd:"" help:"Add a note"`
	// backfill -> id, start, end (optional). If end not set then you are doing switch with a start datetime set
	// status -> info about the current task
	// Tag -> add a tag
	// waiting for person -> who, why, starttime, completetime
	// release -> complete the above

	// Summary SummaryCmd `cmd:"" help:"Show work summary"`
	// Export  ExportCmd  `cmd:"" help:"Export data"`
}

func main() {
	ctx := kong.Parse(&CLI)
	app := storage.NewStore() // loads JSON files
	err := ctx.Run(app)
	ctx.FatalIfErrorf(err)
}
