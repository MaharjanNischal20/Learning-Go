/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	uppercase bool
	lowercase bool
)

// textFormatterCmd represents the textFormatter command
var textFormatterCmd = &cobra.Command{
	Use:   "textfmt [text]",
	Short: "A simple text formatter",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runTextFormatter,
}

func init() {
	rootCmd.AddCommand(textFormatterCmd)
	textFormatterCmd.Flags().BoolVarP(&uppercase, "upper", "u", false, "convert text to uppercase")
	textFormatterCmd.Flags().BoolVarP(&lowercase, "lower", "l", false, "convert text to lowercase")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// textFormatterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// textFormatterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runTextFormatter(cmd *cobra.Command, args []string) {
	var text string
	if len(args) > 0 {
		text = strings.Join(args, " ")
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		text = strings.Join(lines, "\n")
	}
	if text == "" {
		fmt.Println("No text to format")
		return
	}
	result := text
	if uppercase {
		result = strings.ToUpper(result)
	}
	if lowercase {
		result = strings.ToLower(result)
	}
	fmt.Println(result)
}
