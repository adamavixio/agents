package domain

import "errors"

var (
	ErrorInvalid       = errors.New("entity invalid")
	ErrorNotFound      = errors.New("entity not found")
	ErrorAlreadyExists = errors.New("entity already exists")
	ErrorUnimplemented = errors.New("method unimplemented")
)
