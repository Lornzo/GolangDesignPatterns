package example5

import "fmt"

type BagStore struct {
	Factory IBagFactory
}

func (b BagStore) ShowAllBags() {
	totebag := b.Factory.CreateToteBag()
	shoulderbag := b.Factory.CreateShoulderBag()
	fmt.Println("we have the following bags in our store:", totebag, shoulderbag)
}
