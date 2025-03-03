package lazystream

import (
	"cmp"
	"slices"
)

func (s *Stream[T]) All(predicate func(T) bool) bool {
	// all
	for item := range s._seq {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func (s *Stream[T]) Any(predicate func(T) bool) bool {
	// any
	for item := range s._seq {
		if predicate(item) {
			return true
		}
	}
	return false
}

func (s *Stream[T]) Enumerate() *Stream2[int, T] {
	// enumerate
	i := 0
	return &Stream2[int, T]{func(yield func(int, T) bool) {
		for item := range s._seq {
			if !yield(i, item) {
				return
			}
			i++
		}
	}}
}

func (s *Stream[T]) Filter(predicate func(T) bool) *Stream[T] {
	// filter
	return &Stream[T]{func(yield func(T) bool) {
		for item := range s._seq {
			if predicate(item) {
				if !yield(item) {
					return
				}
			}
		}
	}}
}

func (s *Stream[T]) Len() int {
	// len
	len := 0
	for range s._seq {
		len++
	}
	return len
}

func (s *Stream[T]) List() []T {
	// list
	var list []T
	for item := range s._seq {
		list = append(list, item)
	}
	return list
}

func Map[T, R any](s *Stream[T], mapper func(T) R) *Stream[R] {
	// map
	// NOTE: Additional MapFuncs are in map.go
	return &Stream[R]{func(yield func(R) bool) {
		for item := range s._seq {
			if !yield(mapper(item)) {
				return
			}
		}
	}}
}

func (s *Stream[T]) Max(comparator func(T, T) int) T {
	// max
	var max T
	first := true
	for item := range s._seq {
		if first {
			max = item
			first = false
		} else if comparator(item, max) > 0 {
			max = item
		}
	}
	return max
}

func (s *Stream[T]) Min(comparator func(T, T) int) T {
	// min
	var min T
	first := true
	for item := range s._seq {
		if first {
			min = item
			first = false
		} else if comparator(item, min) < 0 {
			min = item
		}
	}
	return min
}

func Range(start, end, step int) *Stream[int] {
	// range
	return &Stream[int]{func(yield func(int) bool) {
		for i := start; i < end; i += step {
			if !yield(i) {
				return
			}
		}
	}}
}

func (s *Stream[T]) Reversed() *Stream[T] {
	// reversed
	list := s.List()
	return &Stream[T]{func(yield func(T) bool) {
		for i := len(list) - 1; i >= 0; i-- {
			if !yield(list[i]) {
				return
			}
		}
	}}
}

func (s *Stream[T]) Slice(start, stop, step int) *Stream[T] {
	// slice
	return &Stream[T]{func(yield func(T) bool) {
		i := 0
		for item := range s._seq {
			if i >= start && i < stop && (i-start)%step == 0 {
				if !yield(item) {
					return
				}
			}
			i++
		}
	}}
}

func (s *Stream[T]) Sum(addFunc func(T, T) T) T {
	// sum
	var sum T
	for item := range s._seq {
		sum = addFunc(sum, item)
	}
	return sum
}

func (s *Stream[T]) Sorted(cmp func(T, T) int) *Stream[T] {
	// sorted
	items := s.List()
	slices.SortFunc(items, cmp)
	return &Stream[T]{func(yield func(T) bool) {
		for _, item := range items {
			if !yield(item) {
				return
			}
		}
	}}
}

func (s *Stream[T]) SortedStable(cmp func(T, T) int) *Stream[T] {
	// sorted
	items := s.List()
	slices.SortStableFunc(items, cmp)
	return &Stream[T]{func(yield func(T) bool) {
		for _, item := range items {
			if !yield(item) {
				return
			}
		}
	}}
}

func Sort[T cmp.Ordered](s *Stream[T]) *Stream[T] {
	// sorted
	items := s.List()
	slices.Sort(items)
	return FromSlice(items)
}

func Zip[K, V any](s1 *Stream[K], s2 *Stream[V]) *Stream2[K, V] {
	// zip
	leftChan := s1.ToChannel()
	rightChan := s2.ToChannel()
	return &Stream2[K, V]{func(yield func(K, V) bool) {
		for {
			left, leftOpen := <-leftChan
			right, rightOpen := <-rightChan
			if !leftOpen || !rightOpen {
				return
			}
			if !yield(left, right) {
				return
			}
		}
	}}
}
