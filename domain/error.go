package domain

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("your requested photo is not found")
	ErrConflict            = errors.New("your photo already exist")
	ErrBadRequest          = errors.New("given param is not valid")
)
