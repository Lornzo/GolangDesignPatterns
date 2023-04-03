package example3

type IMenu interface {
	GetMenuItems() []string
}

type EnglishMenu struct{}

func (e EnglishMenu) GetMenuItems() []string {
	return []string{
		"Breakfast: Milk and Cereal",
		"Lunch: Sandwich",
		"Dinner: Chicken and Rice",
	}
}

type ChineseMenu struct{}

func (c ChineseMenu) GetMenuItems() []string {
	return []string{
		"早餐: 牛奶和谷物",
		"午餐: 三明治",
		"晚餐: 鸡肉和米饭",
	}
}
