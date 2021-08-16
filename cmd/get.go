package cmd

import (
	"fmt"
	"github.com/ronaldoafonso/rconfig/rbox"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&field, "field", "f", "all", "config field")
}

var (
	field  string
	getCmd = &cobra.Command{
		Use:   "get BOXNAME [BOXNAMES...]",
		Short: "Get configuration of remotebox(es)",
		Args:  cobra.MinimumNArgs(1),
		Run:   get,
	}
)

type getResult struct {
	boxname string
	err     error
	info    string
}

func get(cmd *cobra.Command, boxnames []string) {
	results := make(chan getResult)

	for _, boxname := range boxnames {
		go func(boxname string) {
			var err error
			var info string
			b := rbox.NewRBox(boxname)

			switch field {
			case "all":
				err = b.GetConfig()
			case "ssid":
				SSIDs := []string{}
				SSIDs, err = b.GetSSIDs()
				info = strings.Join(SSIDs, " ")
			case "macs":
				MACs := []string{}
				MACs, err = b.GetMACs()
				info = strings.Join(MACs, " ")
			default:
				log.Fatal("Unsuported field.")
			}
			results <- getResult{boxname, err, info}
		}(boxname)
	}

	for i := 0; i < len(boxnames); i++ {
		result := <-results
		if result.err != nil {
			fmt.Printf("%v box: %v.\n", result.boxname, result.err)
		} else {
			if result.info == "" {
				fmt.Printf("%v box: OK.\n", result.boxname)
			} else {
				fmt.Printf("%v: [%v].\n", result.boxname, result.info)
			}
		}
	}
}
