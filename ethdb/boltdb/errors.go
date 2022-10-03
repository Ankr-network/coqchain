package boltdb

import "errors"

var (
	ErrNotFound = errors.New("Not Found")
	ErrEOF      = errors.New("End of file")
)
