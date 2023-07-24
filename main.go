package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Dir struct {
	Name string `yaml:"name"`
}

type DirActions struct {
	Create []Dir `yaml:"create"`
	Delete []Dir `yaml:"delete"`
}

type File struct {
	Name    string `yaml:"name"`
	Content string `yaml:"content"`
}

type FileActions struct {
	Apply  []File `yaml:"apply"`
	Delete []File `yaml:"delete"`
}

type ScriptElement struct {
	Args   []string    `yaml:"args"`
	Output string      `yaml:"output"`
	Dirs   DirActions  `yaml:"dirs"`
	Files  FileActions `yaml:"files"`
}

type Script struct {
	Name        string          `yaml:"name"`
	CliName     string          `yaml:"cli-name"`
	Description string          `yaml:"description"`
	Script      []ScriptElement `yaml:"script"`
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Could not get working dir: %s\n", err.Error())
		os.Exit(1)
	}

	configPath, err := findConfig(wd)
	if err != nil {
		fmt.Printf("Could not locate an .appy file, failed with: %s\n", err.Error())
		os.Exit(1)
	}

	config, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Could not read .appy file, failed with: %s\n", err.Error())
		os.Exit(1)
	}

	configValues := strings.Split(string(config), "\n")

	scriptPath := configValues[0]

	scriptData, err := os.ReadFile(scriptPath)
	if err != nil {
		fmt.Printf("Could not read script file `%s`, failed with: %s\n", scriptPath, err.Error())
		os.Exit(1)
	}

	script := &Script{}
	err = yaml.Unmarshal(scriptData, script)
	if err != nil {
		fmt.Printf("Could not parse yaml in script file `%s`, failed with: %s\n", scriptPath, err.Error())
		os.Exit(1)
	}

	var pos int
	if len(configValues) < 2 {
		pos = 0
	} else {
		pos, err = strconv.Atoi(configValues[1])
		if err != nil {
			pos = 0
		}
	}

	//os.Args[1:]

	currentLine := script.Script[pos]

	if len(currentLine.Args) != len(os.Args)-1 {
		fmt.Printf("\a")
		if os.Getenv("APPY_DEBUG") == "true" {
			fmt.Printf("wrong number of arguments, expected %d, got %d", len(currentLine.Args), len(os.Args)-1)
			os.Exit(1)
		} else {
			os.Exit(0)
		}

	}

	for i, expected := range currentLine.Args {
		got := os.Args[i+1]
		if got != expected {
			fmt.Printf("\a")
			if os.Getenv("APPY_DEBUG") == "true" {
				fmt.Printf("failed to match argument in position %d, expected `%s`, but got `%s`", i+1, expected, got)
				os.Exit(1)
			} else {
				os.Exit(0)
			}
		}
	}
}

func findConfig(dir string) (string, error) {
	config := filepath.Join(dir, ".appy")

	_, err := os.Stat(config)
	if err != nil {
		if os.IsNotExist(err) {
			splitDir := strings.Split(dir, "/")
			if len(splitDir) < 2 {
				return "", fmt.Errorf("could not find .appy in the current dir or any ascendants")
			}
			return findConfig(strings.Join(splitDir[:len(splitDir)-1], "/"))
		} else {
			return "", err
		}
	}

	return config, nil
}
