package db

import "sync"

type memory struct {
	register map[string]string
	sync.RWMutex
}

func newMemory() *memory {
	return &memory{make(map[string]string), sync.RWMutex{}}
}

func (m *memory) Set(name, value string) {
	m.Lock()
	defer m.Unlock()

	m.register[name] = value
}

func (m *memory) Get(name string) string {
	m.RLock()
	defer m.RUnlock()

	value := m.register[name]
	if value == "" {
		return "Nil"
	}
	return value
}

func (m *memory) Unset(name string) {
	m.Lock()
	defer m.Unlock()

	delete(m.register, name)
}

func (m *memory) NumEqualTo(value string) int {
	m.RLock()
	defer m.RUnlock()

	count := 0
	for _, v := range m.register {
		if v == value {
			count++
		}
	}
	return count
}
