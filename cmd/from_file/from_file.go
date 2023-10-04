package from_file

import (
	"github.com/ondrejsika/nela-cli/cmd/root"
	pkg_from_file "github.com/ondrejsika/nela-cli/pkg/from_file"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "from-file <path>",
	Short:   "Print image from file into terminal",
	Aliases: []string{"ff"},
	Args:    cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		pkg_from_file.PrintFromFile(args[0])
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
