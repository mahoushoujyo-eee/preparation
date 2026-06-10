/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"newcli/tool"

	"github.com/spf13/cobra"
)

// grepCmd represents the grep command
var grepCmd = &cobra.Command{
	Use:   "grep",
	Short: "get content which ",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := tool.Grep(args)
		if err != nil {
			log.Fatalf("execute error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(grepCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grepCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grepCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
