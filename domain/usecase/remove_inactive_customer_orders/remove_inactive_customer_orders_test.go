package removeinactivecustomerorders

import (
	m "customer-orders/domain/model"
	. "customer-orders/domain/usecase/mocks"
	"testing"
)

func TestDoNotDeleteOrdersWhenCustomerDoesNotExist(t *testing.T) {
	customerId := 1233
	orderRepository := NewOrderRepositoryMock()
	customerRepository := NewCustomerRepositoryMock().ExpectFindDoesNotReturnCustomer(customerId)
	reader := NewReaderMock().ExpectReturns(customerId)
	writer := NewWriterMock().ExpectReceivesError("customer not found")
	usecase := RemoveInActiveCustomerOrders(customerRepository, orderRepository)

	usecase(reader, writer)

	assertOrderIsNotDeleted(t, reader, writer, customerRepository, orderRepository)
}

func TestDoNotDeleteOrdersOfAnActiveCustomer(t *testing.T) {
	customerId := 4444
	orderRepository := NewOrderRepositoryMock()
	customerRepository := NewCustomerRepositoryMock().ExpectFindReturnsInactiveCustomer(customerId)
	reader := NewReaderMock().ExpectReturns(customerId)
	writer := NewWriterMock().ExpectReceivesError("customer is not active")
	usecase := RemoveInActiveCustomerOrders(customerRepository, orderRepository)

	usecase(reader, writer)

	assertOrderIsNotDeleted(t, reader, writer, customerRepository, orderRepository)
}

func TestDeleteOrdersOfAnInActiveCustomer(t *testing.T) {
	customerId := 4445
	orders := []m.Order{OrderStub{}}
	customerRepository := NewCustomerRepositoryMock().
		ExpectFindReturnsActiveCustomer(customerId)
	orderRepository := NewOrderRepositoryMock().
		ExpectFindReturns(customerId, orders).
		ExpectDeleteIsCalledWith(orders)
	reader := NewReaderMock().ExpectReturns(customerId)
	writer := NewWriterMock().ExpectReceives("orders deleted")
	usecase := RemoveInActiveCustomerOrders(customerRepository, orderRepository)

	usecase(reader, writer)

	assertOrderIsDeleted(t, reader, writer, customerRepository, orderRepository)
}

func assertOrderIsDeleted(
	t *testing.T,
	reader *ReaderMock,
	writer *WriterMock,
	customerRepository *CustomerRepositoryMock,
	orderRepository *OrderRepositoryMock) {
	reader.AssertExpectations(t)
	customerRepository.AssertExpectations(t)
	orderRepository.AssertExpectations(t)
	writer.AssertExpectations(t)
}

func assertOrderIsNotDeleted(
	t *testing.T,
	reader *ReaderMock,
	writer *WriterMock,
	customerRepository *CustomerRepositoryMock,
	orderRepository *OrderRepositoryMock) {
	reader.AssertExpectations(t)
	customerRepository.AssertExpectations(t)
	orderRepository.AssertExpectations(t)
	writer.AssertExpectations(t)
}

type OrderStub struct{}
