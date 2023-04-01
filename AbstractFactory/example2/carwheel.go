package example2

type ICarWheel interface {
	Description() string
}

type RegularWheel struct{}

func (r RegularWheel) Description() string {
	return "普通的輪胎"
}

type LuxuryWheel struct{}

func (l LuxuryWheel) Description() string {
	return "高級的輪胎"
}
