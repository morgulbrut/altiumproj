package cmd

import (
	"io/ioutil"
	"os"

	"github.com/gobuffalo/packr/v2"
	"github.com/morgulbrut/altiumproj/utils"
	"github.com/morgulbrut/colorlog"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init <TEMPLATENAME> <NAME>",
	Short: "Initalize a new project",
	Long:  "Initalize a new project in the directory its called.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			tmp := args[0]
			name := args[1]
			InitializeProject(tmp, name)
			os.Chdir(name)
			FixProject(tmp, name)
		} else {
			colorlog.Fatal("Wrong amount of arguments")
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func InitializeProject(dst string, project string) (err error) {
	err = writeTemplateZip(dst, project)
	if err != nil {
		colorlog.Fatal(err.Error())
		os.Exit(1)
	}
	_, err = utils.Unzip(project+".zip", project)
	if err != nil {
		colorlog.Fatal(err.Error())
		os.Exit(1)
	}
	os.Remove(project + ".zip")

	rename := []string{"BomDoc", "PcbDoc", "PrjPCB", "SchDoc"}
	err = utils.RenameFiles(project, project, rename)
	if err != nil {
		colorlog.Fatal(err.Error())
		os.Exit(1)
	}

	return
}

func writeTemplateZip(tmpl string, project string) (err error) {
	colorlog.Debug("Unzipping template %q to %s", tmpl, project)
	pt := packr.New("projects", "../templates")
	zip, err := pt.Find(tmpl + ".zip")
	if err != nil {
		colorlog.Fatal(err.Error())
	}
	err = ioutil.WriteFile(project+".zip", zip, 0644)
	if err != nil {
		colorlog.Fatal(err.Error())
	}
	return
}
