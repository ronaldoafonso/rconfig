package rbox

import (
	"os"
	"os/exec"
	"strings"
)

type RemoteBox struct {
	boxname string
}

func (b *RemoteBox) GetConfig() error {
	cmd := exec.Command("ssh", b.boxname, "uci show")
	config, err := cmd.Output()
	if err != nil {
		return err
	}

	file, err := os.Create(b.boxname)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(config); err != nil {
		return err
	}
	return nil
}

func (b *RemoteBox) GetSSIDs() ([]string, error) {
	SSIDs := []string{}

	commands := []string{
		"uci get wireless.@wifi-iface[0].ssid",
		"uci get wireless.@wifi-iface[1].ssid",
	}
	for _, command := range commands {
		cmd := exec.Command("ssh", b.boxname, command)
		SSID, err := cmd.Output()
		if err != nil {
			return []string{}, err
		}
		SSIDs = append(SSIDs, strings.TrimSpace(string(SSID)))
	}

	return SSIDs, nil
}

func NewRBox(boxname string) RemoteBox {
	return RemoteBox{
		boxname: boxname,
	}
}
