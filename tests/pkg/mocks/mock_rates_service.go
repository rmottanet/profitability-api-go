package mocks

import (
    "profitability/api/pkg/cache"
    "time"
)

// MockCacheService é uma estrutura de mock para o serviço de cache.
type MockCacheService struct {
    Cache cache.Cache
}

// NewMockCacheService cria uma nova instância do MockCacheService.
func NewMockCacheService() *MockCacheService {
    return &MockCacheService{
        Cache: NewMockCache(), // Aqui utilizamos o mock do cache.
    }
}

// GetRatesCached simula a operação de obter as taxas armazenadas em cache.
func (m *MockCacheService) GetRatesCached() map[string]float64 {
    rates, _, _ := m.Cache.GetAll()
    return rates
}

// SetRatesCached simula a operação de definir as taxas no cache.
func (m *MockCacheService) SetRatesCached(rates map[string]float64, expiration time.Duration) {
    for key, value := range rates {
        m.Cache.Set(key, value, expiration)
    }
}
