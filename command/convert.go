package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gtongy/sideconv/selenium"
	"github.com/gtongy/sideconv/setting"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

const (
	XPATH_SETTINGS_FILE_PATH = "convert-settings/xpaths.yml"
	SIDE_DIR                 = "sides/"
	OUTPUTS_DIR              = "outputs/"
	PERMMISION_ALL_ALLOW     = 0777
)

func Convert(c *cli.Context) {
	filepath.Walk(SIDE_DIR, walkSideFilePaths)
}

func walkSideFilePaths(sidesPath string, info os.FileInfo, err error) error {
	handleError(err)
	if info.IsDir() {
		return nil
	}
	path := strings.Replace(sidesPath, SIDE_DIR, "", 1)
	convertExec(path)
	return nil
}

func convertExec(filePath string) {
	var uploadSideFile selenium.SideFile
	raw, err := ioutil.ReadFile(SIDE_DIR + filePath)
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
	ioutil.WriteFile(XPATH_SETTINGS_FILE_PATH, componentYmlFileBytes, PERMMISION_ALL_ALLOW)
	uploadSideFileBytes, err := json.Marshal(uploadSideFile)
	handleError(err)
	outputFilePath := OUTPUTS_DIR + filePath
	outputDir := filepath.Dir(outputFilePath)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, PERMMISION_ALL_ALLOW)
	}
	ioutil.WriteFile(outputFilePath, uploadSideFileBytes, PERMMISION_ALL_ALLOW)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
