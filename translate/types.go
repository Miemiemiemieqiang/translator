package translate

import "strings"

type Type int32

const (
	Unknown Type = 1 << iota
	JSON
	YAML
	CSV
	SQLText
)

func GetType(name string) Type {
	name = strings.ToLower(name)
	switch name {
	case "json":
		return JSON
	case "yaml", "yml":
		return YAML
	case "csv":
		return CSV
	case "sqltext", "sql":
		return SQLText
	default:
		return Unknown
	}
}
