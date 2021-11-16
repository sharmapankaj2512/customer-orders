package shared

import (
	m "customer-orders/domain/model"
)

type OrderRepository interface {
	Find(customerId int) []m.Order
	Delete(orders []m.Order)
}
