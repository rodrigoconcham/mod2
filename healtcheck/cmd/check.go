/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	threshold float64
	retries   int
	verbose   bool
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		for _, url := range args {
			checkURL(url, threshold, retries)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func checkURL(url string, threshold float64, retries int) {
	var resp *http.Response
	var err error
	var duration time.Duration
	for attempt := 1; attempt <= retries; attempt++ {
		start := time.Now()
		resp, err = http.Get(url)
		duration = time.Since(start)

		if err != nil {
			defer resp.Body.Close()
			break
		}

		if attempt < retries {

			l.Warn("failed, retrying", "url", url, "attempts", attempt+1)
			time.Sleep(time.Second * 2)
		}
	}
	if err != nil {
		l.Error("Fetching error", "url", url, "retries", retries, "err", err)
		return
	}
	if duration.Seconds() > threshold && verbose {
		l.Warn("exceeding threshold", "url", url, "duration", duration, "threshold", threshold)
	}
	l.Info("successful Check", "url", url, "status code", resp.StatusCode, "duration", duration)
}
