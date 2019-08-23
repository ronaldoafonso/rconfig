/* File: config.go */
/* Description: Configuration for remote boxes. */

package rconfig

type boxConfig struct {
	ssid        string
	allowedMacs []string
}

/* Initialize and return a boxConfig structure */
func initBoxConfig() boxConfig {
	return boxConfig{}
}

/* Return false if other is different from config. */
func (config boxConfig) isIgual(other boxConfig) bool {
	if config.ssid != other.ssid {
		return false
	}

	if len(config.allowedMacs) != len(other.allowedMacs) {
		return false
	}

	for i, value := range config.allowedMacs {
		if value != other.allowedMacs[i] {
			return false
		}
	}

	return true
}

/* Update field "ssid" of boxConfig structure */
func (config *boxConfig) updateSsid(ssid string) {
	config.ssid = ssid
}

/* Update field "allowedMacs" of boxConfig structure */
func (config *boxConfig) updateAllowedMacs(allowedMacs []string) {
	config.allowedMacs = allowedMacs
}
