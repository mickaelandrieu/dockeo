package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/urfave/cli.v1"
)

func main() {

	app := cli.NewApp()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Mickaël Andrieu",
			Email: "andrieu.travail@gmail.com",
		},
	}
	app.Name = "Dockeo: the Akeneo docker installer"
	app.Usage = "Add the required files to setup a docker environment for your Akeneo project"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:      "install",
			Aliases:   []string{"a"},
			Usage:     "Install the required files to setup docker environment",
			ArgsUsage: "[version]",
			Action: func(c *cli.Context) error {
				createFiles()
				return nil
			},
		},
		{
			Name:      "remove",
			Aliases:   []string{"c"},
			Usage:     "Remove the files from (install) command",
			ArgsUsage: "[version]",
			Action: func(c *cli.Context) error {
				fmt.Println("Removed Docker files for Akeneo v", c.Args().First())
				return nil
			},
		},
	}

	app.Run(os.Args)
}

func createFiles() {
	searchDir := "files"
	filepaths := []string{}
	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Can't loop into sources folder %q\n", err)
		}

		if f.IsDir() {
			if _, errDir := os.Stat(f.Name()); os.IsNotExist(errDir) {
				if errDir != nil {
					fmt.Printf("Copy of directory failed %q\n", errDir)
				}

				os.Mkdir(f.Name(), 0644)
			}
		} else {

			filepaths = append(filepaths, path)
		}
		return nil
	})

	for _, filePath := range filepaths {
		createFile(filePath)
	}
}

func createFile(filePath string) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("File can't be read %q\n", err)
	}

	copyErr := ioutil.WriteFile(filePath[6:], file, 0644)
	if copyErr != nil {
		fmt.Printf("CopyFile failed %q\n", copyErr)
	}
}
