package steps

import (
	"github.com/mitchellh/multistep"
	"github.com/bfirsh/virtualdocker/virtualbox"
)

type StartVM struct {}

func (s *StartVM) Run(state multistep.StateBag) multistep.StepAction {
	driver := state.Get("driver").(virtualbox.Driver)

	isRunning, err := driver.IsRunning("virtualdocker")
	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	if !isRunning {
		err = driver.VBoxManage("startvm", "virtualdocker")
		if err != nil {
			state.Put("error", err)
			return multistep.ActionHalt
		}
	}




	return multistep.ActionContinue
}

func (s *StartVM) Cleanup(multistep.StateBag) {
}
