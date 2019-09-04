/* File: config_test.go */
/* Description: Test configuration for remote boxes. */

package rconfig

import (
	"testing"
)

var ()

func TestConfigIsIgual(t *testing.T) {
	SSID := "ssid"
	allowedMACs := []string{
		"11:11:11:11:11:11",
		"22:22:22:22:22:22",
	}
	config := Config{
		SSID:        SSID,
		allowedMACs: allowedMACs,
	}

	otherIgual := Config{
		SSID:        SSID,
		allowedMACs: allowedMACs,
	}

	if !config.isIgual(otherIgual) {
		t.Errorf("ConfigIsIgual returned false when compared to another igual"+
			"Config struct. config: %v, other: %v.", config, otherIgual)
	}

	for i := range config.allowedMACs {
		if config.allowedMACs[i] != otherIgual.allowedMACs[i] {
			t.Errorf("ConfigisIgual returned false when compared to another igual"+
				"Config struct(allowedMACs). config: %v, other %v.",
				config, otherIgual)
		}
	}

	otherDiff := []Config{
		Config{
			SSID:        "another ssid",
			allowedMACs: allowedMACs,
		},
		Config{
			SSID: SSID,
			allowedMACs: []string{
				"33:33:33:33:33",
			},
		},
		Config{
			SSID: SSID,
			allowedMACs: []string{
				"44:44:44:44:44",
				"55:55:55:55:55:55",
			},
		},
	}

	for i := range otherDiff {
		if config.isIgual(otherDiff[i]) {
			t.Errorf("ConfigIsIgual returned true when compared to another different "+
				"Config struct. config: %v, other: %v.", config, otherDiff[i])
		}
	}
}

func TestConfigUpdateSSID(t *testing.T) {
	SSID := "ssid"
	config := Config{}
	config.updateSSID(SSID)

	if config.SSID != SSID {
		t.Errorf("updateSSID didn't update field SSID. Want: %v, got: %v.",
			SSID,
			config.SSID)
	}
}

func TestUpdateAllowedMacs(t *testing.T) {
	allowedMACs := []string{
		"11:11:11:11:11:11",
		"22:22:22:22:22:22",
	}
	config := Config{}
	config.updateAllowedMACs(allowedMACs)

	for i := range config.allowedMACs {
		if config.allowedMACs[i] != allowedMACs[i] {
			t.Errorf("ConfigUpdateAllowedMACs didn't update field allowedMACs. "+
				"Want: %v, got: %v.", allowedMACs, config.allowedMACs)
		}
	}
}
