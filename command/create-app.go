package command

import (
	"errors"
	"fmt"
	"os"

	sideconvError "github.com/gtongy/sideconv/error"
	"github.com/urfave/cli"
)

var (
	setupDirectorys = []string{"/convert-settings", "/outputs", "/inputs", "/files"}
	setupFilePaths  = []string{"/convert-settings/xpath.yml", "/convert-settings/file.yml", "/convert-settings/text.yml"}
)

// CreateApp 初期データの作成コマンド
func CreateApp(c *cli.Context) {
	appName := c.Args().Get(0)
	if appName == "" {
		sideconvError.HandleError(errors.New("アプリ名が入力されていません"))
	}
	for _, dir := range setupDirectorys {
		os.MkdirAll(appName+dir, os.ModePerm)
	}
	for _, filePath := range setupFilePaths {
		file, _ := os.OpenFile(appName+filePath, os.O_WRONLY|os.O_CREATE, 0666)
		defer file.Close()
	}
	fmt.Println("アプリの作成が完了しました")
}
