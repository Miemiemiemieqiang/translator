package cmd

import (
	"translator/mgt"
	"translator/translate"
)

func Translate(bytes []byte, r, w string) ([]byte, error) {
	manager := mgt.DefaultManager
	read := translate.GetType(r)
	write := translate.GetType(w)
	if read == translate.Unknown || write == translate.Unknown {
		panic("unknown type")
	}

	return manager.Translate(bytes, read, write)
}

func AutoTranslate(manager *mgt.Manager, bytes []byte, write translate.Type) ([]byte, error) {
	read := manager.PredicateReader(bytes)
	return manager.Translate(bytes, read, write)
}
