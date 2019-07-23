package selenium

import (
	"strings"

	"github.com/gtongy/sideconv/setting"
)

type SideFile struct {
	ID      string `json:"id"`
	Version string `json:"version"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	Tests   []Test `json:"tests"`
	Suites  []struct {
		ID             string `json:"id"`
		Name           string `json:"name"`
		PersistSession bool   `json:"persistSession"`
		Parallel       bool   `json:"parallel"`
		Timeout        int    `json:"timeout"`
		// testのid群がここに入力。もしここでidを指定しなければ処理は実行出来ない
		Tests []string `json:"tests"`
	} `json:"suites"`
	Urls    []string      `json:"urls"`
	Plugins []interface{} `json:"plugins"`
}

type Test struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Commands []Command `json:"commands"`
}

type Command struct {
	ID      string     `json:"id"`
	Comment string     `json:"comment"`
	Command string     `json:"command"`
	Target  string     `json:"target"`
	Targets [][]string `json:"targets"`
	Value   string     `json:"value"`
}

func (c *Command) GetIdRelative() string {
	for _, target := range c.Targets {
		if target[1] == "xpath:idRelative" {
			return target[0]
		}
	}
	return ""
}

func (c *Command) GetTargetXpathKey(xpaths setting.Xpaths) string {
	for key := range xpaths {
		if strings.Index(c.Target, key) != -1 {
			return key
		}
	}
	return ""
}

func (c *Command) GetValueFileKey(files setting.Files) string {
	for key := range files {
		if strings.Index(c.Value, key) != -1 {
			return key
		}
	}
	return ""
}
