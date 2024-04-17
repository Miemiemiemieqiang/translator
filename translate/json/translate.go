package json

import (
	"encoding/json"
	"github.com/Miemiemiemieqiang/translator/core"
	"github.com/Miemiemiemieqiang/translator/translate"
)

type Translate struct {
	Prefix string
	Indent string
}

var (
	_ core.LangBid = Translate{}
)

func NewTranslate() Translate {
	return Translate{
		Prefix: "",
		Indent: "  ",
	}
}

func (t Translate) Predicate(bytes []byte) bool {
	output := make(map[string]interface{})
	err := json.Unmarshal(bytes, &output)
	return err == nil
}

func (t Translate) Read(bytes []byte) (*core.Core, error) {
	var data interface{}
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return core.NewCore(data), nil
}

func (t Translate) Config() core.Config {
	return core.NewConfig(core.WithType(translate.JSON), core.WithRecurse(true))
}

func (t Translate) Write(c *core.Core) ([]byte, error) {
	out, err := json.MarshalIndent(c.Data, t.Prefix, t.Indent)
	return out, err
}
