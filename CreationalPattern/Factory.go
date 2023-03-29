package creationalpattern

// -----------------抽象层---------------

type Factory interface {
	createFruit() Fruit
}

//-------------------实现层--------------

type AppleFactory struct{}

func (appleFac *AppleFactory) createFruit() Fruit {
	return new(Apple)
}

type BananaFactory struct{}

func (bananaFac *BananaFactory) createFruit() Fruit {
	return new(Banana)
}

type GrapeFactory struct{}

func (grapeFac *GrapeFactory) createFruit() Fruit {
	return new(Grape)
}

//------------------业务逻辑---------------
func TestF() {
	// 创建一个苹果工厂
	// 用苹果工厂生产苹果
	var appleFac Factory
	appleFac = new(AppleFactory)

	var apple Fruit
	apple = appleFac.createFruit()
	apple.show()

	// 创建一个葡萄工厂
	// 用葡萄工厂生产葡萄
	grapeFac := new(GrapeFactory)
	grape := grapeFac.createFruit()
	grape.show()

}
