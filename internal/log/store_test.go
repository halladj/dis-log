package log

import (
	"os"
	"testing"
)

var write = []byte("hello world")
var width = uint64(len(write)) + lenWidth

// TODO:4
func TestStoreAppendRead(t *testing.T) {

}

// TODO:1
func TestAppend(t *testing.T, s *store)

// TODO:2
func TestRead(t *testing.T, s *store)

// TODO:3
func TestReadAt(t *testing.T, s *store)

func TestStoreClose(t *testing.T)

func openFile(name string) (
	file *os.File,
	size int64,
	err error,
) {

	return nil, 0, nil
}
