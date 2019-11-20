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
	Use:   "rename <OLDNAME> <NEWNAME>",
	Short: "Renames a project and it's files",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			oldname := args[0]
			newname := args[1]
			RenameFiles(oldname, newname)
			FixProject(oldname, newname)
		} else {
			colorlog.Fatal("Wrong amount of arguments")
			os.Exit(0)
		}
	},
}

func RenameFiles(oldname, newname string) {
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

func FixProject(oldname, newname string) {
	in, err := ioutil.ReadFile(newname + ".PrjPCB")
	if err != nil {
		log.Fatalln(err)
	}
	input := string(in)
	out := strings.ReplaceAll(input, "DocumentPath="+oldname, "DocumentPath="+newname)
	out = strings.ReplaceAll(out, "Project Outputs for "+oldname, "Project Outputs for "+newname)
	err = ioutil.WriteFile(newname+".PrjPCB", []byte(out), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	rootCmd.AddCommand(renameCmd)
}
