package main

import (
	creationalpattern "DesignPattern/CreationalPattern"
	principle "DesignPattern/Principle"
	"fmt"
)

func main() {
	//依赖倒转
	principle.DIP()
	fmt.Println()

	//开闭原则
	principle.OCP()
	fmt.Println()

	// 单一职责
	principle.SRP()
	fmt.Println()

	//简单工厂模式
	creationalpattern.TestSF()
	fmt.Println()

	//工厂模式
	creationalpattern.TestF()
	fmt.Println()
}
