/* File: box.go */
/* Description: Operations on remote boxes. */

package rconfig

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type remoteBox struct {
	boxname string
	config  boxConfig
}

/* Init a remote box struct. */
func initBox(boxname string) remoteBox {
	return remoteBox{
		boxname: boxname,
		config:  boxConfig{},
	}
}

/* Get remote SSID of box */
func (box remoteBox) getRemoteSsid() string {
	cmd := exec.Command(
		"ssh",
		boxname,
		"uci",
		"get",
		"wireless.@wifi-iface[0].ssid",
	)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 32)
	nRead, err := stdout.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%v", string(buffer[:nRead-1]))
}

/* Get remote allowed MACs of box */
func (box remoteBox) getRemoteAllowedMacs() []string {
	cmd := exec.Command(
		"ssh",
		boxname,
		"uci",
		"get",
		"firewall.macs.entry",
	)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 1024)
	nRead, err := stdout.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(buffer[:nRead-1]), " ")
}
