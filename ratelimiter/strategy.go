package ratelimiter

import (
	"go-rate-limiter/config"
	"go-rate-limiter/ratelimiter/fixedwindow"
	"go-rate-limiter/ratelimiter/leakybucket"
	"go-rate-limiter/ratelimiter/tokenbucket"
)

var (
	tokenBucketStrategy = "token-bucket"
	leakyBucketStrategy = "leaky-bucket"
	fixedWindowStrategy = "fixed-window"
)

type RateLimiterAlgo interface {
	Allow() bool
}

type RateLimiter struct {
	rateLimiter RateLimiterAlgo
}

func InitRateLimiter() *RateLimiter {
	r := &RateLimiter{}
	strategy := config.AppConfig.RateLimiter.Strategy
	switch strategy {
	case tokenBucketStrategy:
		r.rateLimiter = tokenbucket.NewTokenBucket()
	case leakyBucketStrategy:
		r.rateLimiter = leakybucket.NewLeakyBucket()
	case fixedWindowStrategy:
		r.rateLimiter = fixedwindow.NewFixedWindow()
	}
	return r
}

func (r *RateLimiter) Allow() bool {
	return r.rateLimiter.Allow()
}
