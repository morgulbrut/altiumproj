package cmd

import (
	"github.com/morgulbrut/colorlog"
	"github.com/spf13/cobra"
)

// sixlayerCmd represents the sixlayer command
var sixlayerCmd = &cobra.Command{
	Use:   "sixlayer",
	Short: "Initializes a new 6 layer altium project: init sixlayer <PROJECTNAME>",
	Run: func(cmd *cobra.Command, args []string) {
		colorlog.Info("Initalizing 6-layer project %s", args[0])
		InitializeProject("sixlayer", args[0])
		FixProject("Template_6L", args[0])
	},
}

func init() {
	initCmd.AddCommand(sixlayerCmd)
}
