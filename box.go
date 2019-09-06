/* File: box.go */
/* Description: Operations on remote boxes. */

package rconfig

import (
	"database/sql"
	"fmt"
	// PostgreSQL driver
	_ "github.com/lib/pq"
	"log"
	"os/exec"
	"strings"
)

var (
	rconfigDb *sql.DB
)

// init box database
func init() {
	connString := "user=rconfig dbname=rconfig password=rconfig " +
		"host=rconfig_db sslmode=disable"
	if db, err := sql.Open("postgres", connString); err != nil {
		log.Fatal(err)
	} else {
		rconfigDb = db
	}
}

// Box ... An OpenWrt box struct
type Box struct {
	boxname string
	Config
}

/* setBoxname: set the name of a box. */
func (b *Box) setBoxname(boxname string) {
	b.boxname = boxname
}

/* Load config from file. */
func (b *Box) loadConfig() error {
	query := fmt.Sprintf("SELECT cfg.ssid, cfg.allowed_macs FROM configs cfg INNER JOIN boxes box USING(config_id) WHERE box.boxname = '%s';", b.boxname)
	var SSID, allowedMACs string

	err := rconfigDb.QueryRow(query).Scan(&SSID, &allowedMACs)
	if err != nil {
		return err
	}

	b.SSID = SSID
	b.allowedMACs = strings.Split(strings.TrimRight(strings.TrimLeft(allowedMACs, "{"), "}"), ",")
	return nil
}

/* Get remote box SSID */
func (b *Box) getRemoteSSID() error {
	uci := []string{
		b.boxname,
		"uci",
		"-q",
		"get",
		"wireless.@wifi-iface[0].ssid",
	}

	SSID, err := exec.Command("ssh", uci...).Output()
	if err != nil {
		return err
	}

	b.SSID = string(SSID[:len(SSID)-1])
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

/* Get remote box allowed MACs */
func (b *Box) getRemoteAllowedMACs() error {
	uci := []string{
		b.boxname,
		"uci",
		"-q",
		"get",
		"firewall.macs.entry",
	}

	MACs, err := exec.Command("ssh", uci...).Output()
	if err != nil {
		return err
	}

	b.allowedMACs = []string{}
	for _, MAC := range strings.Split(string(MACs[:len(MACs)-1]), " ") {
		b.allowedMACs = append(b.allowedMACs, MAC)
	}

	return nil
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
