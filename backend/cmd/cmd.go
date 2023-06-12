package cmd

import (
	"backend/cmd/web"
	"github.com/spf13/cobra"
)

var RootCmd = cobra.Command{}

func init() {
	RootCmd.AddCommand(web.Command)
}
