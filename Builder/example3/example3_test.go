package example3_test

import (
	"testing"

	"github.com/Lornzo/GolangDesignPatterns/GolangDesignPatterns/Builder/example3"
)

func TestExample3(t *testing.T) {

	var (
		director example3.IComputerDirector
	)

	amdBuilder := &example3.AMDBuilder{}
	director = example3.OfficeDirector{}
	director.MakeComputer(amdBuilder)
	admComputer := amdBuilder.GetAMDComputer()
	t.Log(admComputer.Description())

	intelBuilder := &example3.IntelBuilder{}
	director = example3.GamingDirector{}
	director.MakeComputer(intelBuilder)
	intelComputer := intelBuilder.GetIntelComputer()
	t.Log(intelComputer.Description())

}
