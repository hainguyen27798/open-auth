package router

import (
	"github.com/go-open-auth/internal/router/auth"
)

type GroupRouter struct {
	Auth auth.RouterGroup
}

var AppRouter = new(GroupRouter)
