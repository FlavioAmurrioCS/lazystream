package lazystream

func (s *Stream[T]) Append(items ...T) *Stream[T] {
	return &Stream[T]{func(yield func(T) bool) {
		for item := range s._seq {
			if !yield(item) {
				return
			}
		}
		for _, item := range items {
			if !yield(item) {
				return
			}
		}
	}}
}

func (s *Stream[T]) Prepend(items ...T) *Stream[T] {
	return &Stream[T]{func(yield func(T) bool) {
		for _, item := range items {
			if !yield(item) {
				return
			}
		}
		for item := range s._seq {
			if !yield(item) {
				return
			}
		}
	}}
}
