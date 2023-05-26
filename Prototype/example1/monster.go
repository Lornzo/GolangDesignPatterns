package example1

type MonsterPrototype interface {
	Clone() MonsterPrototype
}

type ConcreteMonster struct {
	Name    string
	Attack  int
	Defense int
	Health  int
}

func (c *ConcreteMonster) Clone() MonsterPrototype {
	var prototype ConcreteMonster = *c
	return &prototype
}
