package mgt

import (
	"github.com/Miemiemiemieqiang/translator/core"
	"github.com/Miemiemiemieqiang/translator/translate/csv"
	"github.com/Miemiemiemieqiang/translator/translate/json"
	"github.com/Miemiemiemieqiang/translator/translate/sqltext"
	"github.com/Miemiemiemieqiang/translator/translate/yaml"
)

var DefaultManager = newManager()

func newManager() *Manager {
	manager := NewManager()
	manager.RegisterBid(yaml.NewTranslate())
	manager.RegisterBid(json.NewTranslate())
	manager.RegisterBid(csv.NewTranslate())
	manager.RegisterReader(sqltext.NewTranslate())
	return manager
}

func RegisterReader(reader core.LangReader) {
	DefaultManager.RegisterReader(reader)
}

func RegisterWriter(writer core.LangWriter) {
	DefaultManager.RegisterWriter(writer)
}

func RegisterBid(bid core.LangBid) {
	DefaultManager.RegisterBid(bid)
}
