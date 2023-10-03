package cmd

import (
	"github.com/ondrejsika/nela-cli/cmd/root"
	_ "github.com/ondrejsika/nela-cli/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
