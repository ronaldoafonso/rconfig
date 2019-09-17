/* File: config.go */
/* Description: Configuration for remote boxes. */

package rconfig

import (
	"github.com/ronaldoafonso/rconfig/rmac"
)

// Config ... Configuration parameters for an OpenWrt box
type Config struct {
	SSID      string
	LeaseTime string
	rmac.AllowedMACs
}

/* Return false if other is different from config. */
func (config Config) IsIgual(other Config) bool {
	if config.SSID != other.SSID {
		return false
	}

	if config.LeaseTime != other.LeaseTime {
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

/* Update field "LeaseTime" of Config structure */
func (config *Config) updateLeaseTime(leaseTime string) {
	config.LeaseTime = leaseTime
}

/* Update field "AllowedMACs" of Config structure */
func (config *Config) updateAllowedMACs(AllowedMACs []string) {
	config.AllowedMACs = AllowedMACs
}
