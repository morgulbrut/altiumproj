/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/morgulbrut/colorlog"
	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename <oldname> <newname>",
	Short: "Renames a project and it's files",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			oldname := args[0]
			newname := args[1]
			renameFiles(oldname, newname)
		} else {
			colorlog.Fatal("Wrong amount of arguments")
			os.Exit(0)
		}
	},
}

func renameFiles(oldname, newname string) {
	colorlog.Info("Renaming project %s to %s", oldname, newname)
	files, err := filepath.Glob("*" + oldname + "*")
	if err != nil {
		colorlog.Fatal(err.Error())
	}
	pwd, err := os.Getwd()
	if err != nil {
		colorlog.Fatal(err.Error())
	}

	for _, f := range files {
		colorlog.Debug("renaming %s", f)
		oldpath := filepath.Join(pwd, f)
		filetype := strings.Split(f, ".")[1]
		newpath := filepath.Join(pwd, newname+"."+filetype)
		os.Rename(oldpath, newpath)
	}
}

func fixProject(oldname, newname string) {
	in, err := ioutil.ReadFile(newname + ".PrjPCB")
	if err != nil {
		log.Fatalln(err)
	}

	input := string(in)
	out := []byte(strings.ReplaceAll(input, "DocumentPath="+oldname, "DocumentPath="+newname))

	err = ioutil.WriteFile(newname+".PrjPCB", out, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	rootCmd.AddCommand(renameCmd)
}
