package setting

// TextSetting 変換を行うテキストの設定の構造体
type TextSetting struct {
	Texts Texts `yaml:"texts"`
}

// NewTextSetting 変換を行うテキストの設定の構造体の初期化
func NewTextSetting() TextSetting {
	return TextSetting{
		Texts: make(map[string]string),
	}
}

// GetTemplate 変換を行う定義のテンプレートの値を取得
// ファイルの場合は {text:VAR_NAME} の形式で入力されたものに対して変換を実行
func (fs *TextSetting) GetTemplate(key string) string {
	return "{text:" + key + "}"
}

// Texts 設定ファイルに書き込まれたテキストの設定群の構造体
type Texts map[string]string
