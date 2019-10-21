package setting

// XpathSetting 変換を行うxpathの設定の構造体
type XpathSetting struct {
	Xpaths map[string]string `yaml:"xpaths"`
}

// NewXpathSetting 変換を行うxpathの設定の構造体の初期化
func NewXpathSetting() XpathSetting {
	return XpathSetting{
		Xpaths: make(map[string]string),
	}
}

// GetByTemplate テンプレート形式の入力からxpathを取得する
func (xs *XpathSetting) GetByTemplate(template string) string {
	for key := range xs.Xpaths {
		if t := xs.GetTemplate(key); t == template {
			return xs.Xpaths[key]
		}
	}

	return ""
}

// GetTemplate 変換を行う定義のテンプレートの値を取得
// ファイルの場合は {xpath:VAR_NAME} の形式で入力されたものに対して変換を実行
func (xs *XpathSetting) GetTemplate(key string) string {
	return "{xpath:" + key + "}"
}

// IsAlreadyExists すでにxpathが登録済みかどうかを判定する
func (xs *XpathSetting) IsAlreadyExists(val string) bool {
	for _, xpath := range xs.Xpaths {
		if xpath == val {
			return true
		}
	}
	return false
}
