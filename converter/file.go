package converter

import (
	"io/ioutil"
	"strings"

	"github.com/gtongy/sideconv/selenium"
	"github.com/gtongy/sideconv/setting"
	"gopkg.in/yaml.v2"
)

const (
	fileSettingsFilePath = "convert-settings/file.yml"
)

// File 変換を行うファイルの構造体
// ここで定義されているファイルとは、ファイルの実体そのものではなく、
// 検証時ファイル名をymlで管理するその時のパスを指しているため注意
type File struct {
	uploadSideFile *selenium.SideFile
	fileSetting    *setting.FileSetting
}

// NewFile ファイルの構造体の初期化
func NewFile(uploadSideFile *selenium.SideFile) File {
	fileSetting := setting.NewFileSetting()
	fileSettingRaw, _ := ioutil.ReadFile(fileSettingsFilePath)
	yaml.Unmarshal(fileSettingRaw, &fileSetting)
	return File{
		uploadSideFile: uploadSideFile,
		fileSetting:    &fileSetting,
	}
}

// Exec 処理の実行
func (f *File) Exec(testKey int, commandKey int) {
	fileKey := f.uploadSideFile.Tests[testKey].Commands[commandKey].GetValueFileKey(f.fileSetting.Files)
	if fileKey != "" {
		f.uploadSideFile.Tests[testKey].Commands[commandKey].Value =
			strings.Replace(
				f.uploadSideFile.Tests[testKey].Commands[commandKey].Value,
				f.fileSetting.GetTemplate(fileKey),
				f.fileSetting.BaseURL+"/"+f.fileSetting.Files[fileKey], -1)
	}
}

// After 実行後処理の記述
func (f *File) After() {
	componentYmlFileBytes, _ := yaml.Marshal(&f.fileSetting)
	ioutil.WriteFile(fileSettingsFilePath, componentYmlFileBytes, 0777)
}
