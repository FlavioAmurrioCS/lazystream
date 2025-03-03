package lazystream

import "fmt"

///////////////////////////////////////////////////////////////////////////////
// region: Infinite iterators:
///////////////////////////////////////////////////////////////////////////////

func Count(start int, step int) *Stream[int] {
	// count(start=0, step=1) --> start, start+step, start+2*step, ...
	return &Stream[int]{func(yield func(int) bool) {
		for i := start; ; i += step {
			if !yield(i) {
				return
			}
		}
	}}
}

func Cycle[T any](data []T) *Stream[T] {
	// cycle(p) --> p0, p1, ... plast, p0, p1, ...
	return &Stream[T]{func(yield func(T) bool) {
		for {
			for _, v := range data {
				if !yield(v) {
					return
				}
			}
		}
	}}
}

func (s *Stream[T]) Cycle() *Stream[T] {
	// cycle(p) --> p0, p1, ... plast, p0, p1, ...
	return &Stream[T]{func(yield func(T) bool) {
		for {
			for item := range s._seq {
				if !yield(item) {
					return
				}
			}
		}
	}}
}

func Repeat[T any](elem T, n int) *Stream[T] {
	// repeat(elem [,n]) --> elem, elem, elem, ... endlessly or up to n times
	return &Stream[T]{func(yield func(T) bool) {
		for i := 0; i < n; i++ {
			if !yield(elem) {
				return
			}
		}
	}}
}

///////////////////////////////////////////////////////////////////////////////
// endregion: Infinite iterators:
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// region: Iterators terminating on the shortest input sequence:
///////////////////////////////////////////////////////////////////////////////

func (s *Stream[T]) Accumulate(fn func(T, T) T) *Stream[T] {
	// accumulate(p[, func]) --> p0, p0+p1, p0+p1+p2
	return &Stream[T]{func(yield func(T) bool) {
		var result T
		for item := range s._seq {
			result = fn(result, item)
			if !yield(result) {
				return
			}
		}
	}}
}

func Batched[T any](s *Stream[T], n int) *Stream[[]T] {
	// batched(p, n) --> [p0, p1, ..., p_n-1], [p_n, p_n+1, ..., p_2n-1], ...
	return &Stream[[]T]{func(yield func([]T) bool) {
		var batch []T
		for item := range s._seq {
			batch = append(batch, item)
			if len(batch) == n {
				if !yield(batch) {
					return
				}
				batch = nil
			}
		}
		if len(batch) > 0 {
			yield(batch)
		}
	}}
}

func (s *Stream[T]) Chain(args ...*Stream[T]) *Stream[T] {
	// chain(p, q, ...) --> p0, p1, ... plast, q0, q1, ...
	// chain.from_iterable([p, q, ...]) --> p0, p1, ... plast, q0, q1, ...
	return &Stream[T]{func(yield func(T) bool) {
		for item := range s._seq {
			if !yield(item) {
				return
			}
		}
		for _, arg := range args {
			for item := range arg._seq {
				if !yield(item) {
					return
				}
			}
		}
	}}
}

func (s *Stream[T]) Compress(d *Stream[bool]) *Stream[T] {
	// compress(data, selectors) --> (d[0] if s[0]), (d[1] if s[1]), ...
	return &Stream[T]{func(yield func(T) bool) {
		for left, right := range Zip(s, d)._seq {
			if right {
				if !yield(left) {
					return
				}
			}
		}
	}}
}

func (s *Stream[T]) DropWhile(predicate func(T) bool) *Stream[T] {
	// dropwhile(pred, seq) --> seq[n], seq[n+1], starting when pred fails
	return &Stream[T]{func(yield func(T) bool) {
		skip := true
		for item := range s._seq {
			if skip && predicate(item) {
				continue
			}
			skip = false
			if !yield(item) {
				return
			}
		}
	}}
}

func GroupBy[T any](s *Stream[T], keyFunc func(T) T) *Stream[[]T] {
	// groupby(iterable[, keyfunc]) --> sub-iterators grouped by value of keyfunc(v)
	return &Stream[[]T]{func(yield func([]T) bool) {
		var current string
		acc := []T{}
		if keyFunc == nil {
			keyFunc = func(x T) T { return x }
		}
		for item := range s._seq {
			key := fmt.Sprintf("%#v", keyFunc(item))
			if key != current {
				if len(acc) > 0 {
					yield(acc)
				}
				acc = []T{item}
				current = key
			} else {
				acc = append(acc, item)
			}
		}
	}}
}

func (s *Stream[T]) FilterFalse(predicate func(T) bool) *Stream[T] {
	// filterfalse(pred, seq) --> elements of seq where pred(elem) is False
	return &Stream[T]{func(yield func(T) bool) {
		for item := range s._seq {
			if !predicate(item) {
				if !yield(item) {
					return
				}
			}
		}
	}}
}

func (s *Stream[T]) ISlice(start int, stop int) *Stream[T] {
	// islice(seq, [start,] stop [, step]) --> elements from seq[start:stop:step]
	return &Stream[T]{func(yield func(T) bool) {
		i := 0
		for item := range s._seq {
			if i >= start {
				if !yield(item) {
					return
				}
			}
			i++
			if i >= stop {
				return
			}
		}
	}}
}

func PairWise[T any](s *Stream[T]) *Stream[Pair[T, T]] {
	// pairwise(s) --> (s[0],s[1]), (s[1],s[2]), (s[2], s[3]), ...
	return &Stream[Pair[T, T]]{func(yield func(Pair[T, T]) bool) {
		for item := range Batched(s, 2)._seq {
			if len(item) == 2 {
				if !yield(Pair[T, T]{Left: item[0], Right: item[1]}) {
					return
				}
			}
			var zero T
			if !yield(Pair[T, T]{Left: item[0], Right: zero}) {
				return
			}
		}
	}}
}

// starmap(fun, seq) --> fun(*seq[0]), fun(*seq[1]), ...
// tee(it, n=2) --> (it1, it2 , ... itn) splits one iterator into n

// TakeWhile returns a new Stream that yields items from the original Stream
// as long as the provided predicate function returns true. Once the predicate
// function returns false, the Stream stops yielding items.
//
// The predicate function takes an item of type T and returns a boolean indicating
// whether the item should be included in the resulting Stream.
//
// Example usage:
//
//	stream := NewStream([]int{1, 2, 3, 4, 5})
//	result := stream.TakeWhile(func(n int) bool { return n < 4 })
//	// result will yield 1, 2, 3
//
// Parameters:
//
//	predicate - A function that takes an item of type T and returns a boolean.
//
// Returns:
//
//	A new Stream that yields items from the original Stream as long as the
//	predicate function returns true.
func (s *Stream[T]) TakeWhile(predicate func(T) bool) *Stream[T] {
	// takewhile(pred, seq) --> seq[0], seq[1], until pred fails
	return &Stream[T]{func(yield func(T) bool) {
		for item := range s._seq {
			if !predicate(item) {
				return
			}
			if !yield(item) {
				return
			}
		}
	}}
}

// ZipLongest returns a new Stream where each element is a slice containing
// elements from the original stream and the provided 'other' stream. If the
// streams are of unequal length, the shorter stream will be padded with
// 'fillValue' until both streams are exhausted.
func ZipLongest[T any](s *Stream[T], other *Stream[T], fillValue T) *Stream[[]T] {
	// zip_longest(p, q, ...) --> (p[0], q[0]), (p[1], q[1]), ...
	return &Stream[[]T]{func(yield func([]T) bool) {
		mainChan := s.ToChannel()
		otherChan := other.ToChannel()
		for {
			mainItem, mainOk := <-mainChan
			otherItem, otherOk := <-otherChan
			if !mainOk && !otherOk {
				return
			}
			if !mainOk {
				mainItem = fillValue
			}
			if !otherOk {
				otherItem = fillValue
			}
			if !yield([]T{mainItem, otherItem}) {
				return
			}
		}
	}}
}

///////////////////////////////////////////////////////////////////////////////
// endregion: Iterators terminating on the shortest input sequence:
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// region: Combinatoric generators:
///////////////////////////////////////////////////////////////////////////////
// product(p, q, ... [repeat=1]) --> cartesian product
// permutations(p[, r])
// combinations(p, r)
// combinations_with_replacement(p, r)
///////////////////////////////////////////////////////////////////////////////
// endregion: Combinatoric generators:
///////////////////////////////////////////////////////////////////////////////
