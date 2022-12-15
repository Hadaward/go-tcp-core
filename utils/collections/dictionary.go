package collections

type Dictionary[T any] map[string]T

func (dictionary Dictionary[T]) Has(key string) bool {
	_, has := dictionary[key]
	return has
}

func (dictionary Dictionary[T]) Remove(key string) bool {
	if _, has := dictionary[key]; has {
		delete(dictionary, key)
		return true
	}
	return false
}

func (dictionary Dictionary[T]) Set(key string, value T) {
	dictionary[key] = value
}

func (dictionary Dictionary[T]) Keys() []string {
	keys := []string{}

	for key := range dictionary {
		keys = append(keys, key)
	}

	return keys
}

func (dictionary Dictionary[T]) Values() []T {
	values := []T{}

	for _, value := range dictionary {
		values = append(values, value)
	}

	return values
}
