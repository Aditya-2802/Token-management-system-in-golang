package token

import (
	"fmt"
	"sync"
	"time"
)

// ResetUsagePeriodically starts a background goroutine to reset the usage counts every 24 hours
func ResetUsagePeriodically(tokens []Token, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				// Reset all token usage counts
				for i := range tokens {
					tokens[i].UsageCount = 0
					fmt.Println("Token:",tokens[i].ID,"Usage Count Reset to:",tokens[i].UsageCount)
				}

			}
		}
	}()
}
