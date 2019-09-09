/* File: mac.go */
/* Description: MAC configuration for remote boxes. */

package rconfig

// allowedMACs ... MACs allowed internet access
type allowedMACs []string

/* Return allowedMACs as a string */
func (a allowedMACs) String() string {
	macsAsStr := "{"
	for i, v := range a {
		if i == len(a)-1 {
			macsAsStr += v
		} else {
			macsAsStr += v + ","
		}
	}
	macsAsStr += "}"
	return macsAsStr
}
