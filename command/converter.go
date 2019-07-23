package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gtongy/sideconv/selenium"
	"github.com/gtongy/sideconv/setting"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

const (
	XPATH_SETTINGS_FILE_PATH = "./convert-settings/xpaths.yml"
)

func ConvertExec(c *cli.Context) {
	var uploadSideFile selenium.SideFile
	raw, err := ioutil.ReadFile("./sides/uploadReserve.side")
	json.Unmarshal(raw, &uploadSideFile)
	xpathSetting := setting.NewXpathSetting()
	xpathSettingRaw, err := ioutil.ReadFile(XPATH_SETTINGS_FILE_PATH)
	handleError(err)
	yaml.Unmarshal(xpathSettingRaw, &xpathSetting)
	for testKey, test := range uploadSideFile.Tests {
		for commandKey, command := range test.Commands {
			xpathKey := command.GetTargetXpathKey(xpathSetting.Xpaths)
			if xpathKey != "" {
				uploadSideFile.Tests[testKey].Commands[commandKey].Target = strings.Replace(command.Target, xpathSetting.GetTemplate(xpathKey), xpathSetting.Xpaths[xpathKey], -1)
			}
			if _, ok := xpathSetting.Xpaths[command.ID]; ok {
				continue
			}
			idRelative := command.GetIdRelative()
			if idRelative == "" || xpathSetting.IsAlreadyExists(idRelative) {
				continue
			}
			xpathSetting.Xpaths[command.ID] = idRelative
		}
	}
	componentYmlFileBytes, err := yaml.Marshal(&xpathSetting)
	handleError(err)
	ioutil.WriteFile(XPATH_SETTINGS_FILE_PATH, componentYmlFileBytes, 0664)
	uploadSideFileBytes, err := json.Marshal(uploadSideFile)
	handleError(err)
	ioutil.WriteFile("outputs/uploadReserve.side", uploadSideFileBytes, 0664)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
