package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Could not get working dir: %s", err.Error())
		os.Exit(1)
	}

	_, err = findConfig(wd)
	if err != nil {
		fmt.Printf("Could not locate a .appy file, failed with: %s", err.Error())
		os.Exit(1)
	}

	//os.Args[1:]

}

func findConfig(dir string) (string, error) {
	config := filepath.Join(dir, ".appy")

	_, err := os.Stat(config)
	if err != nil {
		if os.IsNotExist(err) {
			return config, nil
		}
		return "", err
	}

	splitDir := strings.Split(dir, "/")
	return findConfig(strings.Join(splitDir[:len(splitDir)-1], "/"))
}
