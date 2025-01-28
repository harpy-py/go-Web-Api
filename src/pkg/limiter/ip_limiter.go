package limiter

import (
	"sync"
	"golang.org/x/time/rate"
)

// Limiter holds the rate limiter for each IP address
type Limiter struct {
	mu       *sync.RWMutex
	ips map[string]*rate.Limiter
	r rate.Limit
	b int
}

// NewLimiter creates a new Limiter instance
func NewLimiter(r rate.Limit, b int) *Limiter {
	limit := &Limiter{
		ips: make(map[string]*rate.Limiter),
		mu: &sync.RWMutex{},
		r: r,
		b: b,
	}
	return limit
}

// GetLimiter retrieves or creates a rate limiter for a specific IP
func (l *Limiter) AddIP(ip string) *rate.Limiter {
	l.mu.Lock()
	defer l.mu.Unlock()

	limiter := rate.NewLimiter(l.r, l.b)
	l.ips[ip] = limiter
	return limiter
}

func (l *Limiter) GetLimiter(ip string) *rate.Limiter {
	l.mu.Lock()
	limiter, exists := l.ips[ip]

	if !exists {
		l.mu.Unlock()
		return l.AddIP(ip)
	}
	l.mu.Unlock()
	return limiter
}