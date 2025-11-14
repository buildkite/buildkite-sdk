package utils

import (
	"fmt"
	"sort"

	"github.com/iancoleman/orderedmap"
)

type OrderedMap[T any] struct {
	data *orderedmap.OrderedMap
}

func (o *OrderedMap[T]) Keys() []string {
	return o.data.Keys()
}

func (o *OrderedMap[T]) Set(key string, item T) {
	o.data.Set(key, item)
}

func (o *OrderedMap[T]) Get(key string) (T, error) {
	var item T
	val, ok := o.data.Get(key)
	if !ok {
		return item, fmt.Errorf("no value in map found for key \"%s\"", key)
	}

	item, ok = val.(T)
	if !ok {
		return item, fmt.Errorf("incorrect value type for key \"%s\"", key)
	}

	return item, nil
}

func (o *OrderedMap[T]) SortKeys() {
	o.data.SortKeys(sort.Strings)
}

func NewOrderedMap[T any](data map[string]T) *OrderedMap[T] {
	newMap := &OrderedMap[T]{
		data: orderedmap.New(),
	}

	for key, item := range data {
		newMap.Set(key, item)
	}
	newMap.SortKeys()

	return newMap
}
