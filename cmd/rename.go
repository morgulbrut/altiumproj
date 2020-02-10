package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/morgulbrut/altiumproj/utils"
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
			rename := []string{"BomDoc", "PcbDoc", "PrjPCB", "SchDoc"}
			utils.RenameFiles(oldname, newname, rename)
			os.Rename(oldname, newname)
			os.Chdir(newname)
			FixProject(oldname, newname)
		} else {
			colorlog.Fatal("Wrong amount of arguments")
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
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
