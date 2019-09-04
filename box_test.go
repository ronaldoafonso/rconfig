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

func TestSetRemoteSSID(t *testing.T) {
	box := Box{}
	box.setBoxname("boxname")
	box.loadConfig()

	if err := box.setRemoteSSID(); err != nil {
		t.Errorf("setRemoteSSID: Got an error [%v].", err)
	}
}

func TestSetRemoteAllowedMacs(t *testing.T) {
	box := Box{}
	box.setBoxname("boxname")
	box.loadConfig()

	if err := box.setRemoteAllowedMACs(); err != nil {
		t.Errorf("setRemoteAllowedMACs: Got an error [%v].", err)
	}
}
