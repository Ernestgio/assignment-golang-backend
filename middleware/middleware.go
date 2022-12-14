package middleware

import "assignment-golang-backend/hashutils"

type Middleware struct {
	hashUtils hashutils.HashUtils
}

type MiddlewareConfig struct {
	HashUtil hashutils.HashUtils
}

func NewMiddleware(cfg *MiddlewareConfig) *Middleware {
	return &Middleware{
		hashUtils: cfg.HashUtil,
	}
}
