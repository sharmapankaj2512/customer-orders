package deleteinactivecustomerorders

import (
	m "customer-orders/domain/model"
	s "customer-orders/domain/usecase/shared"
	"errors"
)

func DeleteInActiveCustomerOrders(
	customerRepository s.CustomerRepository,
	orderRepository s.OrderRepository) s.Usecase {
	return func(reader s.Reader, writer s.Writer) {
		customerId := reader.Read().(int)
		customer := customerRepository.Find(customerId)
		if no(customer) {
			writer.Write(errors.New("customer not found"))
			return
		}
		if customer.IsNotActive() {
			writer.Write(errors.New("customer is not active"))
			return
		}
		orders := orderRepository.Find(customerId)
		orderRepository.Delete(orders)
		writer.Write("orders deleted")
	}
}

func no(customer m.Customer) bool {
	return customer == nil
}
