package sqltext

import (
	"github.com/Miemiemiemieqiang/translator/core"
	"github.com/Miemiemiemieqiang/translator/translate"
	"strings"
)

type Translate struct {
}

func NewTranslate() *Translate {
	return &Translate{}
}

var (
	_ core.LangReader = Translate{}
)

func (t Translate) Config() core.Config {
	return core.NewConfig(core.WithType(translate.SQLText))
}

func (t Translate) Predicate([]byte) bool {
	// should be manual specify
	return false
}

func (t Translate) Read(bytes []byte) (*core.Core, error) {
	results := make([]map[string]interface{}, 0)
	headers := make([]string, 0)
	text := string(bytes)
	hf := false
	split := strings.Split(text, "\n")
	for _, v := range split {
		if strings.HasPrefix(v, "|") {
			if !hf {
				headers = splitLine(v)
				hf = true
			} else {
				rows := splitLine(v)
				result := make(map[string]interface{})
				for i, header := range headers {
					result[header] = rows[i]
				}
				results = append(results, result)
			}
		}
	}
	return core.NewCore(results), nil
}

func splitLine(line string) []string {
	ls := make([]string, 0)
	cells := strings.Split(line, "|")
	// cells should drop the first and last
	for _, cell := range cells[1 : len(cells)-1] {
		cell = strings.TrimSpace(cell)
		ls = append(ls, cell)
	}
	return ls
}
