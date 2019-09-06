package punkapi

import "errors"

// ErrNoMorePages is returned when a listing call detected an empty request
var ErrNoMorePages = errors.New("No more pages found")
