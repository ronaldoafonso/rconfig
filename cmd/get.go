package cmd

import (
	"github.com/ronaldoafonso/rconfig/rbox"
	"github.com/spf13/cobra"
	"sync"
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

func get(cmd *cobra.Command, boxnames []string) {
	wg := sync.WaitGroup{}
	wg.Add(len(boxnames))

	for _, boxname := range boxnames {
		go func(boxname string) {
			defer wg.Done()
			b := rbox.NewRBox(boxname)
			b.GetConfig()
		}(boxname)
	}

	wg.Wait()
}
