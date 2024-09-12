/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "Base64 encode/decode",
	Long: `
	Base64 encode/decode

	Examples:
	  toolbelt base64 --enc "Hello World"
	  toolbelt base64 --dec "SGVsbG8gV29ybGQ="
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if enc, err := cmd.Flags().GetString("enc"); err == nil {
			encoded := base64.StdEncoding.EncodeToString([]byte(enc))
			fmt.Println(encoded)
		}
		if dec, err := cmd.Flags().GetString("dec"); err == nil {
			decoded, err := base64.StdEncoding.DecodeString(dec)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(decoded))
		}
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)

	base64Cmd.PersistentFlags().String("enc", "", "Encode string to base64")
	base64Cmd.PersistentFlags().String("dec", "", "Decode base64 string")
}
