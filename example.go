package inmemorycache

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type cacheItem struct {
	Value   interface{}
	Expires time.Time
}

type cache struct {
	data        map[string]cacheItem
	mu          sync.Mutex
	ctx         context.Context
	cancel      context.CancelFunc
	cleanupDone sync.WaitGroup
}

type Cache interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}, ttl time.Duration) error
	Delete(key string) error
	Stop()
}

const defaultCleanupInterval = 1 * time.Minute

func (c *cache) Get(key string) (interface{}, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, ok := c.data[key]
	if !ok {
		return nil, fmt.Errorf("key not found")
	}

	if !item.Expires.IsZero() && time.Now().After(item.Expires) {
		delete(c.data, key)
		return nil, fmt.Errorf("key expired")
	}

	return item.Value, nil
}

func (c *cache) Set(key string, value interface{}, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var expires time.Time
	if ttl > 0 {
		expires = time.Now().Add(ttl)
	}

	c.data[key] = cacheItem{
		Value:   value,
		Expires: expires,
	}

	return nil
}

func (c *cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.data[key]
	if !ok {
		return fmt.Errorf("key not found")
	}

	delete(c.data, key)
	return nil
}

func (c *cache) cleanupExpired() {
	ticker := time.NewTicker(defaultCleanupInterval)
	defer ticker.Stop()

	for {
		select {
		case <-c.ctx.Done():
			return
		case <-ticker.C:
			c.mu.Lock()
			now := time.Now()
			for key, item := range c.data {
				if !item.Expires.IsZero() && now.After(item.Expires) {
					delete(c.data, key)
				}
			}
			c.mu.Unlock()
		}
	}
}

func (c *cache) Stop() {
	c.cancel()
	c.cleanupDone.Wait()
}

func New() Cache {
	ctx, cancel := context.WithCancel(context.Background())
	c := &cache{
		data:   make(map[string]cacheItem),
		mu:     sync.Mutex{},
		ctx:    ctx,
		cancel: cancel,
	}
	c.cleanupDone.Add(1)
	go func() {
		defer c.cleanupDone.Done()
		c.cleanupExpired()
	}()
	return c
}