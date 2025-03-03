package lazystream

import "iter"

func Identity[T any](v T) T {
	return v
}

type Pair[T, U any] struct {
	Left  T
	Right U
}

func (pair *Pair[T, U]) Splat() (T, U) {
	return pair.Left, pair.Right
}

type Triplet[A, B, C any] struct {
	Left   A
	Middle B
	Right  C
}

func (triplet *Triplet[A, B, C]) Splat() (A, B, C) {
	return triplet.Left, triplet.Middle, triplet.Right
}

type Stream2[K, V any] struct {
	_seq iter.Seq2[K, V]
}

func ToPairStream[K, V any](s *Stream2[K, V]) *Stream[Pair[K, V]] {
	return &Stream[Pair[K, V]]{func(yield func(Pair[K, V]) bool) {
		for left, right := range s._seq {
			if !yield(Pair[K, V]{Left: left, Right: right}) {
				return
			}
		}
	}}
}
