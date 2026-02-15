package leakybucket

import (
	"go-rate-limiter/config"
	"log"
	"math"
	"sync"
	"time"
)

type LeakyBucket struct {
	totalTokens  float64
	bucketSize   float64
	leakRate     float64
	lastLeakTime time.Time
	mu           sync.Mutex
}

func NewLeakyBucket() *LeakyBucket {
	reqPerMin := config.AppConfig.RateLimiter.ReqPerMin
	return &LeakyBucket{
		totalTokens:  0,
		bucketSize:   float64(reqPerMin),
		leakRate:     float64(reqPerMin) / float64(60),
		lastLeakTime: time.Now(),
	}
}

func (l *LeakyBucket) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.leak()
	if l.totalTokens < l.bucketSize {
		l.totalTokens += 1
		return true
	}

	return false
}

func (l *LeakyBucket) leak() {
	now := time.Now()
	elapsed := now.Sub(l.lastLeakTime).Seconds()
	tokensToRemove := elapsed * l.leakRate
	l.totalTokens = math.Max(l.totalTokens-tokensToRemove, 0)
	l.lastLeakTime = now
	log.Print("total Tokens", l.totalTokens, " leak rate", l.leakRate, " last leak time", l.lastLeakTime)
}
