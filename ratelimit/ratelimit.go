package ratelimit

import (
	"time"
)

// RedisRateLimit enforces a rate limit using Redis.
//
// It tracks the number of times an operation is performed within a given time window.
// If the limit is exceeded, it returns `rateLimited = true`; otherwise, it allows the operation.
//
// Parameters:
// - `key` (string): A unique identifier for tracking usage (e.g., user ID, IP address, or custom key).
// - `ttl` (time.Duration): The time window within which the operation is counted.
// - `maxRequests` (int): The maximum number of times the operation is allowed within the time window.
//
// Returns:
// - `rateLimited` (bool): True if the operation limit is exceeded, false otherwise.
// - `currentHits` (int64): The number of times the operation has been performed so far within the time window.
func (c *RateLimitC) RedisRateLimit(
	key string,
	ttl time.Duration,
	maxRequests int,
) (rateLimited bool, currentHits int64) {
	// Increment operation count for the given key.
	currentHits = c.Redis.Incr(key).Val()

	// Set expiration only when the key is first created.
	if currentHits == 1 {
		c.Redis.Expire(key, ttl)
	}

	return currentHits > int64(maxRequests), currentHits
}
