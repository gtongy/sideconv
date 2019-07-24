package selenium

import (
	"strings"
)

// SideFile .sideファイルの構造体
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

// Test テストケースの構造体
type Test struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Commands []Command `json:"commands"`
}

// Command 各ステップの構造体
type Command struct {
	ID      string     `json:"id"`
	Comment string     `json:"comment"`
	Command string     `json:"command"`
	Target  string     `json:"target"`
	Targets [][]string `json:"targets"`
	Value   string     `json:"value"`
}

// GetIDRelative idRelativeの値
func (c *Command) GetIDRelative() string {
	for _, target := range c.Targets {
		if target[1] == "xpath:idRelative" {
			return target[0]
		}
	}
	return ""
}

// GetTargetXpathKey Target内からXpathの設定のkey名を取得
func (c *Command) GetTargetXpathKey(xpaths map[string]string) string {
	for key := range xpaths {
		if strings.Index(c.Target, key) != -1 {
			return key
		}
	}
	return ""
}

// GetValueFileKey Value内からfileの設定のkey名を取得
func (c *Command) GetValueFileKey(files map[string]string) string {
	for key := range files {
		if strings.Index(c.Value, key) != -1 {
			return key
		}
	}
	return ""
}

// GetTargetTextKey Target内からfileの設定のkey名を取得
func (c *Command) GetTargetTextKey(texts map[string]string) string {
	for key := range texts {
		if strings.Index(c.Target, key) != -1 {
			return key
		}
	}
	return ""
}

// GetValueTextKey Value内からfileの設定のkey名を取得
func (c *Command) GetValueTextKey(texts map[string]string) string {
	for key := range texts {
		if strings.Index(c.Value, key) != -1 {
			return key
		}
	}
	return ""
}
