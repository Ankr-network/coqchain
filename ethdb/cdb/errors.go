package cdb

import "errors"

var (
	ErrNotFound = errors.New("No matching key/data pair found")
)
