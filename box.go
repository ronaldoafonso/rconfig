/* File: box.go */
/* Description: Operations on remote boxes. */

package rconfig

import (
	"fmt"
	"os/exec"
	"strings"
)

type Box struct {
	boxname string
	config  boxConfig
}

/* Init a remote box struct. */
func initBox() Box {
	return Box{}
}

/* Load config from file. */
// TODO: It should be done with a database.
func (b *Box) loadConfig() error {
	b.boxname = "boxname"
	b.config = initBoxConfig()
	b.config.ssid = "ssid"
	b.config.allowedMacs = []string{
		"11:11:11:11:11:11",
		"22:22:22:22:22:22",
	}
	return nil
}

/* Get remote SSID of box */
func (b *Box) getRemoteSsid() (string, error) {
	buffer, err := b.execRemoteUCIGet("wireless.@wifi-iface[0].ssid")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", string(buffer)), nil
}

/* Get remote allowed MACs of box */
func (b Box) getRemoteAllowedMacs() ([]string, error) {
	buffer, err := b.execRemoteUCIGet("firewall.macs.entry")
	if err != nil {
		return nil, nil
	}

	return strings.Split(string(buffer), " "), nil
}

/* Execute remote UCI command */
func (b Box) execRemoteUCIGet(param string) ([]byte, error) {
	cmd := exec.Command(
		"ssh",
		b.boxname,
		"uci",
		"get",
		param,
	)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	buffer := make([]byte, 1024)
	nRead, err := stdout.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer[:nRead-1], nil
}
