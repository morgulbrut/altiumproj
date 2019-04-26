package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/morgulbrut/color"
)

func RenameFiles(dst string, projName string, rename []string) (err error) {
	entries, err := ioutil.ReadDir(dst)
	if err != nil {
		return
	}
	for _, entry := range entries {
		filetype := strings.Split(entry.Name(), ".")[1]
		if StringInSlice(filetype, rename) {
			oldpath := filepath.Join(dst, entry.Name())
			newpath := filepath.Join(dst, projName+"."+filetype)
			color.Yellow("Renaming %s to %s", oldpath, newpath)
			err := os.Rename(oldpath, newpath)
			if err != nil {
				color.Red(err.Error())
			}
		}
	}
	return
}

func CleanUpDir(dst string, del []string) (err error) {
	entries, err := ioutil.ReadDir(dst)
	if err != nil {
		return
	}
	for _, entry := range entries {
		filetype := strings.Split(entry.Name(), ".")[1]
		if StringInSlice(filetype, del) {
			oldpath := filepath.Join(dst, entry.Name())
			color.Yellow("Deleting %s", oldpath)
			err := os.Remove(oldpath)
			if err != nil {
				color.Red(err.Error())
			}
		}
	}
	return
}
