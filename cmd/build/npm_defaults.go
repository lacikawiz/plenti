package build

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"plenti/generated"
	"time"
)

// NpmDefaults creates the node_modules folder with core defaults if it doesn't already exist.
func NpmDefaults() {

	defer Benchmark(time.Now(), "Setting up core NPM packages")

	Log("\nChecking if 'node_modules' directory exists.")

	if _, err := os.Stat("node_modules"); os.IsNotExist(err) {
		for file, content := range generated.Defaults_node_modules {
			// Make file relative to where CLI is executed
			file = "./node_modules/" + file
			// Create the directories needed for the current file
			os.MkdirAll(filepath.Dir(file), os.ModePerm)
			// Create the current default file
			err := ioutil.WriteFile(file, content, os.ModePerm)
			if err != nil {
				fmt.Printf("Unable to write npm dependency file: %v", err)
			}
		}
	}

}
