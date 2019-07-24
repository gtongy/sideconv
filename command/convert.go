package command

import (
	"encoding/json"
	"fmt"
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
	inputsDirPath  = "inputs/"
	outputsDirPath = "outputs/"
)

// Convert 変換処理の実行
func Convert(c *cli.Context) {
	filepath.Walk(inputsDirPath, walkSideFilePaths)
	fmt.Println("変換完了")
}

// walkSideFilePaths 変換を行うファイルのパスの走査
func walkSideFilePaths(sidesPath string, info os.FileInfo, err error) error {
	sideconvError.HandleError(err)
	if info.IsDir() {
		return nil
	}
	path := strings.Replace(sidesPath, inputsDirPath, "", 1)
	if filepath.Ext(path) != ".side" {
		fmt.Printf("%s はsideファイルではないためスキップします\n", path)
		return nil
	}
	convertExec(path)
	return nil
}

// convertExec 変換の実行
func convertExec(filePath string) {
	var uploadSideFile *selenium.SideFile
	raw, err := ioutil.ReadFile(inputsDirPath + filePath)
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

// outPutFile ファイルの書き出し
func outPutFile(filePath string, uploadSideFileBytes []byte) {
	outputFilePath := outputsDirPath + filePath
	outputDir := filepath.Dir(outputFilePath)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, 0777)
	}
	ioutil.WriteFile(outputFilePath, uploadSideFileBytes, 0777)
}
