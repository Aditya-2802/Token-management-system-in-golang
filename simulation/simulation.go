package simulation

import (
	"token-management-system/token"
)

// RunSimulation performs the specified number of operations on the tokens
func RunSimulation(tokens []token.Token, operations int) {
	for i := 0; i < operations; i++ {
		selectedToken := token.SelectToken(tokens)
		selectedToken.UsageCount++
	}
}
