package lazystream

import "fmt"

// | `map(func)/select(func)`
// | Maps `func` onto elements of sequence
// | transformation |

// | `starmap(func)/smap(func)`
// | Applies `func` to sequence with `itertools.starmap`
// | transformation |

// | `filter(func)/where(func)`
// | Filters elements of sequence to only those where `func(element)` is `True`
// | transformation |

// | `filter_not(func)`
// | Filters elements of sequence to only those where `func(element)` is `False`
// | transformation |

// | `flatten()`
// | Flattens sequence of lists to a single sequence
// | transformation |

func Flatten[T any](s *Stream[[]T]) *Stream[T] {
	return &Stream[T]{func(yield func(T) bool) {
		for item := range s._seq {
			for _, inner := range item {
				if !yield(inner) {
					return
				}
			}
		}
	}}
}

// | `flat_map(func)`
// | Maps `func` to each element, then merges the result to one flat sequence. `func` must return an iterable
// | transformation |

func FlatMap[T, R any](s *Stream[T], mapper func(T) []R) *Stream[R] {
	return Flatten(Map(s, mapper))
}

// | `group_by(func)`
// | Groups sequence into `(key, value)` pairs where `key=func(element)` and `value` is from the original sequence
// | transformation |

// | `group_by_key()`
// | Groups sequence of `(key, value)` pairs by `key`
// | transformation |

// | `reduce_by_key(func)`
// | Reduces list of `(key, value)` pairs using `func`
// | transformation |

// | `count_by_key()`
// | Counts occurrence of each `key` in sequence of `(key, value)` pairs
// | transformation |

// | `count_by_value()`
// | Counts occurrence of each value in the sequence
// | transformation |

// | `union(other)`
// | Union of unique elements in sequence and `other`
// | transformation |

// | `intersection(other)`
// | Intersection of unique elements in sequence and `other`
// | transformation |

// | `difference(other)`
// | New sequence with unique elements present in sequence but not in `other`
// | transformation |

// | `symmetric_difference(other)`
// | New sequence with unique elements present in sequence or `other`, but not both
// | transformation |

// | `distinct()`
// | Returns distinct elements of sequence. Elements must be hashable
// | transformation |

// | `distinct_by(func)`
// | Returns distinct elements of sequence using `func` as a key
// | transformation |

func DistinctBy[T any, R comparable](s *Stream[T], keyFunc func(T) R) *Stream[T] {
	seen := make(map[R]bool)
	return &Stream[T]{func(yield func(T) bool) {
		for item := range s._seq {
			key := keyFunc(item)
			if !seen[key] {
				seen[key] = true
				if !yield(item) {
					return
				}
			}
		}
	}}
}

func Unique[T any, R comparable](s *Stream[T], keyFunc func(T) R) *Stream[T] {
	return DistinctBy(s, keyFunc)
}

// | `drop(n)`
// | Drops the first `n` elements of the sequence
// | transformation |
func (s *Stream[T]) Drop(n int) *Stream[T] {
	return &Stream[T]{func(yield func(T) bool) {
		i := 0
		for item := range s._seq {
			if i < n {
				i++
				continue
			}
			if !yield(item) {
				return
			}
		}
	}}
}

// | `drop_right(n)`
// | Drops the last `n` elements of the sequence
// | transformation |

// | `drop_while(func)`
// | Drops elements while `func` evaluates to `True`, returning the rest
// | transformation |

// | `take(n)`
// | Returns sequence of first `n` elements
// | transformation |
func (s *Stream[T]) Take(n int) *Stream[T] {
	return &Stream[T]{func(yield func(T) bool) {
		i := 0
		for item := range s._seq {
			if i >= n {
				return
			}
			if !yield(item) {
				return
			}
			i++
		}
	}}
}

// | `take_while(func)`
// | Takes elements while `func` evaluates to `True`, dropping the rest
// | transformation |

// | `init()`
// | Returns sequence without the last element
// | transformation |

// | `tail()`
// | Returns sequence without the first element
// | transformation |

// | `inits()`
// | Returns consecutive inits of sequence
// | transformation |

// | `tails()`
// | Returns consecutive tails of sequence
// | transformation |

// | `zip(other)`
// | Zips the sequence with `other`
// | transformation |

// | `zip_with_index(start=0)`
// | Zips the sequence with the index starting at `start` on the right side
// | transformation |

// | `enumerate(start=0)`
// | Zips the sequence with the index starting at `start` on the left side
// | transformation |

// | `cartesian(*iterables, repeat=1)`
// | Returns cartesian product from itertools.product
// | transformation |

// | `inner_join(other)`
// | Returns inner join of sequence with `other`. Must be a sequence of `(key, value)` pairs
// | transformation |

// | `outer_join(other)`
// | Returns outer join of sequence with `other`. Must be a sequence of `(key, value)` pairs
// | transformation |

// | `left_join(other)`
// | Returns left join of sequence with `other`. Must be a sequence of `(key, value)` pairs
// | transformation |

// | `right_join(other)`
// | Returns right join of sequence with `other`. Must be a sequence of `(key, value)` pairs
// | transformation |

// | `join(other, join_type='inner')`
// | Returns join of sequence with `other` as specified by `join_type`. Must be a sequence of `(key, value)` pairs
// | transformation |

// | `partition(func)`
// | Partitions the sequence into elements that satisfy `func(element)` and those that don't
// | transformation |

// | `grouped(size)`
// | Partitions the elements into groups of size `size`
// | transformation |

// | `sorted(key=None, reverse=False)/order_by(func)`
// | Returns elements sorted according to python `sorted`
// | transformation |

// | `reverse()`
// | Returns the reversed sequence
// | transformation |

// | `slice(start, until)`
// | Sequence starting at `start` and including elements up to `until`
// | transformation |

// | `head(no_wrap=None)` / `first(no_wrap=None)`
// | Returns first element in sequence (if `no_wrap=True`, the result will never be wrapped with `Sequence`)
// | action         |

// | `head_option(no_wrap=None)`
// | Returns first element in sequence or `None` if its empty (if `no_wrap=True`, the result will never be wrapped with `Sequence`)
// | action         |

// | `last(no_wrap=None)`
// | Returns last element in sequence (if `no_wrap=True`, the result will never be wrapped with `Sequence`)
// | action         |

// | `last_option(no_wrap=None)`
// | Returns last element in sequence or `None` if its empty (if `no_wrap=True`, the result will never be wrapped with `Sequence`)
// | action         |

// | `len()` / `size()`
// | Returns length of sequence
// | action         |

// | `count(func)`
// | Returns count of elements in sequence where `func(element)` is True
// | action         |

func (s *Stream[T]) Count(predicate func(T) bool) int {
	count := 0
	for item := range s._seq {
		if predicate(item) {
			count++
		}
	}
	return count
}

// | `empty()`
// | Returns `True` if the sequence has zero length
// | action         |

// | `non_empty()`
// | Returns `True` if sequence has non-zero length
// | action         |

// | `all()`
// | Returns `True` if all elements in sequence are truthy
// | action         |

// | `exists(func)`
// | Returns `True` if `func(element)` for any element in the sequence is `True`
// | action         |

func (s *Stream[T]) Exists(predicate func(T) bool) bool {
	for item := range s._seq {
		if predicate(item) {
			return true
		}
	}
	return false
}

// | `for_all(func)`
// | Returns `True` if `func(element)` is `True` for all elements in the sequence
// | action         |

// | `find(func)`
// | Returns the first element for which `func(element)` evaluates to `True`
// | action         |

func (s *Stream[T]) Find(predicate func(T) bool) (T, bool) {
	for item := range s._seq {
		if predicate(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

// | `any()`
// | Returns `True` if any element in sequence is truthy
// | action         |

// | `max()`
// | Returns maximal element in sequence
// | action         |

// | `min()`
// | Returns minimal element in sequence
// | action         |

// | `max_by(func)`
// | Returns element with maximal value `func(element)`
// | action         |

// | `min_by(func)`
// | Returns element with minimal value `func(element)`
// | action         |

// | `sum()/sum(projection)`
// | Returns the sum of elements possibly using a projection
// | action         |

// | `product()/product(projection)`
// | Returns the product of elements possibly using a projection
// | action         |

// | `average()/average(projection)`
// | Returns the average of elements possibly using a projection
// | action         |

// | `aggregate(func)/aggregate(seed, func)/aggregate(seed, func, result_map)`
// | Aggregates using `func` starting with `seed` or first element of list then applies `result_map` to the result
// | action         |

// | `fold_left(zero_value, func)`
// | Reduces element from left to right using `func` and initial value `zero_value`
// | action         |

// | `fold_right(zero_value, func)`
// | Reduces element from right to left using `func` and initial value `zero_value`
// | action         |

// | `make_string(separator)`
// | Returns string with `separator` between each `str(element)`
// | action         |
func (s *Stream[T]) MakeString(separator string) string {
	var result string
	for item := range s._seq {
		if result != "" {
			result += separator
		}
		result += fmt.Sprint(item)
	}
	return result
}
func (s *Stream[T]) Join(separator string) string {
	return s.MakeString(separator)
}

// | `dict(default=None)` / `to_dict(default=None)`
// | Converts a sequence of `(Key, Value)` pairs to a `dictionary`. If `default` is not None, it must be a value or zero argument callable which will be used to create a `collections.defaultdict`
// | action         |

// | `list()` / `to_list()`
// | Converts sequence to a list
// | action         |

// | `set() / to_set()`
// | Converts sequence to a set
// | action         |

// | `to_file(path)`
// | Saves the sequence to a file at `path` with each element on a newline
// | action         |

// | `to_csv(path)`
// | Saves the sequence to a csv file at `path` with each element representing a row
// | action         |

// | `to_jsonl(path)`
// | Saves the sequence to a jsonl file with each element being transformed to json and printed to a new line
// | action         |

// | `to_json(path)`
// | Saves the sequence to a json file. The contents depend on if the json root is an array or dictionary
// | action         |

// | `to_sqlite3(conn, tablename_or_query, *args, **kwargs)`
// | Saves the sequence to a SQLite3 db. The target table must be created in advance
// | action         |

// | `to_pandas(columns=None)`
// | Converts the sequence to a pandas DataFrame
// | action         |

// | `cache()`
// | Forces evaluation of sequence immediately and caches the result
// | action         |
func (s *Stream[T]) Cache() *Stream[T] {
	cache := s.List()
	return FromSlice(cache)
}

// | `for_each(func)`
// | Executes `func` on each element of the sequence
// | action         |
func (s *Stream[T]) ForEach(action func(T)) {
	for item := range s._seq {
		action(item)
	}
}

// | `peek(func)`
// | Executes `func` on each element of the sequence and returns it
// | transformation |****
func (s *Stream[T]) Peek(action func(T)) *Stream[T] {
	return &Stream[T]{func(yield func(T) bool) {
		for item := range s._seq {
			action(item)
			if !yield(item) {
				return
			}
		}
	}}
}
