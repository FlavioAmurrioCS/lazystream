package lazystream

func (s *Stream[T]) ToChannel() chan T {
	ch := make(chan T)
	go func() {
		for item := range s._seq {
			ch <- item
		}
		close(ch)
	}()
	return ch
}
