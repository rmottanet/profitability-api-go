package cache

import (
	"fmt"
    "sync"
    "time"
)


type Cache interface {
    Get(key string) (float64, bool)
    Set(key string, value float64, expiration time.Duration)
    GetAll() (map[string]float64, time.Time, error)
    GetTimestamp() time.Time
}


type RatesCache struct {
    cache       map[string]float64
    timeUpdated time.Time
    mutex       sync.RWMutex
}


func NewRatesCache() *RatesCache {
    return &RatesCache{
        cache: make(map[string]float64),
    }
}


func (cc *RatesCache) Get(key string) (float64, bool) {
    cc.mutex.RLock()
    defer cc.mutex.RUnlock()
    value, ok := cc.cache[key]
    return value, ok
}


func (cc *RatesCache) Set(key string, value float64, expiration time.Duration) {
    cc.mutex.Lock()
    defer cc.mutex.Unlock()

    // no update to zero or null
    if value == 0 {
        return
    }

    // check code on cache
    if _, ok := cc.cache[key]; !ok {
        // insert new codes
        cc.cache[key] = value
        cc.timeUpdated = time.Now().UTC()

        go cc.deleteAfter(key, expiration)
    }
}


func (cc *RatesCache) deleteAfter(key string, expiration time.Duration) {
    <-time.After(expiration)
    cc.mutex.Lock()
    defer cc.mutex.Unlock()
    delete(cc.cache, key)
}


func (cc *RatesCache) GetAll() (map[string]float64, time.Time, error) {
    cc.mutex.RLock()
    defer cc.mutex.RUnlock()

    if cc.cache == nil || cc.timeUpdated.IsZero() {
        return nil, time.Time{}, fmt.Errorf("Index rates cache empty or not updated")
    }

    return cc.cache, cc.timeUpdated, nil
}


func (cc *RatesCache) GetTimestamp() time.Time {
    cc.mutex.RLock()
    defer cc.mutex.RUnlock()

    return cc.timeUpdated
}
