package utils

import (
	"fmt"
	"sort"

	"github.com/iancoleman/orderedmap"
)

type OrderedMap[T any] struct {
	Map *orderedmap.OrderedMap
}

func (o *OrderedMap[T]) Set(key string, item T) {
	o.Map.Set(key, item)
}

func (o *OrderedMap[T]) Get(key string) (T, error) {
	var item T
	val, ok := o.Map.Get(key)
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
	o.Map.SortKeys(sort.Strings)
}

func NewOrderedMap[T any]() *OrderedMap[T] {
	return &OrderedMap[T]{
		Map: orderedmap.New(),
	}
}
