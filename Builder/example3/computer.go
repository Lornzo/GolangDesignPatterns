package example3

import (
	"fmt"
	"strings"
)

type IComputer interface {
	Description() string
}

type Computer struct {
	CPU       []IComputerComponent
	RAM       []IComputerComponent
	SSD       []IComputerComponent
	GraphCard []IComputerComponent
}

func (c Computer) GetComponentsDescription() string {
	var arr []string
	var str string
	var componentArr []string

	if len(c.CPU) > 0 {
		str += fmt.Sprint(len(c.CPU), "CPU :")
		componentArr = make([]string, 0)
		for _, cpu := range c.CPU {
			componentArr = append(componentArr, cpu.GetName())
		}
		str += strings.Join(componentArr, ",")
		arr = append(arr, str)
	}

	if len(c.RAM) > 0 {
		str = fmt.Sprint(len(c.RAM), "RAM :")
		componentArr = make([]string, 0)
		for _, cpu := range c.RAM {
			componentArr = append(componentArr, cpu.GetName())
		}
		str += strings.Join(componentArr, ",")
		arr = append(arr, str)
	}

	if len(c.SSD) > 0 {
		str = fmt.Sprint(len(c.SSD), "SSD :")
		componentArr = make([]string, 0)
		for _, cpu := range c.SSD {
			componentArr = append(componentArr, cpu.GetName())
		}
		str += strings.Join(componentArr, ",")
		arr = append(arr, str)
	}

	if len(c.GraphCard) > 0 {
		str = fmt.Sprint(len(c.GraphCard), "GraphCard :")
		componentArr = make([]string, 0)
		for _, cpu := range c.GraphCard {
			componentArr = append(componentArr, cpu.GetName())
		}
		str += strings.Join(componentArr, ",")
		arr = append(arr, str)
	}

	return strings.Join(arr, ",")
}

type AMDBoard struct{}

func (a AMDBoard) AMDSpecial() string {
	return "AMD special board"
}

type IntelBoard struct{}

func (i IntelBoard) IntelSpecial() string {
	return "Intel special board"
}

type AMDComputer struct {
	Computer
	Board AMDBoard
}

func (a AMDComputer) Description() string {
	return fmt.Sprintf("AMDComputer with %s board have : %s", a.Board.AMDSpecial(), a.GetComponentsDescription())

}

type IntelComputer struct {
	Computer
	Board IntelBoard
}

func (a IntelComputer) Description() string {
	return fmt.Sprintf("IntelComputer with %s board have : %s", a.Board.IntelSpecial(), a.GetComponentsDescription())
}
