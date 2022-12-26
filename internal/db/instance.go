package db

type db struct {
	memory map[string]string
}

type DB interface {
	Set(name, value string)
	Get(name string) string
	Unset(name string)
	NumEqualTo(value string) int
}

func New() DB {
	return &db{memory: make(map[string]string)}
}

func (db *db) Set(name, value string) {
	db.memory[name] = value
}

func (db *db) Get(name string) string {
	value := db.memory[name]
	if value == "" {
		return "Nil"
	}
	return value
}

func (db *db) Unset(name string) {
	delete(db.memory, name)
}

func (db *db) NumEqualTo(value string) int {
	count := 0
	for _, v := range db.memory {
		if v == value {
			count++
		}
	}
	return count
}
