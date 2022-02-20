package domain

import "errors"

var (
	ErrUserNotFound  = errors.New("Incorrent login or password")
	ErrGenerateToken = errors.New("Could not login")
	ErrReplayUsername = errors.New("Such username is already in use")
)
