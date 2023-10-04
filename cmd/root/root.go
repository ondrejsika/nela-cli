package root

import (
	"github.com/ondrejsika/nela-cli/pkg/nela"
	"github.com/ondrejsika/nela-cli/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "nela",
	Short: "nela, " + version.Version,
	Run: func(c *cobra.Command, args []string) {
		nela.PrintRandomNela()
	},
}
