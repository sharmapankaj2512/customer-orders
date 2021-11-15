package deleteinactivecustomerorders

import (
	m "customer-orders/domain/model"
	s "customer-orders/domain/usecase/shared"
	"errors"
)

func DeleteInActiveCustomerOrders(orderRepository s.OrderRepository) s.Usecase {
	return func(reader s.Reader, writer s.Writer) {
		customerId := reader.Read().(int)
		orders := orderRepository.Find(customerId)
		if no(orders) {
			writer.Write(errors.New("customer not found"))
		}
	}
}

func no(orders []m.Order) bool {
	return len(orders) == 0
}
