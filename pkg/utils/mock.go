package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockItem struct {
	Args        []interface{}
	ReturnValue interface{}
}

type Mock struct {
	Items map[string][]mockItem
}

func (m *Mock) AddMockItem(name string, args []interface{}, returnValue interface{}) {
	if _, ok := m.Items[name]; !ok {
		m.Items[name] = []mockItem{}
	}

	m.Items[name] = append(m.Items[name], mockItem{
		Args:        args,
		ReturnValue: returnValue,
	})
}

func (m *Mock) GetFirstMockItem(t *testing.T, name string) mockItem {
	assert.Greater(t, len(m.Items[name]), 0, fmt.Sprintf("no mock invocations available for %s", name))
	item := m.Items[name][0]
	m.Items[name] = m.Items[name][1:]
	return item

}

func NewMock() Mock {
	return Mock{
		Items: make(map[string][]mockItem),
	}
}

func runMock(t *testing.T, mock mockItem, args []interface{}) {
	assert.Equal(t, len(mock.Args), len(args), "length of expected args and provided args are different")

	for i, arg := range args {
		assert.Equal(t, mock.Args[i], arg)
	}
}

func InvokeErrorOnlyMock(t *testing.T, mock mockItem, args []interface{}) error {
	runMock(t, mock, args)
	if mock.ReturnValue == nil {
		return nil
	}

	return mock.ReturnValue.(error)
}

func InvokeMock(t *testing.T, mock mockItem, args []interface{}) interface{} {
	runMock(t, mock, args)
	return mock.ReturnValue
}
