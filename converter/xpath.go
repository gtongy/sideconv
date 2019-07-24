package converter

import (
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
	xpathKey := xp.uploadSideFile.Tests[testKey].Commands[commandKey].GetTargetXpathKey(xp.xpathSetting.Xpaths)
	if xpathKey != "" {
		xp.uploadSideFile.Tests[testKey].Commands[commandKey].Target =
			strings.Replace(
				xp.uploadSideFile.Tests[testKey].Commands[commandKey].Target,
				xp.xpathSetting.GetTemplate(xpathKey),
				xp.xpathSetting.Xpaths[xpathKey], -1)
	}
	if _, ok := xp.xpathSetting.Xpaths[xp.uploadSideFile.Tests[testKey].Commands[commandKey].ID]; ok {
		return
	}
	idRelative := xp.uploadSideFile.Tests[testKey].Commands[commandKey].GetIDRelative()
	if idRelative == "" || xp.xpathSetting.IsAlreadyExists(idRelative) {
		return
	}
	xp.xpathSetting.Xpaths[xp.uploadSideFile.Tests[testKey].Commands[commandKey].ID] = idRelative
}

// After 実行後処理の記述
func (xp *Xpath) After() {
	componentYmlFileBytes, _ := yaml.Marshal(&xp.xpathSetting)
	ioutil.WriteFile(xpathSettingsFilePath, componentYmlFileBytes, 0777)
}
