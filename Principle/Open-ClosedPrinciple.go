package principle

import "fmt"

type AbstractBanker interface {
	DoBusi()
}

type SaveBanker struct{}

func (sb *SaveBanker) DoBusi() {
	fmt.Println("进行存款操作")
}

type PayBanker struct{}

func (pb *PayBanker) DoBusi() {
	fmt.Println("进行支付操作")
}

type InvestBanker struct{}

func (ib *InvestBanker) DoBusi() {
	fmt.Println("进行投资操作")
}

func OCP() {
	sb := &SaveBanker{}
	sb.DoBusi()

	pb := &PayBanker{}
	pb.DoBusi()

	ib := &InvestBanker{}
	ib.DoBusi()

}
