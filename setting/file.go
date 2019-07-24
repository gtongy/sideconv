package setting

import "os"

const (
	filesDirPath = "/files"
)

// FileSetting 変換を行うファイルの設定の構造体
type FileSetting struct {
	Files   Files  `yaml:"files"`
	BaseURL string `yaml:"-"`
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

// Files 設定ファイルに書き込まれたファイルの設定群の構造体
type Files map[string]string
