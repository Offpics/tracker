/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"time"
	"tracker/utils"

	"github.com/spf13/cobra"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start [activity to track]",
	Short: "Start tracking your activity",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := ui.Init(); err != nil {
			log.Fatalf("failed to initialize termui: %v", err)
		}
		defer ui.Close()

		timer := createTimer()
		timer.startTimer()

		// activityName := strings.Join(args, "")

		uiEvents := ui.PollEvents()
		for {
			e := <-uiEvents
			switch e.ID {
			case "q", "<C-c>":
				// Save time to csv
				return
			}
		}
	},
}

type Timer struct {
	currentTime int
	ticker      *time.Ticker
}

func (timer *Timer) startTimer() {
	paragraph := widgets.NewParagraph()
	paragraph.Text = utils.SecondsToTime(timer.currentTime)
	paragraph.SetRect(0, 0, 25, 5)

	ui.Render(paragraph)

	go func() {
		for {
			<-timer.ticker.C
			timer.currentTime += 1
			paragraph.Text = utils.SecondsToTime(timer.currentTime)
			ui.Render(paragraph)
		}
	}()
}

func createTimer() Timer {
	ticker := time.NewTicker(1000 * time.Millisecond)
	return Timer{0, ticker}
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
