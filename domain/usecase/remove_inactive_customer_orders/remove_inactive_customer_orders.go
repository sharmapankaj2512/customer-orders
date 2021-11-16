package removeinactivecustomerorders

import (
	m "customer-orders/domain/model"
	s "customer-orders/domain/usecase/shared"
	"errors"
)

func RemoveInActiveCustomerOrders(
	customerRepository s.CustomerRepository,
	orderRepository s.OrderRepository,
) s.Usecase {
	return func(reader s.Reader, writer s.Writer) {
		customerId := reader.Read().(int)
		s.IOPipe(s.CustomerPipe(
			findCustomer(customerRepository, customerId),
			deleteOrder(orderRepository)),
			writer.Write)
	}
}

func findCustomer(
	customerRepository s.CustomerRepository,
	customerId int,
) func() (m.Customer, error) {
	return func() (m.Customer, error) {
		customer := customerRepository.Find(customerId)
		if no(customer) {
			return nil, errors.New("customer not found")
		}
		if customer.IsNotActive() {
			return nil, errors.New("customer is not active")
		}
		return customer, nil
	}
}

func deleteOrder(
	orderRepository s.OrderRepository,
) func(customer m.Customer) interface{} {
	return func(customer m.Customer) interface{} {
		orders := orderRepository.Find(customer.ID())
		for order := range orders {
			orderRepository.Delete(order)
		}
		return "orders deleted"
	}
}

func no(customer m.Customer) bool {
	return customer == nil
}
