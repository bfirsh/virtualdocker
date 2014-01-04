package steps

import (
	"fmt"
	"github.com/mitchellh/cli"
	"github.com/mitchellh/multistep"
	"io"
	"net/http"
	"os"
	"os/user"
	"path"
)

type DownloadIso struct {}

func (s *DownloadIso) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(cli.Ui)

	usr, err := user.Current()
	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	isoPath := path.Join(usr.HomeDir, ".virtualdocker", "boot2docker", "boot2docker-0.3.0.iso")
	state.Put("isoPath", isoPath)

	// TODO: checksum downloaded file to ensure it is valid/complete
	if _, err := os.Stat(isoPath); err == nil {
		return multistep.ActionContinue
	}

	url := "https://github.com/steeve/boot2docker/releases/download/v0.3.0/boot2docker.iso"

	ui.Output(fmt.Sprintf("Downloading %s to %s", url, isoPath))

	if err := os.MkdirAll(path.Dir(isoPath), 0755); err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	out, err := os.Create(isoPath)
	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}
	defer resp.Body.Close()

	io.Copy(out, resp.Body)

	return multistep.ActionContinue
}

func (s *DownloadIso) Cleanup(multistep.StateBag) {
}
