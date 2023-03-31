package example4

type Kichen interface {
	MakeSet() Set
}

type AmericanKichen struct{}

func (a AmericanKichen) MakeSet() Set {

	return AmericanSet{}
}

type JapaneseKichen struct{}

func (j JapaneseKichen) MakeSet() Set {

	return JapaneseSet{}
}

type ChineseKichen struct{}

func (c ChineseKichen) MakeSet() Set {

	return ChineseSet{}
}
