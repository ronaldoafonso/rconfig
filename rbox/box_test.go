/* File: box_test.go */
/* Description: Test operations executed on remote boxes. */

package rbox

import (
	"github.com/ronaldoafonso/rconfig/rconfig"
	"github.com/ronaldoafonso/rconfig/rmac"
	"testing"
)

func TestBoxLoadConfig(t *testing.T) {
	box := Box{}
	box.Boxname = "boxname"
	err := box.LoadConfig()

	if err != nil {
		t.Errorf("LoadConfig: Got an error [%v].", err)
	}

	config := rconfig.Config{
		SSID:      "ssid",
		LeaseTime: "30m",
		AllowedMACs: []string{
			"11:11:11:11:11:11",
			"22:22:22:22:22:22",
		},
	}

	if !box.Config.IsIgual(config) {
		t.Errorf("LoadConfig: box.Config isn't what it's expected.")
		t.Errorf("Want: %v, %v, %v\n", config.SSID, config.LeaseTime, config.AllowedMACs)
		t.Errorf("Got: %v, %v, %v\n", box.Config.SSID, box.Config.LeaseTime, box.Config.AllowedMACs)
	}
}

func TestBoxUpdateConfig(t *testing.T) {
	box := Box{}
	box.Boxname = "boxname"
	box.SSID = "other ssid"
	box.LeaseTime = "60m"
	box.AllowedMACs = rmac.AllowedMACs{
		"aa:aa:aa:aa:aa:aa",
		"bb:bb:bb:bb:bb:bb",
	}

	err := box.UpdateConfig()
	if err != nil {
		t.Errorf("UpdateConfig: Got an error [%v].", err)
	}

	newBox := Box{}
	newBox.Boxname = "boxname"
	newBox.LoadConfig()

	if !box.Config.IsIgual(newBox.Config) {
		t.Errorf("UpdateConfig: box.Config isn't what it's expected.")
		t.Errorf("Want: %v, %v, %v\n", box.Config.SSID, box.Config.LeaseTime, box.Config.AllowedMACs)
		t.Errorf("Got: %v, %v, %v\n", newBox.Config.SSID, newBox.Config.LeaseTime, newBox.Config.AllowedMACs)
	}
}

func TestGetRemoteSSID(t *testing.T) {
	box := Box{}
	box.Boxname = "boxname"
	box.LoadConfig()
	box.SSID = "other ssid"

	if err := box.GetRemoteSSID(); err != nil {
		t.Errorf("GetRemoteSSID: Got an error [%v].", err)
	}

	if box.SSID != "ssid" {
		t.Errorf("GetRemoteSSID: Want: %v, got: %v.", "ssid", box.SSID)
	}
}

func TestSetRemoteSSID(t *testing.T) {
	box := Box{}
	box.Boxname = "boxname"
	box.LoadConfig()

	if err := box.SetRemoteSSID(); err != nil {
		t.Errorf("SetRemoteSSID: Got an error [%v].", err)
	}
}

func TestGetRemoteLeaseTime(t *testing.T) {
	box := Box{}
	box.Boxname = "boxname"
	box.LoadConfig()
	box.LeaseTime = "60m"

	if err := box.GetRemoteLeaseTime(); err != nil {
		t.Errorf("GetRemoteLeaseTime: Got an error [%v].", err)
	}

	if box.LeaseTime != "30m" {
		t.Errorf("GetRemoteLeaseTime: Want: %v, got: %v.", "30m", box.LeaseTime)
	}
}

func TestSetRemoteLeaseTime(t *testing.T) {
	box := Box{}
	box.Boxname = "boxname"
	box.LoadConfig()

	if err := box.SetRemoteLeaseTime(); err != nil {
		t.Errorf("SetRemoteLeaseTime: Got an error [%v].", err)
	}
}

func TestGetRemoteAllowedMACs(t *testing.T) {
	box := Box{}
	box.Boxname = "boxname"
	box.LoadConfig()
	box.AllowedMACs = rmac.AllowedMACs{
		"aa:aa:aa:aa:aa:aa",
		"bb:bb:bb:bb:bb:bb",
	}

	if err := box.GetRemoteAllowedMACs(); err != nil {
		t.Errorf("GetRemoteAllowedMACs: Got an error [%v].", err)
	}

	MAC1 := "11:11:11:11:11:11"
	MAC2 := "22:22:22:22:22:22"

	if box.AllowedMACs[0] != MAC1 {
		t.Errorf("GetRemoteAllowedMACs. Want: %v, got: %v.", MAC1, box.AllowedMACs[0])
	}

	if box.AllowedMACs[1] != MAC2 {
		t.Errorf("GetRemoteAllowedMACs. Want: %v, got: %v.", MAC2, box.AllowedMACs[1])
	}
}

func TestSetRemoteAllowedMACs(t *testing.T) {
	box := Box{}
	box.Boxname = "boxname"
	box.LoadConfig()

	if err := box.SetRemoteAllowedMACs(); err != nil {
		t.Errorf("SetRemoteAllowedMACs: Got an error [%v].", err)
	}
}
