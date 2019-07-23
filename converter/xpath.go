package converter

import (
	"io/ioutil"
	"strings"

	"github.com/gtongy/sideconv/selenium"
	"github.com/gtongy/sideconv/setting"
	"gopkg.in/yaml.v2"
)

const (
	XPATH_SETTINGS_FILE_PATH = "convert-settings/xpath.yml"
	PERMMISION_ALL_ALLOW     = 0777
)

type Xpath struct {
	uploadSideFile *selenium.SideFile
	xpathSetting   *setting.XpathSetting
}

func NewXpath(uploadSideFile *selenium.SideFile) Xpath {
	xpathSetting := setting.NewXpathSetting()
	xpathSettingRaw, _ := ioutil.ReadFile(XPATH_SETTINGS_FILE_PATH)
	yaml.Unmarshal(xpathSettingRaw, &xpathSetting)
	return Xpath{
		uploadSideFile: uploadSideFile,
		xpathSetting:   &xpathSetting,
	}
}

func (xp *Xpath) Exec(testKey int, commandKey int, command selenium.Command) {
	xpathKey := command.GetTargetXpathKey(xp.xpathSetting.Xpaths)
	if xpathKey != "" {
		xp.uploadSideFile.Tests[testKey].Commands[commandKey].Target =
			strings.Replace(command.Target, xp.xpathSetting.GetTemplate(xpathKey), xp.xpathSetting.Xpaths[xpathKey], -1)
	}
	if _, ok := xp.xpathSetting.Xpaths[command.ID]; ok {
		return
	}
	idRelative := command.GetIdRelative()
	if idRelative == "" || xp.xpathSetting.IsAlreadyExists(idRelative) {
		return
	}
	xp.xpathSetting.Xpaths[command.ID] = idRelative
}

func (xp *Xpath) After() {
	componentYmlFileBytes, _ := yaml.Marshal(&xp.xpathSetting)
	ioutil.WriteFile(XPATH_SETTINGS_FILE_PATH, componentYmlFileBytes, PERMMISION_ALL_ALLOW)
}
