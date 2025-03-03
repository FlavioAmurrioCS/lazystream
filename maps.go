package lazystream

func (s *Stream[T]) MapT(mapper func(T) T) *Stream[T] {
	return Map(s, mapper)
}

func (s *Stream[T]) MapInt(mapper func(T) int) *Stream[int] {
	return Map(s, mapper)
}

func (s *Stream[T]) MapString(mapper func(T) string) *Stream[string] {
	return Map(s, mapper)
}

func (s *Stream[T]) MapBool(mapper func(T) bool) *Stream[bool] {
	return Map(s, mapper)
}

func (s *Stream[T]) MapFloat32(mapper func(T) float32) *Stream[float32] {
	return Map(s, mapper)
}
func (s *Stream[T]) MapFloat64(mapper func(T) float64) *Stream[float64] {
	return Map(s, mapper)
}

func (s *Stream[T]) MapInt8(mapper func(T) int8) *Stream[int8] {
	return Map(s, mapper)
}

func (s *Stream[T]) MapInt16(mapper func(T) int16) *Stream[int16] {
	return Map(s, mapper)
}

func (s *Stream[T]) MapInt32(mapper func(T) int32) *Stream[int32] {
	return Map(s, mapper)
}

func (s *Stream[T]) MapInt64(mapper func(T) int64) *Stream[int64] {
	return Map(s, mapper)
}

func (s *Stream[T]) MapUint(mapper func(T) uint) *Stream[uint] {
	return Map(s, mapper)
}

func (s *Stream[T]) MapUint8(mapper func(T) uint8) *Stream[uint8] {
	return Map(s, mapper)
}

func (s *Stream[T]) MapUint16(mapper func(T) uint16) *Stream[uint16] {
	return Map(s, mapper)
}

func (s *Stream[T]) MapUint32(mapper func(T) uint32) *Stream[uint32] {
	return Map(s, mapper)
}

func (s *Stream[T]) MapUint64(mapper func(T) uint64) *Stream[uint64] {
	return Map(s, mapper)
}

func (s *Stream[T]) MapUintptr(mapper func(T) uintptr) *Stream[uintptr] {
	return Map(s, mapper)
}

func (s *Stream[T]) MapRune(mapper func(T) rune) *Stream[rune] {
	return Map(s, mapper)
}

func (s *Stream[T]) MapByte(mapper func(T) byte) *Stream[byte] {
	return Map(s, mapper)
}

func (s *Stream[T]) MapComplex64(mapper func(T) complex64) *Stream[complex64] {
	return Map(s, mapper)
}

func (s *Stream[T]) MapComplex128(mapper func(T) complex128) *Stream[complex128] {
	return Map(s, mapper)
}

func (s *Stream[T]) MapError(mapper func(T) error) *Stream[error] {
	return Map(s, mapper)
}
