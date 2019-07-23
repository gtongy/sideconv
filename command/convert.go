package command

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gtongy/sideconv/converter"
	sideconvError "github.com/gtongy/sideconv/error"
	"github.com/gtongy/sideconv/selenium"
	"github.com/urfave/cli"
)

const (
	XPATH_SETTINGS_FILE_PATH = "convert-settings/xpaths.yml"
	INPUTS_DIR               = "inputs/"
	OUTPUTS_DIR              = "outputs/"
	PERMMISION_ALL_ALLOW     = 0777
)

func Convert(c *cli.Context) {
	filepath.Walk(INPUTS_DIR, walkSideFilePaths)
}

func walkSideFilePaths(sidesPath string, info os.FileInfo, err error) error {
	sideconvError.HandleError(err)
	if info.IsDir() {
		return nil
	}
	path := strings.Replace(sidesPath, INPUTS_DIR, "", 1)
	convertExec(path)
	return nil
}

func convertExec(filePath string) {
	var uploadSideFile *selenium.SideFile
	raw, err := ioutil.ReadFile(INPUTS_DIR + filePath)
	sideconvError.HandleError(err)
	json.Unmarshal(raw, &uploadSideFile)
	xpathConverter := converter.NewXpath(uploadSideFile)
	fileConverter := converter.NewFile(uploadSideFile)
	textConverter := converter.NewText(uploadSideFile)
	for testKey, test := range uploadSideFile.Tests {
		for commandKey, command := range test.Commands {
			xpathConverter.Exec(testKey, commandKey, command)
			fileConverter.Exec(testKey, commandKey, command)
			textConverter.Exec(testKey, commandKey, command)
		}
	}
	xpathConverter.After()
	fileConverter.After()
	textConverter.After()
	uploadSideFileBytes, err := json.Marshal(uploadSideFile)
	sideconvError.HandleError(err)
	outPutFile(filePath, uploadSideFileBytes)
}

func outPutFile(filePath string, uploadSideFileBytes []byte) {
	outputFilePath := OUTPUTS_DIR + filePath
	outputDir := filepath.Dir(outputFilePath)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, PERMMISION_ALL_ALLOW)
	}
	ioutil.WriteFile(outputFilePath, uploadSideFileBytes, PERMMISION_ALL_ALLOW)
}
