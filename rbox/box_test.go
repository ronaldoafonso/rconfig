/* File: box_test.go */
/* Description: Test operations executed on remote boxes. */

package rbox

import (
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

	allowedMACs := rmac.AllowedMACs{
		"11:11:11:11:11:11",
		"22:22:22:22:22:22",
	}

	if len(box.AllowedMACs) != len(allowedMACs) {
		t.Errorf("LoadConfig: Wrong length for AllowedMACs.")
	}

	for i := range box.AllowedMACs {
		if box.AllowedMACs[i] != allowedMACs[i] {
			t.Errorf("LoadConfig: Wrong MAC. Want: %v, got: %v.",
				allowedMACs[i],
				box.AllowedMACs[i])
		}
	}
}

func TestBoxUpdateConfig(t *testing.T) {
	box := Box{}
	box.Boxname = "boxname"
	box.SSID = "ssid"
	box.AllowedMACs = rmac.AllowedMACs{
		"11:11:11:11:11:11",
		"22:22:22:22:22:22",
	}

	err := box.UpdateConfig()
	if err != nil {
		t.Errorf("UpdateConfig: Got an error [%v].", err)
	}

	box.LoadConfig()

	if box.SSID != "ssid" {
		t.Errorf("UpdateConfig: Wrong SSID. Want: %v, got :%v.", "ssid", box.SSID)
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
