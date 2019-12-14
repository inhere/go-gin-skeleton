package main

import (
	"fmt"
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

	// app.SetVerbose(gcli.VerbDebug)
	// app.DefaultCmd("exampl")

	app.Add(handler.GitCommand())
	// app.Add(cmd.ColorCommand())
	app.Add(builtin.GenAutoCompleteScript())

	app.Add(&gcli.Command{
		Name:   "test",
		UseFor: "an test command",
		Func: func(c *gcli.Command, args []string) error {
			fmt.Println("hello")
			return nil
		},
	})

	// fmt.Printf("%+v\n", gcli.CommandNames())
	app.Run()
}
