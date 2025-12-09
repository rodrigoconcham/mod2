/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

var (
	interval time.Duration
)

var monitorCmd = &cobra.Command{
	Use:   "monitor [urls]",
	Short: "Monitor the health of specified URL(s) over time",
	Long:  `Continuously monitors the health of specified URL(s) at the specified interval`,
	Run: func(cmd *cobra.Command, args []string) {
		monitorURLs(args)
	},
}

func init() {

	//flag.DurationVar(&interval, "interval", 2*time.Second, "Interval between healthchecks")
	monitorCmd.Flags().DurationVar(&interval, "interval", 2*time.Second, "Interval between healthchecks")
	rootCmd.AddCommand(monitorCmd)

}
func monitorURLs(urls []string) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		for _, url := range urls {
			checkURL(url, threshold, retries)
		}
	}
}
