package main

import (
	"fmt"
	"math/big"
)

func numSteps(s string) int {

	n := new(big.Int)
	n, success := n.SetString(s, 2)
	if !success {
		fmt.Println("error parsing the binary string")
		return 0
	}

	var countSteps int = 0
	one := big.NewInt(1)
	two := big.NewInt(2)

	for n.Cmp(one) != 0 {
		if new(big.Int).And(n, one).Cmp(one) == 0 {
			n.Add(n, one)
		} else {
			n.Div(n, two)
		}
		countSteps++
	}

	return countSteps
}
