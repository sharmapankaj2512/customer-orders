package shared

import (
	m "customer-orders/domain/model"
)

type CustomerRepository interface{
	Find(customerId int) m.Customer
}