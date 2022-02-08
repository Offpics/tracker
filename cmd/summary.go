/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
	"tracker/types"
	"tracker/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		csvFilePath := viper.GetString("csvFilePath")
		records := getRecordsFromCSV(csvFilePath)
		printSummary(records)
	},
}

// summaryToday represents the summary command
var summaryToday = &cobra.Command{
	Use:   "today",
	Short: "A brief description of your command",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		csvFilePath := viper.GetString("csvFilePath")
		records := getRecordsFromCSV(csvFilePath)
		printSummary(filterRecordsForToday(records))
	},
}

func printSummary(records []types.Record) {
	activities := make(map[string]int)

	for _, record := range records {
		if _, ok := activities[record.ActivityName]; !ok {
			activities[record.ActivityName] = 0
		}
		activities[record.ActivityName] += record.Time
	}

	for activity, time := range activities {
		fmt.Printf("You spent %v minutes on %s \n", utils.SecondsToTime(time), activity)
	}
}

func filterRecordsForToday(records []types.Record) []types.Record {
	yesterday := time.Now().AddDate(0, 0, -1)
	tomorrow := time.Now().AddDate(0, 0, 1)

	filtered_records := make([]types.Record, 0)

	for _, record := range records {
		if record.Date.After(yesterday) && record.Date.Before(tomorrow) {
			filtered_records = append(filtered_records, record)
		}
	}

	return filtered_records
}

func getRecordsFromCSV(filePath string) []types.Record {
	file, err := os.Open(filePath)
	utils.Check(err)
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	utils.Check(err)

	parsedRecords := make([]types.Record, 0, len(rows))

	for _, row := range rows {
		parsedRecords = append(parsedRecords, types.RecordFromStringArray(row))
	}

	file.Close()

	return parsedRecords
}

func init() {
	rootCmd.AddCommand(summaryCmd)
	summaryCmd.AddCommand(summaryToday)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// summaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// summaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
