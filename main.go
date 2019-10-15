package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bitrise-io/step-yml-linter/lint"
	"github.com/bitrise-io/step-yml-linter/step"
	"gopkg.in/yaml.v3"
)

var filePath = flag.String("file", "step.yml", "Path to the file")

func main() {
	flag.Parse()

	stepYMLContent, err := ioutil.ReadFile(*filePath)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	var s step.YML
	if err := yaml.Unmarshal(stepYMLContent, &s); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	lint.SetFilePath(*filePath)

	// checks
	for _, check := range checks {
		check(s)
	}
}
