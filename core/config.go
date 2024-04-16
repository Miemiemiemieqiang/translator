package core

import "translator/translate"

type Config struct {
	Type translate.Type
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
