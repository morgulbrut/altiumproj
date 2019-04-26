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
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gobuffalo/packr/v2"
	"github.com/morgulbrut/altiumproj/utils"
	"github.com/morgulbrut/color"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	/* Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},*/
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func InitializeProject(dst string, project string) (err error) {
	//TODO FUUUUUUUUUUUUUUU!!!!!!!!!!

	err = writeTemplateZip(dst, project)
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
	_, err = utils.Unzip(project+".zip", project)
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
	os.Remove(project + ".zip")

	rename := []string{"BomDoc", "PcbDoc", "PrjPCB", "SchDoc"}
	err = utils.RenameFiles(project, project, rename)
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	return
}

func writeTemplateZip(dst string, project string) (err error) {
	pt := packr.New("projects", "../templates")
	zip, err := pt.Find(dst + ".zip")
	if err != nil {
		color.Red(err.Error())
	}
	err = ioutil.WriteFile(project+".zip", zip, 0644)
	if err != nil {
		color.Red(err.Error())
	}
	return
}
