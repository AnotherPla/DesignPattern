package principle

import (
	"fmt"
)

type ShopClothes struct{}

func (sc *ShopClothes) style() {
	fmt.Println("逛街时候的装扮")
}

type WorkClothes struct{}

func (wc *WorkClothes) style() {
	fmt.Println("工作时候的装扮")
}

func SRP() {
	//逛街时候的类
	sc := ShopClothes{}
	sc.style()

	//工作时候的类
	wc := WorkClothes{}
	wc.style()
}
