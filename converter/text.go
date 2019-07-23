package converter

import (
	"io/ioutil"
	"strings"

	"github.com/gtongy/sideconv/selenium"
	"github.com/gtongy/sideconv/setting"
	"gopkg.in/yaml.v2"
)

const (
	TEXT_SETTINGS_FILE_PATH = "convert-settings/text.yml"
)

type Text struct {
	uploadSideFile *selenium.SideFile
	textSetting    *setting.TextSetting
}

func NewText(uploadSideFile *selenium.SideFile) Text {
	textSetting := setting.NewTextSetting()
	textSettingRaw, _ := ioutil.ReadFile(TEXT_SETTINGS_FILE_PATH)
	yaml.Unmarshal(textSettingRaw, &textSetting)
	return Text{
		uploadSideFile: uploadSideFile,
		textSetting:    &textSetting,
	}
}

func (t *Text) Exec(testKey int, commandKey int, command selenium.Command) {
	textValueKey := command.GetValueFileKey(t.textSetting.Texts)
	if textValueKey != "" {
		t.uploadSideFile.Tests[testKey].Commands[commandKey].Value =
			strings.Replace(command.Value, t.textSetting.GetTemplate(textValueKey), t.textSetting.Texts[textValueKey], -1)
	}
	textTargetKey := command.GetTargetTextKey(t.textSetting.Texts)
	if textTargetKey != "" {
		t.uploadSideFile.Tests[testKey].Commands[commandKey].Target =
			strings.Replace(command.Target, t.textSetting.GetTemplate(textTargetKey), t.textSetting.Texts[textTargetKey], -1)
	}
}

func (t *Text) After() {
	componentYmlFileBytes, _ := yaml.Marshal(&t.textSetting)
	ioutil.WriteFile(FILE_SETTINGS_FILE_PATH, componentYmlFileBytes, PERMMISION_ALL_ALLOW)
}
