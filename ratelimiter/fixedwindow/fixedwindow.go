package fixedwindow

import (
	"go-rate-limiter/config"
	"log"
	"sync"
	"time"
)

type FixedWindow struct {
	totalReq       int
	capacity       int
	lastRefillTime time.Time
	mu             sync.Mutex
}

func NewFixedWindow() *FixedWindow {
	reqPerMin := config.AppConfig.RateLimiter.ReqPerMin
	return &FixedWindow{
		totalReq:       0,
		capacity:       reqPerMin,
		lastRefillTime: time.Now(),
	}
}

func (f *FixedWindow) Allow() bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.reset()
	f.totalReq += 1
	if f.totalReq > f.capacity {
		return false
	}
	return true
}

func (f *FixedWindow) reset() {
	now := time.Now()
	elapsed := now.Sub(f.lastRefillTime).Seconds()
	if int(elapsed) >= 60 {
		f.totalReq = 0
		f.lastRefillTime = now
	}
	log.Print("total Tokens", f.totalReq, " last refill time", f.lastRefillTime," elapsed",elapsed)
}
