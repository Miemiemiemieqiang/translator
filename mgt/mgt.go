package mgt

import (
	"sync"
	"translator/core"
	"translator/translate"
)

type Manager struct {
	sync.RWMutex
	readers map[translate.Type]core.LangReader
	writers map[translate.Type]core.LangWriter
}

func NewManager() *Manager {
	return &Manager{
		readers: make(map[translate.Type]core.LangReader),
		writers: make(map[translate.Type]core.LangWriter),
	}
}

func (m *Manager) RegisterReader(reader core.LangReader) {
	m.Lock()
	defer m.Unlock()
	m.appendReader(reader)
}

func (m *Manager) RegisterWriter(writer core.LangWriter) {
	m.Lock()
	defer m.Unlock()
	m.appendWriter(writer)
}

func (m *Manager) RegisterBid(bid core.LangBid) {
	m.Lock()
	defer m.Unlock()
	m.appendReader(bid)
	m.appendWriter(bid)
}

func (m *Manager) appendReader(reader core.LangReader) {
	t := reader.Config().Type
	if t == 0 {
		panic("reader type is empty")
	}
	if m.readers[t] != nil {
		panic("reader type is duplicated")
	}

	m.readers[t] = reader
}

func (m *Manager) appendWriter(writer core.LangWriter) {
	t := writer.Config().Type
	if t == 0 {
		panic("writer name is empty")
	}
	if m.writers[t] != nil {
		panic("writer name is duplicated")
	}

	m.writers[t] = writer
}

func (m *Manager) GetReaders() []core.LangReader {
	m.RLock()
	defer m.RUnlock()

	result := make([]core.LangReader, 0)
	for _, reader := range m.readers {
		result = append(result, reader)
	}
	return result
}

func (m *Manager) GetWriters() []core.LangWriter {
	m.RLock()
	defer m.RUnlock()

	result := make([]core.LangWriter, 0)
	for _, writer := range m.writers {
		result = append(result, writer)
	}
	return result
}

func (m *Manager) GetReader(t translate.Type) core.LangReader {
	m.RLock()
	defer m.RUnlock()

	if reader, ok := m.readers[t]; ok {
		return reader
	} else {
		panic("reader not found")
	}
}

func (m *Manager) GetWriter(t translate.Type) core.LangWriter {
	m.RLock()
	defer m.RUnlock()

	if writer, ok := m.writers[t]; ok {
		return writer
	} else {
		panic("writer not found")
	}
}

func (m *Manager) PredicateReader(bytes []byte) translate.Type {
	m.RLock()
	defer m.RUnlock()

	for _, reader := range m.readers {
		if reader.Predicate(bytes) {
			return reader.Config().Type
		}
	}
	panic("reader not found")
}

func (m *Manager) Translate(bytes []byte, r, w translate.Type) ([]byte, error) {
	reader := m.GetReader(r)
	writer := m.GetWriter(w)
	read, err := reader.Read(bytes)
	if err != nil {
		return nil, err
	}
	return writer.Write(read)
}
