package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davecgh/go-spew/spew"

	"gopkg.in/yaml.v3"
)

type test struct {
	V    interface{}
	L, C int
}

func (t *test) UnmarshalYAML(node *yaml.Node) error {
	if err := node.Decode(&t.V); err != nil {
		return err
	}

	t.L, t.C = node.Line, node.Column
	return nil
}

func main() {
	filePath := "step.yml"

	stepyml, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	var sm map[test]test
	if err := yaml.Unmarshal(stepyml, &sm); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	// for _, warning := range lint.Warnings {
	// 	for _, msg := range warning.Messages {
	// 		lint.Log(filePath, warning.Line, warning.Column, msg)
	// 	}
	// }

	spew.Dump(sm)
}
