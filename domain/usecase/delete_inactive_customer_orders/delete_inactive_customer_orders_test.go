package deleteinactivecustomerorders

import (
	. "customer-orders/domain/usecase/mocks"
	"testing"
)

func TestDoNotDeleteOrdersWhenCustomerDoesNotExist(t *testing.T) {
	customerId := 1233
	customerRepository := NewCustomerRepositoryMock().ExpectFindIsCalled(customerId)
	reader := NewReaderMock().ExpectReturns(customerId)
	writer := NewWriterMock().ExpectReceivesError("customer not found")
	usecase := DeleteInActiveCustomerOrders(customerRepository)

	usecase(reader, writer)

	assertOrderIsNotDeleted(t, reader, writer, customerRepository)
}

func TestDoNotDeleteOrdersOfAnActiveCustomer(t *testing.T) {
	t.Skip("to be implemented")
}

func TestDeleteOrdersOfAnInActiveCustomer(t *testing.T) {
	t.Skip("to be implemented")
}

func assertOrderIsNotDeleted(
	t *testing.T,
	reader *ReaderMock,
	writer *WriterMock,
	customerRepository *CustomerRepositoryMock) {
	reader.AssertExpectations(t)
	customerRepository.AssertExpectations(t)
	customerRepository.AssertNotCalled(t, "Save")
	writer.AssertExpectations(t)
}
