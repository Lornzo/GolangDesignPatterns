package example5_test

import (
	"testing"

	"github.com/Lornzo/GolangDesignPatterns/Builder/example5"
)

func TestExample5(t *testing.T) {

	var (
		director     example5.IProfessionDirector
		humanBuilder example5.HumanBuilder
		elfBuilder   example5.ElfBuilder
	)

	director = &example5.WarriorDirector{}
	director.SetActor(&humanBuilder)
	director.Make()
	human := humanBuilder.GetHuman()
	t.Log(human)

	director = &example5.WizardDirector{}
	director.SetActor(&elfBuilder)
	director.Make()
	elf := elfBuilder.GetElf()
	t.Log(elf)

}
