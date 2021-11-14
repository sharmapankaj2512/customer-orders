package mocks

import "github.com/stretchr/testify/mock"

type OrderRepositoryMock struct {
	mock.Mock
}

func NewOrderRepositoryMock() *OrderRepositoryMock {
	return new(OrderRepositoryMock)
}

func (m *OrderRepositoryMock) ExpectFindIsCalled(customerId int) *OrderRepositoryMock {
	m.On("Find", customerId).Return(nil)
	return m
}