package deleteinactivecustomerorders

import (
	. "customer-orders/domain/usecase/mocks"
	"testing"
)

func TestDeleteNoOrderWhenTheCustomerDoesNotExist(t *testing.T) {
	customerId := 1233
	orderRepository := NewOrderRepositoryMock().ExpectFindIsCalled(customerId)
	reader := NewReaderMock().ExpectReturns(customerId)
	writer := NewWriterMock().ExpectReceivesError("customer not found")
	usecase := DeleteInActiveCustomerOrders(orderRepository)

	usecase(reader, writer)

	assertOrderIsNotFound(t, reader, writer, orderRepository)
}

func assertOrderIsNotFound(
	t *testing.T,
	reader *ReaderMock,
	writer *WriterMock,
	orderRepository *OrderRepositoryMock) {
	reader.AssertExpectations(t)
	orderRepository.AssertExpectations(t)
	orderRepository.AssertNotCalled(t, "Save")
	writer.AssertExpectations(t)	
}
