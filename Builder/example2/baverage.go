package example2

import "fmt"

type IBeverage interface {
	Description() string
}

type Coffee struct {
	Concentration int
	Sugar         int
	Milk          int
}

func (c Coffee) Description() string {
	return fmt.Sprintf("Coffee: %d%%, %dg sugar, %dml milk", c.Concentration, c.Sugar, c.Milk)
}

type Tea struct {
	TeaType     string
	Temperature float64
	Sugar       int
}

func (t Tea) Description() string {
	return fmt.Sprintf("Tea: %s, %fC, %dg sugar", t.TeaType, t.Temperature, t.Sugar)
}

type Juice struct {
	Fruit string
	Sugar int
}

func (j Juice) Description() string {
	return fmt.Sprintf("Juice: %s, %dg sugar", j.Fruit, j.Sugar)
}

type Soda struct {
	Style    string
	Carbonic int
}

func (s Soda) Description() string {
	return fmt.Sprintf("Soda: %s, %d%% carbonic", s.Style, s.Carbonic)
}
