package main

import (
	"runtime"

	"github.com/gookit/gcli/v2/builtin"
	"github.com/inhere/go-gin-skeleton/cli/cmd"
)

// for test run: go build ./demo/cliapp.go && ./cliapp
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cliapp.NewApp()
	app.Version = "1.0.3"
	app.Description = "this is my cli application"

	app.SetVerbose(cliapp.VerbDebug)
	// app.DefaultCmd("exampl")

	app.Add(cmd.GitCommand())
	// app.Add(cmd.ColorCommand())
	app.Add(builtin.GenShAutoComplete())
	// fmt.Printf("%+v\n", cliapp.CommandNames())
	app.Run()
}
