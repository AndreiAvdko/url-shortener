package storage

import "errors"

var (
	ErrURLNotFound = errors.New("url not fuond")
	ErrURLExists   = errors.New("url exists")
)
