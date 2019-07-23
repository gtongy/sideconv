package setting

type XpathSetting struct {
	Xpaths map[string]string `yaml:"xpaths"`
}

func NewXpathSetting() XpathSetting {
	return XpathSetting{
		Xpaths: make(map[string]string),
	}
}

func (xs *XpathSetting) GetTemplate(key string) string {
	return "{xpath:" + key + "}"
}

func (xs *XpathSetting) IsAlreadyExists(val string) bool {
	for _, xpath := range xs.Xpaths {
		if xpath == val {
			return true
		}
	}
	return false
}

type Xpaths map[string]string
