package db

type replyAction func(BaseOperation)

type transaction struct {
	*memory
	replyActions []replyAction
}

func newTransaction() *transaction {
	return &transaction{memory: newMemory()}
}

func (t *transaction) replay(db BaseOperation) {
	for _, action := range t.replyActions {
		action(db)
	}
}

func (t *transaction) Set(name, value string) {
	t.replyActions = append(t.replyActions, func(bo BaseOperation) { bo.Set(name, value) })
	t.memory.Set(name, value)
}

func (t *transaction) Get(name string) string {
	// reply will not do any effect
	return t.memory.Get(name)
}

func (t *transaction) Unset(name string) {
	t.replyActions = append(t.replyActions, func(bo BaseOperation) { bo.Unset(name) })
	t.memory.Unset(name)
}

func (t *transaction) NumEqualTo(value string) int {
	// reply will not do any effect
	return t.memory.NumEqualTo(value)
}
