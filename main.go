package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"bitbucket.org/hameesys/sidegen/selenium"
	"gopkg.in/yaml.v2"
	yml "gopkg.in/yaml.v2"
)

func main() {
	var uploadSideFile selenium.SideFile
	raw, err := ioutil.ReadFile("./sides/uploadReserve.side")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &uploadSideFile)

	components := selenium.NewComponents()
	componentRaw, err := ioutil.ReadFile("./component.yml")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	yml.Unmarshal(componentRaw, &components)
	for testKey, test := range uploadSideFile.Tests {
		for commandKey, command := range test.Commands {
			xpathKey := command.GetTargetXpathKey(components.Xpaths)
			if xpathKey != "" {
				uploadSideFile.Tests[testKey].Commands[commandKey].Target = strings.Replace(command.Target, "${"+xpathKey+"}", components.Xpaths[xpathKey], -1)
			}
			if _, ok := components.Xpaths[command.ID]; ok {
				continue
			}
			idRelative := command.GetIdRelative()
			if idRelative == "" {
				continue
			}
			components.Xpaths[command.ID] = idRelative
		}
	}
	componentYmlFileBytes, err := yaml.Marshal(&components)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	uploadSideFileBytes, err := json.Marshal(uploadSideFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	ioutil.WriteFile("./component.yml", componentYmlFileBytes, 0664)
	ioutil.WriteFile("outputs/uploadReserve.side", uploadSideFileBytes, 0664)
}
