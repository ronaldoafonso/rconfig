/* File: config.go */
/* Description: Configuration for remote boxes. */

package rconfig

// Config ... Configuration parameters for an OpenWrt box
type Config struct {
	SSID string
	allowedMACs
}

/* Return false if other is different from config. */
func (config Config) isIgual(other Config) bool {
	if config.SSID != other.SSID {
		return false
	}

	if len(config.allowedMACs) != len(other.allowedMACs) {
		return false
	}

	for i, value := range config.allowedMACs {
		if value != other.allowedMACs[i] {
			return false
		}
	}

	return true
}

/* Update field "SSID" of Config structure */
func (config *Config) updateSSID(SSID string) {
	config.SSID = SSID
}

/* Update field "allowedMACs" of Config structure */
func (config *Config) updateAllowedMACs(allowedMACs []string) {
	config.allowedMACs = allowedMACs
}
