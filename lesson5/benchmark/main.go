// Протестируйте производительность операций чтения и записи на множестве действительных чисел,
// безопасность которого обеспечивается sync.Mutex и sync.RWMutex для разных вариантов
// использования: 10% запись, 90% чтение; 50% запись, 50% чтение; 90% запись, 10% чтение
//
package main

import "sync"

type Set struct {
sync.Mutex
mm map[int]int
}

func NewSet() *Set {
return &Set{
mm: map[int]int{},
}
}

func (s *Set) Add(i int) {
s.Lock()
defer s.Unlock()
s.mm[i] = i
}

func (s *Set) Has(i int) bool {
s.Lock()
defer s.Unlock()
_, ok := s.mm[i]
return ok
}

type SetRW struct {
sync.RWMutex
mm map[int]int
}

func NewSetRW() *SetRW {
return &SetRW{
mm: map[int]int{},
}
}

func (s *SetRW) AddRW(i int) {
s.Lock()
defer s.Unlock()
s.mm[i] = i
}

func (s *SetRW) HasRW(i int) bool {
s.RLock()
defer s.RUnlock()
_, ok := s.mm[i]
return ok
}