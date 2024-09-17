package model

import "errors"

var (
	ErrRequired        = errors.New("required value")
	ErrNotFound        = errors.New("not found")
	ErrNil             = errors.New("nil data")
	ErrNegative        = errors.New("negative value")
	ErrInvalidUserID   = errors.New("invalid user ID")
	ErrInvalidShowIDs  = errors.New("invalid show IDs")
	ErrNoUserInContext = errors.New("no user in context")
)
