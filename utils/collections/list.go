package collections

import (
	"sort"
)

type List[T comparable] []T

func (list List[T]) Add(value T) List[T] {
	return append(list, value)
}

func (list List[T]) Includes(value T) bool {
	return list.IndexOf(value) != -1
}

func (list List[T]) IndexOf(value T) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}

	return -1
}

func (list List[T]) RemoveAt(index int) List[T] {
	if (index < 0) || (index >= len(list)) {
		return list
	}

	return append(list[:index], list[index+1:]...)
}

func (list List[T]) Remove(value T) List[T] {
	index := list.IndexOf(value)

	if index < 0 {
		return list
	}

	return list.RemoveAt(index)
}

func (list List[T]) Concat(arrays ...[]T) List[T] {
	for _, array := range arrays {
		list = append(list, array...)
	}

	return list
}

func (list List[T]) Filter(predicate func(T) bool) List[T] {
	var result List[T]

	for _, v := range list {
		if predicate(v) {
			result = append(result, v)
		}
	}

	return result
}

func (list List[T]) Map(mapper func(T, int) T) List[T] {
	var result List[T]

	for k, v := range list {
		result = append(result, mapper(v, k))
	}

	return result
}

func (list List[T]) Pop() ([]T, T) {
	return list[:len(list)-1], list[len(list)-1]
}

func (list List[T]) Shift() ([]T, T) {
	return list[1:], list[0]
}

func (list List[T]) Some(predicate func(T) bool) bool {
	for _, v := range list {
		if predicate(v) {
			return true
		}
	}

	return false
}

func (list List[T]) Every(predicate func(T) bool) bool {
	for _, v := range list {
		if !predicate(v) {
			return false
		}
	}

	return true
}

func (list List[T]) Sort(sorter func(T, T) bool) {
	sort.SliceStable(list, func(i, j int) bool {
		return sorter(list[i], list[j])
	})
}

func (list List[T]) Keys() List[int] {
	keys := List[int]{}

	for key := range list {
		keys = append(keys, key)
	}

	return keys
}

func (list List[T]) Values() List[T] {
	return append(List[T]{}, list...)
}
