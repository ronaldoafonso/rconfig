package cmd

import (
	"fmt"
	"github.com/ronaldoafonso/rconfig/rbox"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get BOXNAME [BOXNAMES...]",
	Short: "Get configuration of remotebox(es)",
	Args:  cobra.MinimumNArgs(1),
	Run:   get,
}

type Result struct {
	boxname string
	err     error
}

func get(cmd *cobra.Command, boxnames []string) {
	results := make(chan Result)

	for _, boxname := range boxnames {
		go func(boxname string) {
			b := rbox.NewRBox(boxname)
			err := b.GetConfig()
			results <- Result{boxname, err}
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
