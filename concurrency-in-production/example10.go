package main

import (
	"context"
	"time"
)

// START OMIT
func (s *OrdersStore) UpdateWithCollectionByCustomerData(ctx context.Context, key string, data httpservice.FulfilmentCollectionByCustomer) error {
	order, err := s.Read(ctx, key)
	if err != nil {
		return err
	}

	order.FulfilmentMethod = "collection-by-customer"
	order.DateModified = time.Now().Format(time.RFC3339)
	order.FulfilmentCollectionByCustomer = data

	return s.Store(ctx, order)
}

// END OMIT
