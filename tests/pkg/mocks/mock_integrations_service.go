package mocks

import (
    "time"
)

// MockIntegrationService é uma estrutura de mock para serviços de integração com APIs externas.
type MockIntegrationService struct {
    SelicRate float64
    IpcaRate  float64
    SelicTime time.Time
    IpcaTime  time.Time
}

// NewMockIntegrationService cria uma nova instância do MockIntegrationService.
func NewMockIntegrationService() *MockIntegrationService {
    return &MockIntegrationService{}
}

// GetSELICData simula a operação de obter dados da taxa SELIC de uma API externa.
func (m *MockIntegrationService) GetSELICData() (float64, time.Time, error) {
    return m.SelicRate, m.SelicTime, nil
}

// GetIPCAData simula a operação de obter dados do IPCA de uma API externa.
func (m *MockIntegrationService) GetIPCAData() (float64, time.Time, error) {
    return m.IpcaRate, m.IpcaTime, nil
}
