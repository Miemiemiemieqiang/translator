package mgt

import (
	"translator/core"
	"translator/translate/json"
	"translator/translate/sqltext"
	"translator/translate/yaml"
)

var DefaultManager = newManager()

func newManager() *Manager {
	manager := NewManager()
	manager.RegisterBid(yaml.NewTranslate())
	manager.RegisterBid(json.NewTranslate())
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
