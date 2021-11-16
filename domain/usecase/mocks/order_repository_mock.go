package mocks

import (
	. "customer-orders/domain/model"

	"github.com/stretchr/testify/mock"
)

type OrderRepositoryMock struct {
	mock.Mock
}

func NewOrderRepositoryMock() *OrderRepositoryMock {
	return new(OrderRepositoryMock)
}

func (m *OrderRepositoryMock) Find(customerId int) []Order {
	args := m.Called(customerId)
	return args.Get(0).([]Order)
}

func (m *OrderRepositoryMock) Delete(order Order) {
	m.Called(order)
}

func (m *OrderRepositoryMock) ExpectFindIsCalled(customerId int) *OrderRepositoryMock {
	m.On("Find", customerId).Return([]Order{})
	return m
}

func (m *OrderRepositoryMock) ExpectFindReturns(customerId int, orders []Order) *OrderRepositoryMock {
	m.On("Find", customerId).Return(orders)
	return m
}

func (m *OrderRepositoryMock) ExpectOrdersAreDeleted(orders []Order) *OrderRepositoryMock {
	for order := range orders {
		m.On("Delete", order).Return(nil)
	}
	return m
}
