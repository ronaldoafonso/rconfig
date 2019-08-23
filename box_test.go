/* File: box_test.go */
/* Description: Test operations executed on remote boxes. */

package rconfig

import (
	"reflect"
	"testing"
)

var (
	boxname          = "boxname"
	otherIgualConfig = boxConfig{
		ssid: "ssid",
		allowedMacs: []string{
			"11:11:11:11:11:11",
			"22:22:22:22:22:22",
		},
	}
)

func TestInitBox(t *testing.T) {
	box := initBox()

	if reflect.TypeOf(box) != reflect.TypeOf(Box{}) {
		t.Errorf("initBox didn't return a Box type.")
	}

	if box.boxname != "" {
		t.Errorf("initBox: boxname must be an empty string.")
	}

	if reflect.TypeOf(box.config) != reflect.TypeOf(boxConfig{}) {
		t.Errorf("initBox: config must be a boxConfig type.")
	}
}

func TestLoadBoxConfigOk(t *testing.T) {
	box := initBox()
	err := box.loadConfig()

	if err != nil {
		t.Errorf("loadConfig: Got an error [%v].", err)
	}

	if box.boxname != boxname {
		t.Errorf("loadConfig: Didn't set boxname. Want: %v, got: %v.",
			boxname, box.boxname)
	}

	if !box.config.isIgual(otherIgualConfig) {
		t.Errorf("loadConfig: Didn't set config. Want: %v, got: %v.",
			otherIgualConfig, box.config)
	}
}

func TestGetRemoteBoxSsid(t *testing.T) {
	box := initBox()
	box.loadConfig()
	ssid, err := box.getRemoteSsid()

	if err != nil {
		t.Errorf("getRemoteSsid: Got an error [%v].", err)
	}

	if ssid != "ssid" {
		t.Errorf("getRemoteSsid: Want: %v, got: %v.", "ssid", ssid)
	}
}

func TestGetRemoteAllowedMacs(t *testing.T) {
	box := initBox()
	box.loadConfig()
	allowedMacsWanted := []string{
		"11:11:11:11:11:11",
		"22:22:22:22:22:22",
	}

	allowedMacs, err := box.getRemoteAllowedMacs()

	if err != nil {
		t.Errorf("getRemoteAllowedMacs: Got an error [%v].", err)
	}

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
