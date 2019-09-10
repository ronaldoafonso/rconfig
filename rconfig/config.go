/* File: config.go */
/* Description: Configuration for remote boxes. */

package rconfig

import (
	"github.com/ronaldoafonso/rconfig/rmac"
)

// Config ... Configuration parameters for an OpenWrt box
type Config struct {
	SSID string
	rmac.AllowedMACs
}

/* Return false if other is different from config. */
func (config Config) isIgual(other Config) bool {
	if config.SSID != other.SSID {
		return false
	}

	if len(config.AllowedMACs) != len(other.AllowedMACs) {
		return false
	}

	for i, value := range config.AllowedMACs {
		if value != other.AllowedMACs[i] {
			return false
		}
	}

	return true
}

/* Update field "SSID" of Config structure */
func (config *Config) updateSSID(SSID string) {
	config.SSID = SSID
}

/* Update field "AllowedMACs" of Config structure */
func (config *Config) updateAllowedMACs(AllowedMACs []string) {
	config.AllowedMACs = AllowedMACs
}
