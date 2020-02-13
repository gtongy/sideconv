package setting

import (
	"os"
	"strings"
)

const (
	filesDirPath = "/files"
)

// FileSetting 変換を行うファイルの設定の構造体
type FileSetting struct {
	Files   map[string]string `yaml:"files"`
	BaseURL string            `yaml:"-"`
}

// NewFileSetting 変換を行うファイルの設定の構造体の初期化
func NewFileSetting() FileSetting {
	p, _ := os.Getwd()
	return FileSetting{
		Files:   make(map[string]string),
		BaseURL: p + filesDirPath,
	}
}

// getTemplate 変換を行う定義のテンプレートの値を取得
// ファイルの場合は {file:VAR_NAME} の形式で入力されたものに対して変換を実行
func (fs *FileSetting) getTemplate(key string) string {
	return "{file:" + key + "}"
}

// GetTemplates テンプレート形式が含まれた入力からfileを取得する
func (fs *FileSetting) GetTemplates(s string) map[string]string {
	templates := make(map[string]string)

	for key := range fs.Files {
		template := fs.getTemplate(key)
		if strings.Contains(s, template) {
			templates[template] = fs.Files[key]
		}
	}

	return templates
}
