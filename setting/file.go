package setting

import "os"

const (
	FILE_PATH = "/files"
)

type FileSetting struct {
	Files   map[string]string `yaml:"files"`
	BaseUrl string            `yaml:"-"`
}

func NewFileSetting() FileSetting {
	p, _ := os.Getwd()
	return FileSetting{
		Files:   make(map[string]string),
		BaseUrl: p + FILE_PATH,
	}
}

func (fs *FileSetting) GetTemplate(key string) string {
	return "{file:" + key + "}"
}

type Files map[string]string
