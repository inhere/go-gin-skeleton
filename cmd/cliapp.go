package main

import (
	"runtime"

	"github.com/gookit/gcli/v2"
	"github.com/gookit/gcli/v2/builtin"
	"github.com/inhere/go-gin-skeleton/cmd/handler"
)

// for test run: go build ./demo/cliapp.go && ./cliapp
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := gcli.NewApp()
	app.Version = "1.0.3"
	app.Description = "this is my cli application"

	app.SetVerbose(cliapp.VerbDebug)
	// app.DefaultCmd("exampl")

	app.Add(handler.GitCommand())
	// app.Add(cmd.ColorCommand())
	app.Add(builtin.GenShAutoComplete())
	// fmt.Printf("%+v\n", cliapp.CommandNames())
	app.Run()
}
