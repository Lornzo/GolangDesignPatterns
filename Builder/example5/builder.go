package example5

type IActorBuilder interface {
	SetStrength(strength int)
	SetAgility(agility int)
	SetIntelligence(intelligence int)
}

type HumanBuilder struct {
	Actor Human
}

func (h *HumanBuilder) SetStrength(strength int) {
	h.Actor.Strength = strength
}

func (h *HumanBuilder) SetAgility(agility int) {
	h.Actor.Agility = agility
}

func (h *HumanBuilder) SetIntelligence(intelligence int) {
	h.Actor.Intelligence = intelligence
}

func (h *HumanBuilder) GetHuman() Human {
	return h.Actor
}

type ElfBuilder struct {
	Actor Elf
}

func (e *ElfBuilder) SetStrength(strength int) {
	e.Actor.Strength = strength
}

func (e *ElfBuilder) SetAgility(agility int) {
	e.Actor.Agility = agility
}

func (e *ElfBuilder) SetIntelligence(intelligence int) {
	e.Actor.Intelligence = intelligence
}

func (e *ElfBuilder) GetElf() Elf {
	return e.Actor
}
