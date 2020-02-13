package converter

import (
	"io/ioutil"
	"strings"

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

	for template, text := range t.textSetting.GetTemplates(command.Value) {
		command.Value = strings.Replace(command.Value, template, text, -1)
	}

	for template, text := range t.textSetting.GetTemplates(command.Target) {
		command.Target = strings.Replace(command.Target, template, text, -1)
	}
}

// After 実行後処理の記述
func (t *Text) After() {
	componentYmlFileBytes, _ := yaml.Marshal(&t.textSetting)
	ioutil.WriteFile(textSettingsFilePath, componentYmlFileBytes, 0777)
}
