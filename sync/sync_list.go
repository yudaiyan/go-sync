package sync

import (
	"fmt"
	"iter"
	"sync"
)

type SyncList[T comparable] struct {
	items []T
	mu    sync.RWMutex
}

type ISyncList[T comparable] interface {
	Add(item T)
	Get(i int) T
	Remove(i int)
	RemoveByVal(val T)
	Find(val T) (int, bool)
	Size() int
	Iter() iter.Seq2[int, T]
}

func (s *SyncList[T]) Add(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, item)
}

func (s *SyncList[T]) Get(i int) T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.items[i]
}

func (s *SyncList[T]) Remove(i int) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.items = append(s.items[:i], s.items[i+1:]...)
}

func (s *SyncList[T]) RemoveByVal(val T) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if i, ok := s.Find(val); ok {
		s.items = append(s.items[:i], s.items[i+1:]...)
	}
}

func (s *SyncList[T]) Find(val T) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for i, item := range s.items {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func (s *SyncList[T]) Size() int {
	return len(s.items)
}

func (s *SyncList[T]) String() string {
	return fmt.Sprintf("%v", s.items)
}

func (s *SyncList[T]) Iter() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, item := range s.items {
			if !yield(i, item) {
				return
			}
		}
	}
}

func NewInstance[T any]() T {
	return *new(T)
}
