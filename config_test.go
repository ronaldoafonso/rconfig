/* File: config_test.go */
/* Description: Test configuration for remote boxes. */

package rconfig

import (
	"reflect"
	"testing"
)

var (
	ssid = "ssid"
)

func TestInitBoxConfig(t *testing.T) {
	config := initBoxConfig(ssid)

	if reflect.TypeOf(config) != reflect.TypeOf(boxConfig{}) {
		t.Errorf("initBoxConfig didn't return a boxConfig type.")
	}

	if config.ssid != ssid {
		t.Errorf("initBoxConfig didn't set ssid. Want: %v, got: %v.",
			ssid, config.ssid)
	}
}

func TestIsIgual(t *testing.T) {
	config := initBoxConfig(ssid)
	otherIgual := boxConfig{
		ssid: ssid,
	}

	if !config.isIgual(otherIgual) {
		t.Errorf("isIgual returned false when compared to another igual"+
			"boxConfig struct. config: %v, other: %v.", config, otherIgual)
	}

	otherDifferent := boxConfig{
		ssid: "another ssid",
	}

	if config.isIgual(otherDifferent) {
		t.Errorf("isIgual returned true when compared to another different"+
			"boxConfig struct. config: %v, other: %v.", config, otherDifferent)
	}
}

func TestUpDateSsid(t *testing.T) {
	config := initBoxConfig(ssid)
	otherSsid := "other ssid"
	config.upDateSsid(otherSsid)

	if config.ssid != otherSsid {
		t.Errorf("upDateSsid didn't update field ssid. Want: %v, got: %v.",
			otherSsid, config.ssid)
	}
}
