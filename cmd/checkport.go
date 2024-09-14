/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// checkportCmd represents the checkport command
var checkportCmd = &cobra.Command{
	Use:   "checkport",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			fmt.Println(err)
		}
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			fmt.Println(err)
		}
		ports := strings.Split(port, ",")
		CheckPort(host, ports)
	},
}

func init() {
	rootCmd.AddCommand(checkportCmd)

	checkportCmd.PersistentFlags().String("host", "", "Host to check")
	checkportCmd.PersistentFlags().String("port", "", "Port to check")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func CheckPort(host string, ports []string) {
	for _, port := range ports {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		if err != nil {
			fmt.Printf("Port %s is closed\n", port)
			continue
		}
		defer conn.Close()
		fmt.Printf("Port %s is open\n", port)
	}
}
