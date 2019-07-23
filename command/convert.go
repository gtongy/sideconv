package command

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	sideconvError "github.com/gtongy/sideconv/error"
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
	sideconvError.HandleError(err)
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
	sideconvError.HandleError(err)
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
	sideconvError.HandleError(err)
	ioutil.WriteFile(XPATH_SETTINGS_FILE_PATH, componentYmlFileBytes, PERMMISION_ALL_ALLOW)
	uploadSideFileBytes, err := json.Marshal(uploadSideFile)
	sideconvError.HandleError(err)
	outputFilePath := OUTPUTS_DIR + filePath
	outputDir := filepath.Dir(outputFilePath)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, PERMMISION_ALL_ALLOW)
	}
	ioutil.WriteFile(outputFilePath, uploadSideFileBytes, PERMMISION_ALL_ALLOW)
	sideconvError.HandleError(err)
}
