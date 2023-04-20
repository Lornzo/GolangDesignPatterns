package example4

import (
	"fmt"
	"strings"
)

type ISandwich interface {
	Description() string
}

type AbstractSandwich struct {
	Bread   string
	Size    int
	Toasted bool
	Recipt  []string
}

type BeefSandwich struct {
	AbstractSandwich
}

func (b BeefSandwich) Description() string {
	var description string = "Beef Sandwich with "

	if b.Toasted {
		description += "toasted "
	}

	description += fmt.Sprintf("%d inch %s and ", b.Size, b.Bread)
	description += strings.Join(b.Recipt, ", ")

	return description
}

type ChickenSandwich struct {
	AbstractSandwich
}

func (c ChickenSandwich) Description() string {
	var description string = "Chicken Sandwich with "

	if c.Toasted {
		description += "toasted "
	}

	description += fmt.Sprintf("%d inch %s and ", c.Size, c.Bread)
	description += strings.Join(c.Recipt, ", ")

	return description
}
