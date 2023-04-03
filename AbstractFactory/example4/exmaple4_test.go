package example4_test

import (
	"testing"

	"github.com/Lornzo/GolangDesignPatterns/GolangDesignPatterns/AbstractFactory/example4"
)

func TestExample4(t *testing.T) {
	var client example4.Client

	client.System = &example4.Linux{}
	client.Run()

	client.System = &example4.Windows{}
	client.Run()
}
