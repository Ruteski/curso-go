package tax

import (
	"github.com/stretchr/testify/mock"
)

type TaxaRepositoryMock struct {
	mock.Mock
}

func (m *TaxaRepositoryMock) SaveTax(tax float64) error {
	args := m.Called(tax)
	return args.Error(0)
}
