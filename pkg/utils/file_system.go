package utils

import (
	"os"
	"testing"
)

type FileSystem interface {
	NewDirectory(name string) error
	NewFile(name, contents string) error
}

type fileSystem struct{}

func (fileSystem) NewDirectory(name string) error {
	return os.Mkdir(name, os.ModePerm)
}

func (fileSystem) NewFile(name, contents string) error {
	return os.WriteFile(name, []byte(contents), os.ModePerm)
}

var FS = fileSystem{}

// Mocks
type mockFileSystem struct {
	T     *testing.T
	Mocks *Mock
}

func (m mockFileSystem) NewDirectory(name string) error {
	mock := m.Mocks.GetFirstMockItem(m.T, "NewDirectory")
	return InvokeErrorOnlyMock(m.T, mock, []interface{}{name})
}

func (m mockFileSystem) NewFile(name, contents string) error {
	mock := m.Mocks.GetFirstMockItem(m.T, "NewFile")
	return InvokeErrorOnlyMock(m.T, mock, []interface{}{name, contents})
}

func NewMockFS(t *testing.T) mockFileSystem {
	return mockFileSystem{
		T: t,
		Mocks: &Mock{
			Items: make(map[string][]mockItem),
		},
	}
}
