package yaml

import (
	"github.com/Miemiemiemieqiang/translator/core"
	"github.com/Miemiemiemieqiang/translator/translate"
	"gopkg.in/yaml.v2"
)

type Translate struct {
}

func NewTranslate() *Translate {
	return &Translate{}
}

var (
	_ core.LangBid = Translate{}
)

func (t Translate) Predicate(bytes []byte) bool {
	var data interface{}
	err := yaml.Unmarshal(bytes, &data)
	return err == nil
}

func (t Translate) Read(bytes []byte) (*core.Core, error) {
	var data interface{}
	err := yaml.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return core.NewCore(data), nil
}

func (t Translate) Config() core.Config {
	return core.NewConfig(core.WithType(translate.YAML), core.WithRecurse(true))
}

func (t Translate) Write(core *core.Core) ([]byte, error) {
	out, err := yaml.Marshal(core.Data)
	return out, err
}
