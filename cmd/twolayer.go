package cmd

import (
	"github.com/morgulbrut/colorlog"
	"github.com/spf13/cobra"
)

// twolayerCmd represents the twolayer command
var twolayerCmd = &cobra.Command{
	Use:   "twolayer",
	Short: "Initializes a new 2 layer altium project: init twolayer <PROJECTNAME>",
	Run: func(cmd *cobra.Command, args []string) {
		colorlog.Info("Initalizing 2-layer project %s", args[0])
		InitializeProject("twolayer", args[0])
		FixProject("Template_2L", args[0])
	},
}

func init() {
	initCmd.AddCommand(twolayerCmd)
}
