/* File: mac.go */
/* Description: MAC configuration for remote boxes. */

package rmac

// AllowedMACs ... MACs allowed internet access
type AllowedMACs []string

/* Return AllowedMACs as a string */
func (a AllowedMACs) String() string {
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
