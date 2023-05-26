package example1

type MonsterManager struct {
	prototypes map[string]MonsterPrototype
}

func (m *MonsterManager) SetPrototype(name string, prototype MonsterPrototype) {
	if m.prototypes == nil {
		m.prototypes = make(map[string]MonsterPrototype)
	}
	m.prototypes[name] = prototype
}

func (m *MonsterManager) GetPrototype(name string) MonsterPrototype {
	return m.prototypes[name]
}
