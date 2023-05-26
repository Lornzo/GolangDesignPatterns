package example1_test

import (
	"testing"

	"github.com/Lornzo/GolangDesignPatterns/Prototype/example1"
)

func TestExample1(t *testing.T) {

	type monsterManager interface {
		SetPrototype(name string, prototype example1.MonsterPrototype)
		GetPrototype(name string) example1.MonsterPrototype
	}

	var (
		manager monsterManager = &example1.MonsterManager{}
	)

	manager.SetPrototype("orc", &example1.ConcreteMonster{
		Name:    "orc",
		Attack:  10,
		Defense: 20,
		Health:  100,
	})

	orcPrototypeClone := manager.GetPrototype("orc").Clone().(*example1.ConcreteMonster)
	orcPrototypeClone.Name = "orc clone"

	orc1 := orcPrototypeClone.Clone().(*example1.ConcreteMonster)
	orc2 := orcPrototypeClone.Clone().(*example1.ConcreteMonster)
	t.Log("orc1", orc1)
	t.Log("orc2", orc2)

}
