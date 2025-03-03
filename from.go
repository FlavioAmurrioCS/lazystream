package lazystream

import (
	"bufio"
	"iter"
	"os"
)

type Stream[T any] struct {
	_seq iter.Seq[T]
}

func FromSlice[T any](data []T) *Stream[T] {
	return &Stream[T]{func(yield func(T) bool) {
		for _, v := range data {
			if !yield(v) {
				return
			}
		}
	}}
}

func FromChannel[T any](ch <-chan T) *Stream[T] {
	return &Stream[T]{func(yield func(T) bool) {
		for val := range ch {
			if !yield(val) {
				return
			}
		}
	}}
}

func FromSeq[T any](seq iter.Seq[T]) *Stream[T] {
	return &Stream[T]{seq}
}

func FromStdin() *Stream[string] {
	return &Stream[string]{func(yield func(string) bool) {
		info, _ := os.Stdin.Stat()
		if (info.Mode() & os.ModeCharDevice) == 0 {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				if !yield(scanner.Text()) {
					return
				}
			}
		}
	}}
}

func FromFile(path string) *Stream[string] {
	return &Stream[string]{func(yield func(string) bool) {
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				return
			}
		}
	}}
}
