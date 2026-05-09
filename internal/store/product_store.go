package store

import (
	"sync"

	"github.com/misiuleczek02/zadanie7-server/internal/model"
)

type ProductStore struct {
	mu       sync.RWMutex
	products []model.Product
}

func NewProductStore() *ProductStore {
	return &ProductStore{
		products: []model.Product{
			{ID: 1, Name: "Klawiatura Mechaniczna", Price: 350.00},
			{ID: 2, Name: "Myszka Bezprzewodowa", Price: 150.00},
			{ID: 3, Name: "Monitor 27 cali", Price: 1200.00},
		},
	}
}

func (s *ProductStore) All() []model.Product {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]model.Product, len(s.products))
	copy(out, s.products)
	return out
}
