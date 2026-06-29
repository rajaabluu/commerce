package exception

import "errors"

var (
	ErrTooManyProductImages = errors.New("maximum 5 product images allowed")
	ErrNoProductImages      = errors.New("no product images uploaded")
)
