package structuralpattern

import (
	"fmt"
)

type Goods struct {
	kind   string
	isTrue bool
}

// ---------------------抽象层---------------
// 抽象的购物主题
type Shopping interface {
	Buy(good *Goods)
}

// ---------------------实现层---------------
// 购物的具体实现

type USAShopping struct{}

func (usa *USAShopping) Buy(good *Goods) {
	fmt.Println("went to us buy  ", good.kind)
}

type UKShopping struct{}

func (uk *UKShopping) Buy(good *Goods) {
	fmt.Println("went to us buy  ", good.kind)
}

// 代理实现
// 代理实际上也是一个Shopping类，不过它组合了其它购物方式，
// 内部实现一些代理逻辑并调用其它购物方式的方法
type Proxy struct {
	shopping Shopping
}

// 代理的构造方法，目的是组合其它的shopping
func NewProxy(shopping Shopping) Shopping {
	return &Proxy{shopping: shopping}
}

// 代理的验证方法
func (p *Proxy) vertify(good *Goods) bool {
	fmt.Println("vartifieing goods.....")
	if !good.isTrue {
		fmt.Println(good.kind, " is fake, no buying")
	}
	return good.isTrue
}

// 代理的海关检验方法
func (p *Proxy) check(good *Goods) {
	fmt.Println(good.kind, "  passing")
}

func (p *Proxy) Buy(good *Goods) {
	// 1.验货
	if p.vertify(good) {
		// 2. 购买
		p.shopping.Buy(good)
		// 3. 入海关
		p.check(good)
	}
}

// ---------------------- 业务逻辑层------------------
func TestProxy() {
	g1 := Goods{
		"ps5",
		true,
	}

	g2 := Goods{
		"psp",
		false,
	}

	var shopping = new(USAShopping)

	// 若不使用代理
	// 则需要将所有逻辑代码写在逻辑层
	/* if !g1.isTrue {
		fmt.Println(g1.kind, " is fake, no buying")
	} else {
		shopping.Buy(&g1)
		fmt.Println(g1.kind, "  passing")
	} */

	// 使用代理
	var overseaProxy = NewProxy(shopping)
	overseaProxy.Buy(&g1)
	overseaProxy.Buy(&g2)
}
