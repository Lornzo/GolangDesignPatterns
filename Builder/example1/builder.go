package example1

type IBuilder interface {
	BuildRoom(num int)
	BuildLivingRoom(num int)
	BuildBathRoom(num int)
	BuildBedRoom(num int)
	BuildKitchen(num int)
	BuildBalcony(num int)
	BuildGarage(num int)
}

type StrawHouseBuilder struct {
	house StrawHouse
}

func (s *StrawHouseBuilder) BuildRoom(num int) {
	s.house.UseStrawNum += 3 * num
	s.house.Room = num
}

func (s *StrawHouseBuilder) BuildLivingRoom(num int) {
	s.house.UseStrawNum += 5 * num
	s.house.LivingRoom = num
}

func (s *StrawHouseBuilder) BuildBathRoom(num int) {
	s.house.UseStrawNum += 2 * num
	s.house.BathRoom = num
}

func (s *StrawHouseBuilder) BuildBedRoom(num int) {
	s.house.UseStrawNum += 4 * num
	s.house.BedRoom = num
}

func (s *StrawHouseBuilder) BuildKitchen(num int) {
	s.house.UseStrawNum += 2 * num
	s.house.Kitchen = num
}

func (s *StrawHouseBuilder) BuildBalcony(num int) {
	s.house.UseStrawNum += 1 * num
	s.house.Balcony = num
}

func (s *StrawHouseBuilder) BuildGarage(num int) {
	s.house.UseStrawNum += 10 * num
	s.house.Garage = num
}

func (s *StrawHouseBuilder) GetHouse() StrawHouse {
	return s.house
}

type WoodHouseBuilder struct {
	house WoodHouse
}

func (w *WoodHouseBuilder) BuildRoom(num int) {
	w.house.UseWoodNum += 3 * num
	w.house.Room = num
}

func (w *WoodHouseBuilder) BuildLivingRoom(num int) {
	w.house.UseWoodNum += 5 * num
	w.house.LivingRoom = num
}

func (w *WoodHouseBuilder) BuildBathRoom(num int) {
	w.house.UseWoodNum += 2 * num
	w.house.BathRoom = num
}

func (w *WoodHouseBuilder) BuildBedRoom(num int) {
	w.house.UseWoodNum += 4 * num
	w.house.BedRoom = num
}

func (w *WoodHouseBuilder) BuildKitchen(num int) {
	w.house.UseWoodNum += 2 * num
	w.house.Kitchen = num
}

func (w *WoodHouseBuilder) BuildBalcony(num int) {
	w.house.UseWoodNum += 1 * num
	w.house.Balcony = num
}

func (w *WoodHouseBuilder) BuildGarage(num int) {
	w.house.UseWoodNum += 10 * num
	w.house.Garage = num
}

func (w *WoodHouseBuilder) GetHouse() WoodHouse {
	return w.house
}

type BrickHouseBuilder struct {
	house BrickHouse
}

func (b *BrickHouseBuilder) BuildRoom(num int) {
	b.house.UseBrickNum += 3 * num
	b.house.Room = num
}

func (b *BrickHouseBuilder) BuildLivingRoom(num int) {
	b.house.UseBrickNum += 5 * num
	b.house.LivingRoom = num
}

func (b *BrickHouseBuilder) BuildBathRoom(num int) {
	b.house.UseBrickNum += 2 * num
	b.house.BathRoom = num
}

func (b *BrickHouseBuilder) BuildBedRoom(num int) {
	b.house.UseBrickNum += 4 * num
	b.house.BedRoom = num
}

func (b *BrickHouseBuilder) BuildKitchen(num int) {
	b.house.UseBrickNum += 2 * num
	b.house.Kitchen = num
}

func (b *BrickHouseBuilder) BuildBalcony(num int) {
	b.house.UseBrickNum += 1 * num
	b.house.Balcony = num
}

func (b *BrickHouseBuilder) BuildGarage(num int) {
	b.house.UseBrickNum += 10 * num
	b.house.Garage = num
}

func (b *BrickHouseBuilder) GetHouse() BrickHouse {
	return b.house
}
