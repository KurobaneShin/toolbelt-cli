/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// epochCmd represents the epoch command
var epochCmd = &cobra.Command{
	Use:   "epoch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if epoch, err := cmd.Flags().GetString("convert"); err == nil {
			fmt.Println(epochToHuman(epoch))
		}
	},
}

func init() {
	rootCmd.AddCommand(epochCmd)

	epochCmd.PersistentFlags().String("convert", "", "epoch to convert")
}

func epochToHuman(epoch string) string {
	date, err := strconv.ParseInt(epoch, 10, 64)
	if err != nil {
		return "Invalid epoch"
	}
	return time.Unix(date, 0).Format("2006-01-02 15:04:05")
}
