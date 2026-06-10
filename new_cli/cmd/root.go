/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// go build -ldflags "-X 'newcli/cmd.version=0.1.0'" .
var version = "1.1.1"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "newcli",
	Version: version,
	Short:   "A small learning-purpose CLI that re-implements basic Unix tools (cat, grep) in Go.",
	Long: `newcli is a minimal command-line toolkit written in Go with the Cobra framework.
It is built as a study project to practice CLI design, file I/O and subcommand
organization by re-implementing a couple of classic Unix utilities.

Currently supported subcommands:
  cat   - print the contents of a file to standard output, line by line.
  grep  - print every line in a file that contains a given substring.

Usage examples:
  newcli cat ./test.txt
  newcli grep hello ./test.txt

Run "newcli <command> --help" to see the detailed help for any subcommand.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var debug bool

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.newcli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable logging")
}
