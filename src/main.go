package main

import (
	"fmt"
	"os"

	"github.com/aleroxac/cernner/skill/skill"
	"github.com/spf13/cobra"
)

func create_root_cmd(use string, short string) *cobra.Command {
	return &cobra.Command{
		Use:   "cern [kind] [options]",
		Short: "Cernner command line tool",
	}
}

func add_cmd(c *cobra.Command) {
	c.AddCommand(skill.SkillCmd)
}

func run_cmd(c *cobra.Command) {
	if err := c.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	rootCmd := create_root_cmd("cern [kind] [options]", "Cernner command line tool")
	add_cmd(rootCmd)
	run_cmd(rootCmd)
}
