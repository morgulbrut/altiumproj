// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"

	"github.com/morgulbrut/altiumproj/utils"
	"github.com/spf13/cobra"
)

// twolayerCmd represents the twolayer command
var twolayerCmd = &cobra.Command{
	Use:   "twolayer",
	Short: "Initializes a new 2 layer altium project: init twolayer <PROJECTNAME>",
	Run: func(cmd *cobra.Command, args []string) {
		err := utils.CopyDir(`D:\Tillo\Documents\hw_altium\trunk\Templates\PCB_Project_2Layer`, args[0])
		if err != nil {
			log.Fatal(err)
		}
		err = RenameFiles(args[0], args[0])
		if err != nil {
			log.Fatal(err)
		}
		err = CleanUpDir(args[0])
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	initCmd.AddCommand(twolayerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// twolayerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// twolayerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
