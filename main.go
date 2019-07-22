package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"bitbucket.org/hameesys/sidegen/selenium"
)

func main() {
	var loginSideFile selenium.SideFile
	loginRaw, err := ioutil.ReadFile("./setup/login.side")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(loginRaw, &loginSideFile)
	var uploadSideFile selenium.SideFile
	raw, err := ioutil.ReadFile("./sides/uploadReserve.side")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &uploadSideFile)
	for i, test := range uploadSideFile.Tests {
		uploadSideFile.Tests[i].Commands = append(loginSideFile.Tests[0].Commands, test.Commands...)
	}
	uploadSideFileBytes, err := json.Marshal(uploadSideFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	ioutil.WriteFile("outputs/uploadReserve.side", uploadSideFileBytes, 0664)
}
