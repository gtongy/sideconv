package setting

type TextSetting struct {
	Texts map[string]string `yaml:"texts"`
}

func NewTextSetting() TextSetting {
	return TextSetting{
		Texts: make(map[string]string),
	}
}

func (fs *TextSetting) GetTemplate(key string) string {
	return "{text:" + key + "}"
}

type Texts map[string]string
