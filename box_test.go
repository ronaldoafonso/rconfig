/* File: box_test.go */
/* Description: Test operations executed on remote boxes. */

package rconfig

import (
	"reflect"
	"testing"
)

var (
	boxname = "boxname"
)

func TestInitBox(t *testing.T) {
	box := initBox(boxname)

	if reflect.TypeOf(box) != reflect.TypeOf(remoteBox{}) {
		t.Errorf("initBox didn't return a remoteBox type.")
	}

	if box.boxname != boxname {
		t.Errorf("initBox didn't set boxname. Want: %v, got: %v.",
			boxname, box.boxname)
	}

	if !box.config.isIgual(boxConfig{}) {
		t.Errorf("initBox didn't set an empty boxConfig struct.")
	}
}

func TestGetRemoteBoxSsid(t *testing.T) {
	box := initBox(boxname)
	ssid := box.getRemoteSsid()

	if ssid != "ssid" {
		t.Errorf("getRemoteSsid error. Want: %v, got: %v.", "ssid", ssid)
	}
}

func TestGetRemoteAllowedMacs(t *testing.T) {
	box := initBox(boxname)
	allowedMacsWanted := []string{
		"11:11:11:11:11:11",
		"22:22:22:22:22:22",
	}

	allowedMacs := box.getRemoteAllowedMacs()

	if len(allowedMacs) == 0 {
		t.Errorf("getRemoteAllowedMacs returned an empty list of MACs.")
	}

	if len(allowedMacs) != len(allowedMacsWanted) {
		t.Errorf("getRemoteAllowedMacs returned a wrong list of MACs. "+
			"Want: %v, got: %v", allowedMacsWanted, allowedMacs)
	}

	for i, _ := range allowedMacs {
		if allowedMacs[i] != allowedMacsWanted[i] {
			t.Errorf("getRemoteAllowedMacs error. Want: %v, got: %v.",
				allowedMacsWanted[i],
				allowedMacs[i])
		}
	}
}
