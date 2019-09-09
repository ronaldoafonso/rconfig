/* File: box_test.go */
/* Description: Test operations executed on remote boxes. */

package rconfig

import (
	"testing"
)

func TestBoxSetBoxname(t *testing.T) {
	box := Box{}
	box.setBoxname("boxname")

	if box.boxname != "boxname" {
		t.Errorf("setBoxname: Want: %v, got: %v.", "boxname", box.boxname)
	}
}

func TestBoxLoadConfig(t *testing.T) {
	box := Box{}
	box.setBoxname("boxname")

	err := box.loadConfig()

	if err != nil {
		t.Errorf("loadConfig: Got an error [%v].", err)
	}

	allowedMACs := []string{
		"11:11:11:11:11:11",
		"22:22:22:22:22:22",
	}

	if len(box.allowedMACs) != len(allowedMACs) {
		t.Errorf("loadConfig: Wrong length for allowedMACs.")
	}

	for i := range box.allowedMACs {
		if box.allowedMACs[i] != allowedMACs[i] {
			t.Errorf("loadConfig: Wrong MAC. Want: %v, got: %v.",
				allowedMACs[i],
				box.allowedMACs[i])
		}
	}
}

func TestBoxUpdateConfig(t *testing.T) {
	box := Box{}
	box.setBoxname("boxname")
	box.SSID = "new ssid"
	box.allowedMACs = []string{
		"11:11:11:11:11:11",
		"22:22:22:22:22:22",
		"33:33:33:33:33:33",
		"44:44:44:44:44:44",
	}

	err := box.updateConfig()
	if err != nil {
		t.Errorf("updateConfig: Got an error [%v].", err)
	}

	box.loadConfig()

	if box.SSID != "new ssid" {
		t.Errorf("updateConfig: Wrong SSID. Want: %v, got :%v.", "new ssid", box.SSID)
	}
}

func TestGetRemoteSSID(t *testing.T) {
	box := Box{}
	box.setBoxname("boxname")
	box.loadConfig()
	box.SSID = "other ssid"

	if err := box.getRemoteSSID(); err != nil {
		t.Errorf("getRemoteSSID: Got an error [%v].", err)
	}

	if box.SSID != "ssid" {
		t.Errorf("getRemoteSSID: Want: %v, got: %v.", "ssid", box.SSID)
	}
}

func TestSetRemoteSSID(t *testing.T) {
	box := Box{}
	box.setBoxname("boxname")
	box.loadConfig()

	if err := box.setRemoteSSID(); err != nil {
		t.Errorf("setRemoteSSID: Got an error [%v].", err)
	}
}

func TestGetRemoteAllowedMACs(t *testing.T) {
	box := Box{}
	box.setBoxname("boxname")
	box.loadConfig()
	box.allowedMACs = []string{
		"aa:aa:aa:aa:aa:aa",
		"bb:bb:bb:bb:bb:bb",
	}

	if err := box.getRemoteAllowedMACs(); err != nil {
		t.Errorf("getRemoteAllowedMACs: Got an error [%v].", err)
	}

	MAC1 := "11:11:11:11:11:11"
	MAC2 := "22:22:22:22:22:22"

	if box.allowedMACs[0] != MAC1 {
		t.Errorf("getRemoteAllowedMACs. Want: %v, got: %v.", MAC1, box.allowedMACs[0])
	}

	if box.allowedMACs[1] != MAC2 {
		t.Errorf("getRemoteAllowedMACs. Want: %v, got: %v.", MAC2, box.allowedMACs[1])
	}
}

func TestSetRemoteAllowedMACs(t *testing.T) {
	box := Box{}
	box.setBoxname("boxname")
	box.loadConfig()

	if err := box.setRemoteAllowedMACs(); err != nil {
		t.Errorf("setRemoteAllowedMACs: Got an error [%v].", err)
	}
}
