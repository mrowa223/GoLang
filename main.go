package main

import (
	"fmt"
	"pkg"
)

var (
	base = pkg.BasePc{}
	home = pkg.HomePc{
		Cpu:          4,
		GraphicsCard: 1,
		Wrapper:      base,
	}
	server = pkg.ServerPc{
		Cpu:     16,
		Memory:  256,
		Wrapper: base,
	}
)

func main() {
	fmt.Printf("Base [%f]", base.GetPrice())
}
