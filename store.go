package main

import "sync"

type ReceiptStore struct {
	sync.RWMutex
	receipts map[string]Receipt
}

func NewReceiptStore() *ReceiptStore {
	return &ReceiptStore{
		receipts: make(map[string]Receipt),
	}
}

func (s *ReceiptStore) Add(receipt Receipt) {
	s.Lock()
	defer s.Unlock()
	s.receipts[receipt.ID.String()] = receipt
}

func (s *ReceiptStore) Get(id string) (Receipt, bool) {
	s.RLock()
	defer s.RUnlock()
	receipt, exists := s.receipts[id]
	return receipt, exists
}
