package main

import (
	creationalpattern "DesignPattern/CreationalPattern"
	principle "DesignPattern/Principle"
	structuralpattern "DesignPattern/StructuralPattern"
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

	//单例模式  原子操作+互斥锁实现线程安全的单例模式
	s := creationalpattern.GetInstance()
	s.Something()

	//单例模式  Once 实现线程安全的单例模式
	onceS := creationalpattern.GetInstance2()
	onceS.Something()
	fmt.Println()

	//代理模式
	structuralpattern.TestProxy()

}
