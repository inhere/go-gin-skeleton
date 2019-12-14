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

	cli := gcli.NewApp()
	cli.Version = "1.0.3"
	cli.Description = "this is my cli application"

	// cli.SetVerbose(gcli.VerbDebug)
	// cli.DefaultCmd("exampl")

	cli.Add(handler.GitCommand())
	// cli.Add(cmd.ColorCommand())
	cli.Add(builtin.GenAutoCompleteScript())

	cli.Add(&gcli.Command{
		Name:   "test",
		UseFor: "an test command",
		Func: func(c *gcli.Command, args []string) error {
			fmt.Println("hello")
			return nil
		},
	})

	// fmt.Printf("%+v\n", gcli.CommandNames())
	cli.Run()
}
