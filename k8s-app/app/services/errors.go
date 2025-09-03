package services

import "errors"

// Определяем ошибки сервиса
var (
	ErrAppNotReady    = errors.New("application not ready")
	ErrInvalidRequest = errors.New("invalid request")
	ErrInternalError  = errors.New("internal server error")
)
