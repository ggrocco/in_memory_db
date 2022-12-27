package db

type memory struct {
	register map[string]string
}

func newMemory() *memory {
	return &memory{register: make(map[string]string)}
}

func (m *memory) Set(name, value string) {
	m.register[name] = value
}

func (m *memory) Get(name string) string {
	value := m.register[name]
	if value == "" {
		return "Nil"
	}
	return value
}

func (m *memory) Unset(name string) {
	delete(m.register, name)
}

func (m *memory) NumEqualTo(value string) int {
	count := 0
	for _, v := range m.register {
		if v == value {
			count++
		}
	}
	return count
}
