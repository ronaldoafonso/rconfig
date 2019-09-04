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
	Config
}

/* setBoxname: set the name of a box. */
func (b *Box) setBoxname(boxname string) {
	b.boxname = boxname
}

/* Load config from file. */
// TODO: It should be done with a database.
func (b *Box) loadConfig() error {
	b.SSID = "ssid"
	b.allowedMACs = []string{
		"11:11:11:11:11:11",
		"22:22:22:22:22:22",
	}
	return nil
}

/* Set remote box SSID */
func (b Box) setRemoteSSID() error {
	SSID24 := fmt.Sprintf("wireless.@wifi-iface[0].ssid=%s", b.SSID)
	SSID50 := fmt.Sprintf("wireless.@wifi-iface[1].ssid=%s", b.SSID)
	uci := []string{
		b.boxname,
		"uci", "set", SSID24, "&&",
		"uci", "set", SSID50, "&&",
		"uci", "commit", "wireless", "&&",
		"/etc/init.d/network", "reload",
	}

	return exec.Command("ssh", uci...).Run()
}

/* Set remote box allowed MACs */
func (b Box) setRemoteAllowedMACs() error {
	allowedMACs := "uci delete firewall.macs.entry && "
	for _, MAC := range b.allowedMACs {
		allowedMACs += "uci add_list firewall.macs.entry='" + MAC + "' && "
	}
	allowedMACs += "uci commit firewall && "
	uci := []string{b.boxname}

	for _, value := range strings.Split(allowedMACs, " ") {
		uci = append(uci, value)
	}
	uci = append(uci, "/etc/init.d/firewall", "restart")

	return exec.Command("ssh", uci...).Run()
}
