package middleware

import (
	pkgLog "gitlab.com/tantai-smap/authenticate-api/pkg/log"
	pkgScope "gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

type Middleware struct {
	l          pkgLog.Logger
	jwtManager pkgScope.Manager
}

func New(l pkgLog.Logger, jwtManager pkgScope.Manager) Middleware {
	return Middleware{
		l:          l,
		jwtManager: jwtManager,
	}
}
