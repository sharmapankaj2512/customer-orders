package shared

import (
	m "customer-orders/domain/model"
)

func CustomerPipe(
	supplier func() (m.Customer, error),
	consumer func(m.Customer) interface{},
) func() (interface{}, error) {
	return func() (interface{}, error) {
		customer, err := supplier()
		if err != nil {
			return err, nil
		}
		return consumer(customer), nil
	}
}

func IoPipe(
	supplier func() (interface{}, error),
	consumer func(interface{})) {
	if output, err := supplier(); err != nil {
		consumer(err)
	} else {
		consumer(output)
	}
}
