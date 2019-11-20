package cmd

import (
	"os"

	"github.com/morgulbrut/colorlog"
	"github.com/spf13/cobra"
)

// fourlayerCmd represents the fourlayer command
var fourlayerCmd = &cobra.Command{
	Use:   "fourlayer",
	Short: "Initializes a new 4 layer altium project: init fourlayer <PROJECTNAME>",
	Run: func(cmd *cobra.Command, args []string) {
		colorlog.Info("Initalizing 4-layer project %s", args[0])
		InitializeProject("fourlayer", args[0])
		os.Chdir(args[0])
		FixProject("Template_4L", args[0])
	},
}

func init() {
	initCmd.AddCommand(fourlayerCmd)
}
