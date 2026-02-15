package tokenbucket

import (
	"fmt"
	"go-rate-limiter/config"
	"runtime"
	"sync"
	"testing"
)

func TestRaceCondition(t *testing.T) {
	config.LoadAppConfigurations("../../config")
	config.AppConfig.RateLimiter.ReqPerMin = 100
	tb := NewTokenBucket()

	var wg sync.WaitGroup
	allowed := 0
	var countMu sync.Mutex

	// Launch 100 goroutines simultaneously
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if tb.Allow() {
				countMu.Lock()
				allowed++
				countMu.Unlock()
			}
		}()
	}

	wg.Wait()

	// Should be exactly 100, but without mutex in TokenBucket,
	// you might see 105, 110, etc. (more than you had tokens!)
	fmt.Printf("Allowed: %d, Remaining tokens: %f, NumCpu: %d\n", allowed, tb.totalTokens, runtime.NumCPU())
}
