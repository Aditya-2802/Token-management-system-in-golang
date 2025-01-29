package token

import (
	"fmt"
	"math/rand"
	"time"
)

// Token represents a single token with ID and usage count
type Token struct {
	ID         string
	UsageCount int
}

// InitializeTokens creates and initializes a list of tokens
func InitializeTokens(count int) []Token {
	tokens := make([]Token, count)
	for i := 0; i < count; i++ {
		tokens[i] = Token{ID: fmt.Sprintf("Token %d", i+1), UsageCount: 0}
	}
	return tokens
}

// SelectToken selects the token with the least usage count (randomly among ties)
func SelectToken(tokens []Token) *Token {
	minUsage := tokens[0].UsageCount
	candidates := []*Token{}

	for i := range tokens {
		if tokens[i].UsageCount < minUsage {
			minUsage = tokens[i].UsageCount
			candidates = []*Token{&tokens[i]}
		} else if tokens[i].UsageCount == minUsage {
			candidates = append(candidates, &tokens[i])
		}
	}

	rand.Seed(time.Now().UnixNano())
	return candidates[rand.Intn(len(candidates))]
}

// DisplayResults shows the token usage counts and least-used tokens
func DisplayResults(tokens []Token) {
	fmt.Println("Token Usage Counts:")
	leastUsed := []Token{}
	minUsage := tokens[0].UsageCount

	for _, token := range tokens {
		fmt.Printf("%s: %d uses\n", token.ID, token.UsageCount)
		if token.UsageCount < minUsage {
			minUsage = token.UsageCount
			leastUsed = []Token{token}
		} else if token.UsageCount == minUsage {
			leastUsed = append(leastUsed, token)
		}
	}

	fmt.Println("\nLeast Used Tokens:")
	for _, token := range leastUsed {
		fmt.Printf("%s (%d uses)\n", token.ID, token.UsageCount)
	}
}
