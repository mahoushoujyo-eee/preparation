/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"newcli/tool"

	"github.com/spf13/cobra"
)

// grepCmd represents the grep command
var grepCmd = &cobra.Command{
	Use:   "grep <pattern> <file>",
	Short: "Print every line in a file that contains the given substring.",
	Long: `grep scans the given file line by line and prints only those lines that
contain the supplied pattern as a plain substring. It is a minimal
re-implementation of the classic Unix "grep" command, intended for
learning purposes.

Arguments (positional, in order):
  1. pattern - the substring to search for (no regular expressions yet).
  2. file    - the path to the file to scan.

Behaviour:
  - The match is a simple, case-sensitive substring check
    (strings.Contains), not a regular expression.
  - Matching lines are printed in the order they appear in the file.
  - If fewer than two arguments are provided, or the file cannot be
    opened, the program reports the error and exits with a non-zero
    status.

Examples:
  newcli grep hello ./test.txt
  newcli grep TODO  ./main.go`,
	Run: func(cmd *cobra.Command, args []string) {
		if debug {
			log.Printf("command grep run, args: %v\n---------", args)
		}
		err := tool.Grep(args)
		if err != nil {
			log.Fatalf("execute error: %v", err)
		} else if debug {
			fmt.Println("---------")
			log.Println("command grep run successfully")
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
