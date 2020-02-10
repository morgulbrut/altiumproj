/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/morgulbrut/colorlog"
	"github.com/spf13/cobra"
)

// outjobsCmd represents the outjobs command
var outjobsCmd = &cobra.Command{
	Use:   "outjobs",
	Short: "updates the outjob files to a newer version",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			tmp := args[0]
			removeOutJobs()
			InitializeProject(tmp, "temp")
			moveOutJobs()
			os.RemoveAll("temp")
		} else {
			colorlog.Fatal("Wrong amount of arguments")
			os.Exit(0)
		}
	},
}

func moveOutJobs() {
	files, err := filepath.Glob("temp/*.OutJob")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		dest := strings.Replace(f, "temp", ".", 1)
		if err := os.Rename(f, dest); err != nil {
			panic(err)
		}
	}
}

func removeOutJobs() {
	files, err := filepath.Glob("./*.OutJob")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}
}

func init() {
	updateCmd.AddCommand(outjobsCmd)
}
