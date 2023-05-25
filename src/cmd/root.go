/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/aleroxac/cernner/task"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cernner",
	Short: "A platform to grok yourself.",
	Long:  "A platform to grok yourself. A place where is possivel organize, to see and to show what you know and who are you.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(task.TaskCmd)
}
