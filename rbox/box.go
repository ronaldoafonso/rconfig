/* File: box.go */
/* Description: Operations on remote boxes. */

package rbox

import (
	"database/sql"
	"fmt"
	// PostgreSQL driver
	_ "github.com/lib/pq"
	"github.com/ronaldoafonso/rconfig/rconfig"
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
	Boxname string
	rconfig.Config
}

/* Load config from database. */
func (b *Box) LoadConfig() error {
	query := fmt.Sprintf("SELECT ssid, allowed_macs FROM boxes WHERE boxname = '%s';", b.Boxname)
	var SSID, allowedMACs string

	err := rconfigDb.QueryRow(query).Scan(&SSID, &allowedMACs)
	if err != nil {
		return err
	}

	b.SSID = SSID
	b.AllowedMACs = strings.Split(strings.TrimRight(strings.TrimLeft(allowedMACs, "{"), "}"), ",")
	return nil
}

/* Update box database config. */
func (b *Box) UpdateConfig() error {
	query := fmt.Sprintf("UPDATE boxes SET ssid = '%s', allowed_macs = '%s' WHERE boxname = '%s';",
		b.SSID,
		fmt.Sprintf("%s", b.AllowedMACs),
		b.Boxname)
	_, err := rconfigDb.Exec(query)
	return err
}

/* Get remote box SSID */
func (b *Box) GetRemoteSSID() error {
	uci := []string{
		b.Boxname,
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
func (b Box) SetRemoteSSID() error {
	SSID24 := fmt.Sprintf("wireless.@wifi-iface[0].ssid=%s", b.SSID)
	SSID50 := fmt.Sprintf("wireless.@wifi-iface[1].ssid=%s", b.SSID)
	uci := []string{
		b.Boxname,
		"uci", "set", SSID24, "&&",
		"uci", "set", SSID50, "&&",
		"uci", "commit", "wireless", "&&",
		"/etc/init.d/network", "reload",
	}

	return exec.Command("ssh", uci...).Run()
}

/* Get remote box allowed MACs */
func (b *Box) GetRemoteAllowedMACs() error {
	uci := []string{
		b.Boxname,
		"uci",
		"-q",
		"get",
		"firewall.macs.entry",
	}

	MACs, err := exec.Command("ssh", uci...).Output()
	if err != nil {
		return err
	}

	b.AllowedMACs = []string{}
	for _, MAC := range strings.Split(string(MACs[:len(MACs)-1]), " ") {
		b.AllowedMACs = append(b.AllowedMACs, MAC)
	}

	return nil
}

/* Set remote box allowed MACs */
func (b Box) SetRemoteAllowedMACs() error {
	allowedMACs := "uci delete firewall.macs.entry && "
	for _, MAC := range b.AllowedMACs {
		allowedMACs += "uci add_list firewall.macs.entry='" + MAC + "' && "
	}
	allowedMACs += "uci commit firewall && "
	uci := []string{b.Boxname}

	for _, value := range strings.Split(allowedMACs, " ") {
		uci = append(uci, value)
	}
	uci = append(uci, "/etc/init.d/firewall", "restart")

	return exec.Command("ssh", uci...).Run()
}
