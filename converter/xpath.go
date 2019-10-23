package converter

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gtongy/sideconv/selenium"
	"github.com/gtongy/sideconv/setting"
	"gopkg.in/yaml.v2"
)

const (
	xpathSettingsFilePath = "convert-settings/xpath.yml"
)

// Xpath 変換を行うXpathの構造体
type Xpath struct {
	uploadSideFile *selenium.SideFile
	xpathSetting   *setting.XpathSetting
}

// NewXpath 変換を行うXpathの構造体の初期化
func NewXpath(uploadSideFile *selenium.SideFile) Xpath {
	xpathSetting := setting.NewXpathSetting()
	xpathSettingRaw, _ := ioutil.ReadFile(xpathSettingsFilePath)
	yaml.Unmarshal(xpathSettingRaw, &xpathSetting)
	return Xpath{
		uploadSideFile: uploadSideFile,
		xpathSetting:   &xpathSetting,
	}
}

// Exec 処理の実行
func (xp *Xpath) Exec(testKey int, commandKey int) {
	command := &xp.uploadSideFile.Tests[testKey].Commands[commandKey]

	for template, xpath := range xp.xpathSetting.GetTemplates(command.Target) {
		command.Target = strings.Replace(command.Target, template, fmt.Sprintf("xpath=%s", xpath), 1)
	}
	if _, ok := xp.xpathSetting.Xpaths[command.ID]; ok {
		return
	}
	idRelative := command.GetIDRelative()
	if idRelative == "" || xp.xpathSetting.IsAlreadyExists(idRelative) {
		return
	}
	xp.xpathSetting.Xpaths[command.ID] = idRelative
}

// After 実行後処理の記述
func (xp *Xpath) After() {
	componentYmlFileBytes, _ := yaml.Marshal(&xp.xpathSetting)
	ioutil.WriteFile(xpathSettingsFilePath, componentYmlFileBytes, 0777)
}
