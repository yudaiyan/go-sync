package sync

import (
	"iter"
	"log"
	"sync"

	"github.com/go-errors/errors"
)

type SyncMap[K, V any] struct {
	m sync.Map
}

type ISyncMap[K, V any] interface {
	Store(key K, val V)
	Load(key K) (V, bool)
	LoadOrStore(key K, val V) (V, bool)
	Delete(key K)
	LoadAndDelete(key K) (V, bool)
	Iter() iter.Seq2[K, V]
	ContainsKey(k K) bool
}

func (s *SyncMap[K, V]) Store(key K, val V) {
	s.m.Store(key, val)
}

func (s *SyncMap[K, V]) retVal(val any, ok bool) (V, bool) {
	if ok {
		switch val.(type) {
		case nil:
			return *new(V), ok
		default:
			return val.(V), ok
		}
	} else {
		return *new(V), ok
	}
}

func (s *SyncMap[K, V]) Load(key K) (V, bool) {
	val, ok := s.m.Load(key)
	return s.retVal(val, ok)

}

func (s *SyncMap[K, V]) LoadOrStore(key K, val V) (V, bool) {
	actual, loaded := s.m.LoadOrStore(key, val)
	return s.retVal(actual, loaded)
}

func (s *SyncMap[K, V]) Delete(key K) {
	s.m.Delete(key)
}

func (s *SyncMap[K, V]) LoadAndDelete(key K) (V, bool) {
	val, loaded := s.m.LoadAndDelete(key)
	return s.retVal(val, loaded)
}

type Entry[K, V any] struct {
	key K
	val V
}

func (s *SyncMap[K, V]) Iter() iter.Seq2[K, V] {
	var entries []Entry[K, V]
	s.m.Range(func(key, value any) bool {
		entries = append(entries, Entry[K, V]{key.(K), value.(V)})
		return true
	})

	return func(yield func(K, V) bool) {
		for _, entry := range entries {
			if !yield(entry.key, entry.val) {
				return
			}
		}
	}
}

func (s *SyncMap[K, V]) ContainsKey(k K) bool {
	_, ok := s.Load(k)
	return ok
}

func (s *SyncMap[K, V]) ContainsVal(v V) bool {
	log.Println(errors.New("ContainsVal not support").ErrorStack())
	return false
}

type ISyncMapComparableVal[K, V any, C comparable] interface {
	ISyncMap[K, V]
	ContainsVal(v V) bool
}

type SyncMapComparableVal[K, V any, C comparable] struct {
	SyncMap[K, V]
	ToComparable func(V) C
}

func (s *SyncMapComparableVal[K, V, C]) ContainsVal(v V) bool {
	if s.ToComparable == nil {
		log.Println(errors.New("not support, please implement ToComparable").ErrorStack())
		return false
	}

	comparableV := s.ToComparable(v)
	for _, val := range s.Iter() {
		if s.ToComparable(val) == comparableV {
			return true
		}
	}
	return false
}
