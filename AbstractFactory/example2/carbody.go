package example2

type ICarBody interface {
	Description() string
}

type VanBody struct{}

func (v VanBody) Description() string {
	return "烤漆斑駁的廂型車"
}

type SportCarBody struct{}

func (s SportCarBody) Description() string {
	return "光亮的跑車"
}
