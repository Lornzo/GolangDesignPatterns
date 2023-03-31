package example4_test

import (
	"fmt"
	"testing"

	"github.com/Lornzo/GolangDesignPatterns/GolangDesignPatterns/FactoryMethod/example4"
)

func TestExample4(t *testing.T) {
	american := example4.AmericanKichen{}
	japanese := example4.JapaneseKichen{}
	chinese := example4.ChineseKichen{}

	americanSet := american.MakeSet()
	japaneseSet := japanese.MakeSet()
	chineseSet := chinese.MakeSet()

	custom(americanSet)
	custom(japaneseSet)
	custom(chineseSet)

}

func custom(set example4.Set) {
	fmt.Println("湯品", set.GetSoup())
	fmt.Println("主菜", set.GetMain())
	fmt.Println("甜點", set.GetDessert())
}
