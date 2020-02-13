package setting

import "strings"

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
func (ts *TextSetting) getTemplate(key string) string {
	return "{text:" + key + "}"
}

// GetTemplates テンプレート形式が含まれた入力からfileを取得する
func (ts *TextSetting) GetTemplates(s string) map[string]string {
	templates := make(map[string]string)

	for key := range ts.Texts {
		template := ts.getTemplate(key)
		if strings.Contains(s, template) {
			templates[template] = ts.Texts[key]
		}
	}

	return templates
}
