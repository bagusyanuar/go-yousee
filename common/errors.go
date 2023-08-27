package common

import "errors"

var (
	ErrPasswordNotMatch = errors.New("password did not match")
	ErrUserNotFound     = errors.New("user not found")
)
