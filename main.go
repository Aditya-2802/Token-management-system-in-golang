package main

import (
	"fmt"
	"sync"
	"token-management-system/simulation"
	"token-management-system/token"
)

func main() {

	tokens := token.InitializeTokens(1000)

	var wg sync.WaitGroup

	token.ResetUsagePeriodically(tokens, &wg)
	for {
		var operations int
		fmt.Print("Enter number of operations to simulate: ")
		fmt.Scanln(&operations)

		if operations == 0 {
			break
		}
		simulation.RunSimulation(tokens, operations)

		token.DisplayResults(tokens)
	}
	wg.Wait()
}
