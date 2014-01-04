package commands

import (
	"github.com/mitchellh/cli"
	"github.com/mitchellh/multistep"
	"github.com/bfirsh/virtualdocker/steps"
	"github.com/bfirsh/virtualdocker/virtualbox"
	"strings"
)

type UpCommand struct {
	Ui cli.Ui
}

func (c *UpCommand) Help() string {
	helpText := `
Usage: virtualdocker up

  Brings up a Docker virtual machine.
`
	return strings.TrimSpace(helpText)
}

func (c *UpCommand) Run(args []string) int {
	driver, err := virtualbox.NewDriver()
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	steps := []multistep.Step{
		new(steps.DownloadIso),
		new(steps.CreateVM),
		new(steps.StartVM),
	}

	state := new(multistep.BasicStateBag)
	state.Put("driver", driver)
	state.Put("ui", c.Ui)

	runner := &multistep.BasicRunner{Steps: steps}
    runner.Run(state)

	
	if err, ok := state.GetOk("error"); ok {
		c.Ui.Error(err.(error).Error())
		return 1
	}

	return 0
}

func (c *UpCommand) Synopsis() string {
	return "Brings up a Docker virtual machine"
}
