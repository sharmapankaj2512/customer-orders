package deleteinactivecustomerorders

import (
	s "customer-orders/domain/usecase/shared"
)

func DeleteInActiveCustomerOrders(orderRepository s.OrderRepository) s.Usecase {
	return func(reader s.Reader, writer s.Writer) {

	}
}