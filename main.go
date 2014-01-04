package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"github.com/bfirsh/virtualdocker/commands"
	"os"
)


func main() {
        os.Exit(realMain())
}

func realMain() int {
	ui := &cli.BasicUi{Writer: os.Stdout}

	Commands := map[string]cli.CommandFactory{
		"up": func() (cli.Command, error) {
			return &commands.UpCommand{
				Ui: ui,
			}, nil
		},
	}

	cli := &cli.CLI{
		Args: os.Args[1:],
		Commands: Commands,
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}
