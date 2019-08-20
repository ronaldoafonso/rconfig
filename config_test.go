/* File: config_test.go */
/* Description: Test configuration for remote boxes. */

package rconfig

import (
	"reflect"
	"testing"
)

var (
	ssid        = "ssid"
	allowedMacs = []string{
		"11:11:11:11:11:11",
		"22:22:22:22:22:22",
	}
)

func TestInitBoxConfig(t *testing.T) {
	config := initBoxConfig(ssid, allowedMacs)

	if reflect.TypeOf(config) != reflect.TypeOf(boxConfig{}) {
		t.Errorf("initBoxConfig didn't return a boxConfig type.")
	}

	if config.ssid != ssid {
		t.Errorf("initBoxConfig didn't set ssid. Want: %v, got: %v.",
			ssid, config.ssid)
	}

	for i := range allowedMacs {
		if config.allowedMacs[i] != allowedMacs[i] {
			t.Error("initBoxConfig didn't set allowedMacs correctly. "+
				"Want: %v, got: %v.", allowedMacs[i], config.allowedMacs[i])
		}
	}
}

func TestIsIgual(t *testing.T) {
	config := initBoxConfig(ssid, allowedMacs)
	otherIgual := boxConfig{
		ssid:        ssid,
		allowedMacs: allowedMacs,
	}

	if !config.isIgual(otherIgual) {
		t.Errorf("isIgual returned false when compared to another igual"+
			"boxConfig struct. config: %v, other: %v.", config, otherIgual)
	}

	for i := range config.allowedMacs {
		if config.allowedMacs[i] != otherIgual.allowedMacs[i] {
			t.Errorf("isIgual returned false when compared to another igual"+
				"boxConfig struct(allowedMacs). config: %v, other %v.",
				config, otherIgual)
		}
	}

	otherDiff := []boxConfig{
		boxConfig{
			ssid:        "another ssid",
			allowedMacs: allowedMacs,
		},
		boxConfig{
			ssid: ssid,
			allowedMacs: []string{
				"33:33:33:33:33",
			},
		},
		boxConfig{
			ssid: ssid,
			allowedMacs: []string{
				"44:44:44:44:44",
				"55:55:55:55:55:55",
			},
		},
	}

	for i := range otherDiff {
		if config.isIgual(otherDiff[i]) {
			t.Errorf("isIgual returned true when compared to another different "+
				"boxConfig struct. config: %v, other: %v.", config, otherDiff[i])
		}
	}
}

func TestUpDateSsid(t *testing.T) {
	config := initBoxConfig(ssid, allowedMacs)
	otherSsid := "other ssid"
	config.updateSsid(otherSsid)

	if config.ssid != otherSsid {
		t.Errorf("updateSsid didn't update field ssid. Want: %v, got: %v.",
			otherSsid, config.ssid)
	}
}

func TestUpDateAllowedMacs(t *testing.T) {
	config := initBoxConfig(ssid, allowedMacs)
	otherAllowedMacs := append(allowedMacs, "33:33:33:33:33:33")
	config.updateAllowedMacs(otherAllowedMacs)

	for i := range config.allowedMacs {
		if config.allowedMacs[i] != otherAllowedMacs[i] {
			t.Errorf("updateAllowedMacs didn't update field allowedMacs. "+
				"Want: %v, got: %v.", otherAllowedMacs, config.allowedMacs)
		}
	}
}
