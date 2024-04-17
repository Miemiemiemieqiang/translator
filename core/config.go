package core

import "github.com/Miemiemiemieqiang/translator/translate"

type Config struct {
	Type    translate.Type
	Recurse bool
}

type ConfigFunc func(c Config) Config

func NewConfig(fs ...ConfigFunc) Config {
	config := Config{}
	for _, f := range fs {
		config = f(config)
	}
	return config
}

func WithType(t translate.Type) ConfigFunc {
	return func(c Config) Config {
		c.Type = t
		return c
	}
}

func WithRecurse(recurse bool) ConfigFunc {
	return func(c Config) Config {
		c.Recurse = recurse
		return c
	}
}
