package mocks

import (
    "time"
)

// MockCache é uma estrutura de mock para o cache.
type MockCache struct {
    Cache map[string]float64
    TimeUpdated time.Time
}

// NewMockCache cria uma nova instância do MockCache.
func NewMockCache() *MockCache {
    return &MockCache{
        Cache: make(map[string]float64),
        TimeUpdated: time.Now(),
    }
}

// Get simula a operação de obter um valor do cache.
func (mc *MockCache) Get(key string) (float64, bool) {
    value, ok := mc.Cache[key]
    return value, ok
}

// Set simula a operação de definir um valor no cache.
func (mc *MockCache) Set(key string, value float64, expiration time.Duration) {
    mc.Cache[key] = value
    mc.TimeUpdated = time.Now()
}

// GetAll simula a operação de obter todos os itens do cache.
func (mc *MockCache) GetAll() (map[string]float64, time.Time, error) {
    return mc.Cache, mc.TimeUpdated, nil
}

// GetTimestamp simula a operação de obter o tempo da última atualização do cache.
func (mc *MockCache) GetTimestamp() time.Time {
    return mc.TimeUpdated
}
