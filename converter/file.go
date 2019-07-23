package converter

import (
	"io/ioutil"
	"strings"

	"github.com/gtongy/sideconv/selenium"
	"github.com/gtongy/sideconv/setting"
	"gopkg.in/yaml.v2"
)

const (
	FILE_SETTINGS_FILE_PATH = "convert-settings/file.yml"
)

type File struct {
	uploadSideFile *selenium.SideFile
	fileSetting    *setting.FileSetting
}

func NewFile(uploadSideFile *selenium.SideFile) File {
	fileSetting := setting.NewFileSetting()
	fileSettingRaw, _ := ioutil.ReadFile(FILE_SETTINGS_FILE_PATH)
	yaml.Unmarshal(fileSettingRaw, &fileSetting)
	return File{
		uploadSideFile: uploadSideFile,
		fileSetting:    &fileSetting,
	}
}

func (f *File) Exec(testKey int, commandKey int, command selenium.Command) {
	fileKey := command.GetValueFileKey(f.fileSetting.Files)
	if fileKey != "" {
		f.uploadSideFile.Tests[testKey].Commands[commandKey].Value =
			strings.Replace(command.Value, f.fileSetting.GetTemplate(fileKey), f.fileSetting.BaseUrl+"/"+f.fileSetting.Files[fileKey], -1)
	}
}

func (f *File) After() {
	componentYmlFileBytes, _ := yaml.Marshal(&f.fileSetting)
	ioutil.WriteFile(FILE_SETTINGS_FILE_PATH, componentYmlFileBytes, PERMMISION_ALL_ALLOW)
}
