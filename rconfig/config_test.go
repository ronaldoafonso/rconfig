/* File: config_test.go */
/* Description: Test configuration for remote boxes. */

package rconfig

import (
	"github.com/ronaldoafonso/rconfig/rmac"
	"testing"
)

func TestConfigIsIgual(t *testing.T) {
	SSID := "ssid"
	leaseTime := "30m"
	allowedMACs := rmac.AllowedMACs{
		"11:11:11:11:11:11",
		"22:22:22:22:22:22",
	}

	config := Config{
		SSID:        SSID,
		LeaseTime:   leaseTime,
		AllowedMACs: allowedMACs,
	}

	otherIgual := Config{
		SSID:        SSID,
		LeaseTime:   leaseTime,
		AllowedMACs: allowedMACs,
	}

	if !config.IsIgual(otherIgual) {
		t.Errorf("ConfigIsIgual returned false when compared to another igual"+
			"Config struct. config: %v, other: %v.", config, otherIgual)
	}

	otherDiff := []Config{
		Config{
			SSID:        "another ssid",
			LeaseTime:   leaseTime,
			AllowedMACs: allowedMACs,
		},
		Config{
			SSID:        SSID,
			LeaseTime:   "60m",
			AllowedMACs: allowedMACs,
		},
		Config{
			SSID:      SSID,
			LeaseTime: leaseTime,
			AllowedMACs: rmac.AllowedMACs{
				"33:33:33:33:33",
			},
		},
		Config{
			SSID:      SSID,
			LeaseTime: leaseTime,
			AllowedMACs: rmac.AllowedMACs{
				"44:44:44:44:44",
				"55:55:55:55:55:55",
			},
		},
	}

	for _, other := range otherDiff {
		if config.IsIgual(other) {
			t.Errorf("ConfigIsIgual returned true when compared to another different "+
				"Config struct. config: %v, other: %v.", config, other)
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

func TestConfigUpdateLeaseTime(t *testing.T) {
	leaseTime := "30m"
	config := Config{}
	config.updateLeaseTime(leaseTime)

	if config.LeaseTime != leaseTime {
		t.Errorf("updateLeaseTime didn't update field LeaseTime. Want: %v, got: %v.",
			leaseTime,
			config.LeaseTime)
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
