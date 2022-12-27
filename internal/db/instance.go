package db

import "errors"

var ErrorNoTransaction = errors.New("NO TRANSACTION")

type db struct {
	memory       *memory
	transactions []*transaction
}

type BaseOperation interface {
	Set(name, value string)
	Get(name string) string
	Unset(name string)
	NumEqualTo(value string) int
}

type TransactionOperation interface {
	Begin()
	Commit() error
	Rollback() error
}

type DB interface {
	BaseOperation
	TransactionOperation
}

func New() DB {
	return &db{memory: newMemory()}
}

func (db *db) Set(name, value string) {
	db.currentTransaction(func(t BaseOperation) {
		t.Set(name, value)
	})
}

func (db *db) Get(name string) (output string) {
	db.currentTransaction(func(t BaseOperation) {
		output = t.Get(name)
	})

	return
}

func (db *db) Unset(name string) {
	db.currentTransaction(func(t BaseOperation) {
		t.Unset(name)
	})
}

func (db *db) NumEqualTo(value string) (count int) {
	db.currentTransaction(func(t BaseOperation) {
		count = t.NumEqualTo(value)
	})

	return
}

func (db *db) currentTransaction(action func(t BaseOperation)) {
	if len(db.transactions) > 0 {
		action(db.transactions[len(db.transactions)-1])
	} else {
		action(db.memory)
	}
}

func (db *db) Begin() {
	db.transactions = append(db.transactions, newTransaction())
}

func (db *db) Commit() error {
	if len(db.transactions) == 0 {
		return ErrorNoTransaction
	}

	for i := len(db.transactions); i > 0; i-- {
		db.transactions[i-1].replay(db.memory)
	}

	db.transactions = nil
	return nil
}

func (db *db) Rollback() error {
	if len(db.transactions) == 0 {
		return ErrorNoTransaction
	}

	db.transactions = db.transactions[:len(db.transactions)-1]
	return nil
}
