package tokenbucket

import (
	"go-rate-limiter/config"
	"log"
	"math"
	"sync"
	"time"
)

type TokenBucket struct {
	totalTokens    float64
	bucketSize     float64
	refillRate     float64
	lastRefillTime time.Time
	mu             sync.Mutex
}

func NewTokenBucket() *TokenBucket {
	reqPerMin := config.AppConfig.RateLimiter.TokenBucket.ReqPerMin
	return &TokenBucket{
		totalTokens:    float64(reqPerMin),
		bucketSize:     float64(reqPerMin),
		refillRate:     float64(reqPerMin) / float64(60),
		lastRefillTime: time.Now(),
	}
}

func (t *TokenBucket) Allow() bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.refill()
	if t.totalTokens >= 1 {
		t.totalTokens -= 1
		return true
	}

	return false
}

// refill func takes care of continuous refil intervals based on elapsed time
func (t *TokenBucket) refill() {
	now := time.Now()

	elapsed := now.Sub(t.lastRefillTime).Seconds()
	tokensToAdd := elapsed * t.refillRate
	t.totalTokens = math.Min(tokensToAdd+t.totalTokens, t.bucketSize) // to cap the total token at bucket size
	t.lastRefillTime = now
	log.Print("total Tokens", t.totalTokens, " refill rate", t.refillRate, " last refill time", t.lastRefillTime)
}
