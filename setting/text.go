package setting

// TextSetting 変換を行うテキストの設定の構造体
type TextSetting struct {
	Texts map[string]string `yaml:"texts"`
}

// NewTextSetting 変換を行うテキストの設定の構造体の初期化
func NewTextSetting() TextSetting {
	return TextSetting{
		Texts: make(map[string]string),
	}
}

// getTemplate 変換を行う定義のテンプレートの値を取得
// ファイルの場合は {text:VAR_NAME} の形式で入力されたものに対して変換を実行
func (fs *TextSetting) getTemplate(key string) string {
	return "{text:" + key + "}"
}

// GetByTemplate テンプレート形式の入力からtextを取得する
func (fs *TextSetting) GetByTemplate(template string) string {
	for key := range fs.Texts {
		if t := fs.getTemplate(key); t == template {
			return fs.Texts[key]
		}
	}

	return ""
}
