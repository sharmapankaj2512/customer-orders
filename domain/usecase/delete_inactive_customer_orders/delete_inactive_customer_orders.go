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
		if output, err := customerPipe(
			findCustomer(customerRepository, customerId),
			deleteOrder(orderRepository)); err != nil {
			writer.Write(err)
		} else {
			writer.Write(output)
		}
	}
}

func customerPipe(
	supplier func() (m.Customer, error),
	consumer func(m.Customer) interface{}) (interface{}, error) {
	customer, err := supplier()
	if err != nil {
		return err, nil
	}
	return consumer(customer), nil
}

func findCustomer(
	customerRepository s.CustomerRepository,
	customerId int) func() (m.Customer, error) {
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

func deleteOrder(orderRepository s.OrderRepository) func(customer m.Customer) interface{} {
	return func(customer m.Customer) interface{} {
		orders := orderRepository.Find(customer.ID())
		orderRepository.Delete(orders)
		return "orders deleted"
	}
}

func no(customer m.Customer) bool {
	return customer == nil
}
