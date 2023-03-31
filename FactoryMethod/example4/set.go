package example4

type Set interface {
	GetSoup() string
	GetMain() string
	GetDessert() string
}

type AmericanSet struct{}

func (a AmericanSet) GetSoup() string {
	return "玉米濃湯"
}

func (a AmericanSet) GetMain() string {
	return "紐約客牛排"
}

func (a AmericanSet) GetDessert() string {
	return "甜甜圈"
}

type JapaneseSet struct{}

func (j JapaneseSet) GetSoup() string {
	return "味噌湯"
}

func (j JapaneseSet) GetMain() string {
	return "壽司"
}

func (j JapaneseSet) GetDessert() string {
	return "麻糬丸子"
}

type ChineseSet struct{}

func (c ChineseSet) GetSoup() string {
	return "酸辣湯"
}

func (c ChineseSet) GetMain() string {
	return "牛肉麵"
}

func (c ChineseSet) GetDessert() string {
	return "豆沙包子"
}
