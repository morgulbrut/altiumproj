package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/gobuffalo/packr/v2"
	"github.com/morgulbrut/altiumproj/utils"
	"github.com/morgulbrut/colorlog"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init <TEMPLATENAME> <NAME>",
	Short: "Initalize a new project",
	Long:  "Initalize a new project",
	Run: func(cmd *cobra.Command, args []string) {
		tmp := args[0]
		name := args[1]
		InitializeProject(tmp, name)
		os.Chdir(name)
		FixProject(tmp, name)
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

	// writeProjectFile(project)
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

func writeProjectFile(project string) (err error) {
	colorlog.Debug("Writing %s.PrjPCB", project)
	tpl, err := template.ParseFiles(filepath.Join(project, "Template.PrjPCBTmpl"))
	if err != nil {
		colorlog.Fatal(err.Error())
	}
	projFile, err := os.Create(filepath.Join(project, project+".PrjPCB"))
	if err != nil {
		colorlog.Fatal(err.Error())
	}
	err = tpl.Execute(projFile, project)
	if err != nil {
		colorlog.Fatal(err.Error())
	}
	os.Remove(filepath.Join(project, "Template.PrjPCBTmpl"))
	return
}
