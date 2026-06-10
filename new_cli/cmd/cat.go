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

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat <file>",
	Short: "Print the contents of a file to standard output.",
	Long: `cat reads the file given as the first argument and writes its contents,
line by line, to standard output. It is a minimal re-implementation of the
classic Unix "cat" command, intended for learning purposes.

Behaviour:
  - Exactly one positional argument is required: the path to the file.
  - The file is read with a buffered scanner; each line is printed as-is
    followed by a newline.
  - If the file cannot be opened, or a read error occurs mid-stream, the
    program reports the error and exits with a non-zero status.

Examples:
  newcli cat ./test.txt
  newcli cat /etc/hosts`,
	Run: func(cmd *cobra.Command, args []string) {
		if debug {
			log.Printf("command cat run, args: %v\n---------", args)
		}
		err := tool.Cat(args)
		if err != nil {
			log.Fatalf("execute error: %v\n", err)
		} else if debug {
			fmt.Println("---------")
			log.Println("command cat run successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(catCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// catCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
