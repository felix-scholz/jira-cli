package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/ankitpokhrel/jira-cli/internal/cmd/root"
)

func RootCmd() *cobra.Command {
	return root.NewCmdRoot()
}

func main() {
	rootCmd := root.NewCmdRoot()
	if _, err := rootCmd.ExecuteC(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
