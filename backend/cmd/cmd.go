package cmd

import (
	"github.com/spf13/cobra"

	"backend/cmd/web"
)

var RootCmd = cobra.Command{}

func init() {
	RootCmd.AddCommand(web.Command)
}
