package example5

type Actor struct {
	Strength     int
	Agility      int
	Intelligence int
}

type Human struct {
	Actor
}

type Elf struct {
	Actor
}
