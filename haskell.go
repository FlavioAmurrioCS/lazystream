package lazystream

func (s *Stream[T]) Uncons() (T, *Stream[T]) {
	ch := s.ToChannel()
	item, ok := <-ch
	if !ok {
		panic("Stream is empty")
	}
	return item, FromChannel(ch)
}

func (s *Stream[T]) Head() T {
	// all
	x, _ := s.Uncons()
	return x
}

func (s *Stream[T]) Tail() *Stream[T] {
	// all
	_, xs := s.Uncons()
	return xs
}

func (s *Stream[T]) Unsnoc() (*Stream[T], T) {
	// all
	list := s.List()
	if len(list) == 0 {
		panic("Stream is empty")
	}
	return FromSlice(list[:len(list)-1]), list[len(list)-1]
}

func (s *Stream[T]) Init() *Stream[T] {
	// all
	xs, _ := s.Unsnoc()
	return xs
}

func (s *Stream[T]) Last() T {
	// all
	_, x := s.Unsnoc()
	return x
}
