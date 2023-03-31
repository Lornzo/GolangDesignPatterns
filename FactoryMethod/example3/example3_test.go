package example3_test

import (
	"fmt"
	"testing"

	"github.com/Lornzo/GolangDesignPatterns/GolangDesignPatterns/FactoryMethod/example3"
)

func TestExample3(t *testing.T) {
	generateTempalte(example3.DarkCreator{})
	generateTempalte(example3.SimpleCreator{})
	generateTempalte(example3.BrightCreator{})
}

func generateTempalte(creator example3.StyleCreator) {

	var style example3.Style = creator.CreateStyle()

	fmt.Println(style.GetTemplate())

}
