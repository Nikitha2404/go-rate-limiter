package ratelimiter

import (
	"go-rate-limiter/config"
	"go-rate-limiter/ratelimiter/tokenbucket"
)

var (
	tokenBucketStrategy = "token-bucket"
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
	}
	return r
}

func (r *RateLimiter) Allow() bool {
	return r.rateLimiter.Allow()
}
