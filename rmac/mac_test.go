/* File: mac_test.go */
/* Description: Test MAC operation */

package rmac

import (
	"fmt"
	"testing"
)

func TestMACAllowedMACsAsString(t *testing.T) {
	type MACs struct {
		AllowedMACs
		asString string
	}

	testMACs := []MACs{
		MACs{
			AllowedMACs{},
			"{}",
		},
		MACs{
			AllowedMACs{
				"11:11:11:11:11:11",
			},
			"{11:11:11:11:11:11}",
		},
		MACs{
			AllowedMACs{
				"11:11:11:11:11:11",
				"22:22:22:22:22:22",
			},
			"{11:11:11:11:11:11,22:22:22:22:22:22}",
		},
		MACs{
			AllowedMACs{
				"11:11:11:11:11:11",
				"22:22:22:22:22:22",
				"33:33:33:33:33:33",
			},
			"{11:11:11:11:11:11,22:22:22:22:22:22,33:33:33:33:33:33}",
		},
	}

	for _, MAC := range testMACs {
		if fmt.Sprintf("%s", MAC.AllowedMACs) != MAC.asString {
			t.Errorf("allowedMACAsString: Want: %v, got %v.", MAC.asString, fmt.Sprintf("%s", MAC.AllowedMACs))
		}
	}
}
