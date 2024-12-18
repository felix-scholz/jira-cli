package jira

import (
	"github.com/ankitpokhrel/jira-cli/internal/cmd/root"
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	return root.NewCmdRoot()
}
