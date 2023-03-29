package creationalpattern

import "fmt"

// -----------------抽象层---------------
type Fruit interface {
	show()
	//...
}

//-------------------实现层--------------
type Apple struct{}

func (apple *Apple) show() {
	fmt.Println("create a apple...")
}

type Banana struct{}

func (banana *Banana) show() {
	fmt.Println("create a banana...")
}

type Grape struct{}

func (grape *Grape) show() {
	fmt.Println("create a grape...")
}

//------------------工厂-----------------
type FruitFactory struct{}

func (factory *FruitFactory) CreateFruit(kind string) (Fruit, error) {
	var fruit Fruit
	switch kind {
	case "apple":
		// apple的初始化代码
		fruit = new(Apple)
	case "banana":
		fruit = new(Banana)
	case "grape":
		fruit = new(Grape)
	default:
		return nil, fmt.Errorf("not support yet")
	}
	return fruit, nil
}

//------------------业务逻辑---------------
func TestSF() {
	factory := new(FruitFactory)
	grape, err := factory.CreateFruit("grape")
	if err == nil {
		grape.show()
	} else {
		fmt.Println(err)
	}

	orange, err := factory.CreateFruit("orange")
	if err == nil {
		orange.show()
	} else {
		fmt.Println(err)
	}

}
