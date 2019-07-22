package selenium

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
	ID      string        `json:"id"`
	Comment string        `json:"comment"`
	Command string        `json:"command"`
	Target  string        `json:"target"`
	Targets []interface{} `json:"targets"`
	Value   string        `json:"value"`
}
