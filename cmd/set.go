package cmd

import (
	"fmt"
	"github.com/ronaldoafonso/rconfig/rbox"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringVarP(&SSID, "ssid", "s", "z3n", "SSID for box")
}

var (
	SSID   string
	setCmd = &cobra.Command{
		Use:   "set BOXNAME [BOXNAMES...]",
		Short: "Set configuration of remotebox(es)",
		Args:  cobra.MinimumNArgs(1),
		Run:   set,
	}
)

type setResult struct {
	boxname string
	err     error
}

func set(cmd *cobra.Command, boxnames []string) {
	results := make(chan setResult)

	for _, boxname := range boxnames {
		go func(boxname string) {
			b := rbox.NewRBox(boxname)
			err := b.SetSSIDs(SSID)
			results <- setResult{boxname, err}
		}(boxname)
	}

	for i := 0; i < len(boxnames); i++ {
		result := <-results
		if result.err != nil {
			fmt.Printf("%v box: %v.\n", result.boxname, result.err)
		} else {
			fmt.Printf("%v box: OK.\n", result.boxname)
		}
	}
}
