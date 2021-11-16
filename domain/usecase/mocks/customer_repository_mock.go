package mocks

import (
	"customer-orders/domain/model"

	"github.com/stretchr/testify/mock"
)

type CustomerRepositoryMock struct {
	mock.Mock
}

func NewCustomerRepositoryMock() *CustomerRepositoryMock {
	return new(CustomerRepositoryMock)
}

func (m *CustomerRepositoryMock) Find(customerId int) model.Customer {
	args := m.Called(customerId)
	if args.Get(0) != nil {
		return args.Get(0).(model.Customer)
	}
	return nil
}

func (m *CustomerRepositoryMock) ExpectFindDoesNotReturnCustomer(customerId int) *CustomerRepositoryMock {
	m.On("Find", customerId).Return(nil)
	return m
}

func (m *CustomerRepositoryMock) ExpectFindReturnsInactiveCustomer(customerId int) *CustomerRepositoryMock {
	m.On("Find", customerId).Return(&CustomerStub{customerId, false})
	return m
}

func (m *CustomerRepositoryMock) ExpectFindReturnsActiveCustomer(customerId int) *CustomerRepositoryMock {
	m.On("Find", customerId).Return(&CustomerStub{customerId, true})
	return m
}

type CustomerStub struct {
	id int
	active bool
}

func (s *CustomerStub) ID() int {
	return s.id
}

func (s *CustomerStub) IsNotActive() bool {
	return !s.active
}
