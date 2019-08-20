/* File: config.go */
/* Description: Configuration for remote boxes. */

package rconfig

type boxConfig struct {
	ssid string
}

/* Initialize and return a boxConfig structure */
func initBoxConfig(ssid string) boxConfig {
	return boxConfig{
		ssid: ssid,
	}
}

/* Return false if other is different from config. */
func (config boxConfig) isIgual(other boxConfig) bool {
	if config.ssid != other.ssid {
		return false
	}
	return true
}

/* Update field "ssid" of boxConfig structure */
func (config *boxConfig) upDateSsid(ssid string) {
	config.ssid = ssid
}
