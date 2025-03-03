package lazystream

// NOTE: Fix when MethodGenerics are supported
func Reduce[T, R any](s *Stream[T], reducer func(R, T) R, initial R) R {
	result := initial
	for item := range s._seq {
		result = reducer(result, item)
	}
	return result
}

func (s *Stream[T]) Reduce(reducer func(T, T) T, initial T) T {
	return Reduce(s, reducer, initial)
}

func Chunked[T any](s *Stream[T], predicate func(T) bool) *Stream[[]T] {
	return &Stream[[]T]{func(yield func([]T) bool) {
		var chunk []T
		for item := range s._seq {
			if predicate(item) && len(chunk) > 0 {
				if !yield(chunk) {
					return
				}
				chunk = nil
			}
			chunk = append(chunk, item)
		}
		if len(chunk) > 0 {
			yield(chunk)
		}
	}}
}
