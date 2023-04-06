package example1

import "fmt"

type IHouse interface {
	Interduce() string
}

type House struct {
	Room       int
	LivingRoom int
	BathRoom   int
	BedRoom    int
	Kitchen    int
	Balcony    int
	Garage     int
}

func (h *House) Interduce() string {
	return fmt.Sprintf("the house have: %d living room , %d bath room , %d bed room , %d kitchen , %d balcony , %d garage", h.LivingRoom, h.BathRoom, h.BedRoom, h.Kitchen, h.Balcony, h.Garage)
}

type StrawHouse struct {
	UseStrawNum int
	House
}

func (s StrawHouse) Interduce() string {
	var interduce string = "This is a straw house, it use straw: " + fmt.Sprintf("%d", s.UseStrawNum) + " to build it. " + s.House.Interduce()
	return interduce
}

type WoodHouse struct {
	UseWoodNum int
	House
}

func (w WoodHouse) Interduce() string {
	var interduce string = "This is a wood house, it use wood: " + fmt.Sprintf("%d", w.UseWoodNum) + " to build it. " + w.House.Interduce()
	return interduce
}

type BrickHouse struct {
	UseBrickNum int
	House
}

func (b BrickHouse) Interduce() string {
	var interduce string = "This is a brick house, it use brick: " + fmt.Sprintf("%d", b.UseBrickNum) + " to build it. " + b.House.Interduce()
	return interduce
}
