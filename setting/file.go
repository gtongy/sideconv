package setting

import "os"

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

// GetTemplate 変換を行う定義のテンプレートの値を取得
// ファイルの場合は {file:VAR_NAME} の形式で入力されたものに対して変換を実行
func (fs *FileSetting) GetTemplate(key string) string {
	return "{file:" + key + "}"
}

// GetByTemplate テンプレート形式の入力からxpathを取得する
func (fs *FileSetting) GetByTemplate(template string) string {
	for key := range fs.Files {
		if t := fs.GetTemplate(key); t == template {
			return fs.Files[key]
		}
	}

	return ""
}
