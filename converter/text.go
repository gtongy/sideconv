package converter

import (
	"io/ioutil"

	"github.com/gtongy/sideconv/selenium"
	"github.com/gtongy/sideconv/setting"
	"gopkg.in/yaml.v2"
)

const (
	textSettingsFilePath = "convert-settings/text.yml"
)

// Text 変換を行うテキストの構造体
type Text struct {
	uploadSideFile *selenium.SideFile
	textSetting    *setting.TextSetting
}

// NewText 変換を行うテキストの構造体の初期化
func NewText(uploadSideFile *selenium.SideFile) Text {
	textSetting := setting.NewTextSetting()
	textSettingRaw, _ := ioutil.ReadFile(textSettingsFilePath)
	yaml.Unmarshal(textSettingRaw, &textSetting)
	return Text{
		uploadSideFile: uploadSideFile,
		textSetting:    &textSetting,
	}
}

// Exec 処理の実行
func (t *Text) Exec(testKey int, commandKey int) {
	command := &t.uploadSideFile.Tests[testKey].Commands[commandKey]

	if value := t.textSetting.GetByTemplate(command.Value); value != "" {
		command.Value = value
	}
	if target := t.textSetting.GetByTemplate(command.Target); target != "" {
		command.Target = target
	}
}

// After 実行後処理の記述
func (t *Text) After() {
	componentYmlFileBytes, _ := yaml.Marshal(&t.textSetting)
	ioutil.WriteFile(textSettingsFilePath, componentYmlFileBytes, 0777)
}
