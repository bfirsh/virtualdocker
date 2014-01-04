package steps

import (
	"github.com/mitchellh/multistep"
	"github.com/bfirsh/virtualdocker/virtualbox"
)

type CreateVM struct {}

func (s *CreateVM) Run(state multistep.StateBag) multistep.StepAction {
	driver := state.Get("driver").(virtualbox.Driver)
	isoPath := state.Get("isoPath").(string)

	isCreated, err := driver.IsCreated("virtualdocker")
	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	if isCreated {
		return multistep.ActionContinue
	}

	err = driver.VBoxManage("createvm",
		"--name", "virtualdocker",
		"--ostype", "Linux_64",
		"--register")
	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	err = driver.VBoxManage("modifyvm", "virtualdocker",
		"--memory", "1024",
		"--cpus", "2")
	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	err = driver.VBoxManage("storagectl", "virtualdocker",
		"--name", "IDE1",
		"--add", "ide")
	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	err = driver.VBoxManage("storageattach", "virtualdocker",
		"--storagectl", "IDE1",
		"--port", "0",
		"--device", "1",
		"--type", "dvddrive",
		"--medium", isoPath)
	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	err = driver.VBoxManage("storageattach", "virtualdocker",
		"--storagectl", "IDE1",
		"--port", "0",
		"--device", "1",
		"--type", "dvddrive",
		"--medium", isoPath)
	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}
	return multistep.ActionContinue
}

func (s *CreateVM) Cleanup(multistep.StateBag) {
}
