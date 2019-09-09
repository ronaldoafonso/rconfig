/* File: mac_test.go */
/* Description: Test MAC operation */

package rconfig

import (
	"fmt"
	"testing"
)

func TestMACAllowedMACsAsString(t *testing.T) {
	type MACs struct {
		allowedMACs
		asString string
	}

	testMACs := []MACs{
		MACs{
			allowedMACs{},
			"{}",
		},
		MACs{
			allowedMACs{
				"11:11:11:11:11:11",
			},
			"{11:11:11:11:11:11}",
		},
		MACs{
			allowedMACs{
				"11:11:11:11:11:11",
				"22:22:22:22:22:22",
			},
			"{11:11:11:11:11:11,22:22:22:22:22:22}",
		},
		MACs{
			allowedMACs{
				"11:11:11:11:11:11",
				"22:22:22:22:22:22",
				"33:33:33:33:33:33",
			},
			"{11:11:11:11:11:11,22:22:22:22:22:22,33:33:33:33:33:33}",
		},
	}

	for _, MAC := range testMACs {
		if fmt.Sprintf("%s", MAC.allowedMACs) != MAC.asString {
			t.Errorf("allowedMACAsString: Want: %v, got %v.", MAC.asString, fmt.Sprintf("%s", MAC.allowedMACs))
		}
	}
}
