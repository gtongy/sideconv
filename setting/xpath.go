package setting

type XpathSetting struct {
	Xpaths map[string]string `yaml:"xpaths"`
}

func NewXpathSetting() XpathSetting {
	return XpathSetting{
		Xpaths: make(map[string]string),
	}
}

type Xpaths map[string]string
