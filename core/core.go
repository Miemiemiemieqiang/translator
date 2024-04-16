package core

type Core struct {
	Data    interface{}
	Recurse bool
}

func NewCore(data interface{}) *Core {
	return &Core{Data: data, Recurse: isRecurse(data)}
}

// isRecurse checks whether the data is a map[string]interface{} or []interface{}
func isRecurse(data interface{}) bool {
	switch value := data.(type) {
	case map[string]interface{}:
		for _, v := range value {
			if _, ok := v.(map[string]interface{}); ok {
				return true
			}
		}
	case []interface{}:
		for _, v := range value {
			if _, ok := v.(map[string]interface{}); ok {
				return true
			}
		}
	}

	return false
}

// LangBid is a bidirectional reader and writer
type LangBid interface {
	LangReader
	LangWriter
}

// LangReader is a reader translator
type LangReader interface {
	// Config returns the configuration of the reader
	Config() Config
	// Predicate can be used to determine whether the reader can read the data
	Predicate([]byte) bool
	// Read reads the data
	Read([]byte) (*Core, error)
}

// LangWriter is a writer translator
type LangWriter interface {
	// Config returns the configuration of the writer
	Config() Config
	// Write writes the data
	Write(*Core) ([]byte, error)
}
