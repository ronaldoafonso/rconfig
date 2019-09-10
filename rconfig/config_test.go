/* File: config_test.go */
/* Description: Test configuration for remote boxes. */

package rconfig

import (
	"github.com/ronaldoafonso/rconfig/rmac"
	"testing"
)

func TestConfigIsIgual(t *testing.T) {
	SSID := "ssid"
	allowedMACs := rmac.AllowedMACs{
		"11:11:11:11:11:11",
		"22:22:22:22:22:22",
	}
	config := Config{
		SSID:        SSID,
		AllowedMACs: allowedMACs,
	}

	otherIgual := Config{
		SSID:        SSID,
		AllowedMACs: allowedMACs,
	}

	if !config.isIgual(otherIgual) {
		t.Errorf("ConfigIsIgual returned false when compared to another igual"+
			"Config struct. config: %v, other: %v.", config, otherIgual)
	}

	for i := range config.AllowedMACs {
		if config.AllowedMACs[i] != otherIgual.AllowedMACs[i] {
			t.Errorf("ConfigisIgual returned false when compared to another igual"+
				"Config struct(AllowedMACs). config: %v, other %v.",
				config, otherIgual)
		}
	}

	otherDiff := []Config{
		Config{
			SSID:        "another ssid",
			AllowedMACs: allowedMACs,
		},
		Config{
			SSID: SSID,
			AllowedMACs: rmac.AllowedMACs{
				"33:33:33:33:33",
			},
		},
		Config{
			SSID: SSID,
			AllowedMACs: rmac.AllowedMACs{
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

func TestConfigUpdateAllowedMacs(t *testing.T) {
	allowedMACs := rmac.AllowedMACs{
		"11:11:11:11:11:11",
		"22:22:22:22:22:22",
	}
	config := Config{}
	config.updateAllowedMACs(allowedMACs)

	for i := range config.AllowedMACs {
		if config.AllowedMACs[i] != allowedMACs[i] {
			t.Errorf("ConfigUpdateAllowedMACs didn't update field AllowedMACs. "+
				"Want: %v, got: %v.", allowedMACs, config.AllowedMACs)
		}
	}
}
