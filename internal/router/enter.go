package router

import (
	"github.com/open-auth/internal/router/admin"
	"github.com/open-auth/internal/router/auth"
)

type GroupRouter struct {
	Auth  auth.RouterGroup
	Admin admin.RouterGroup
}

var AppRouter = new(GroupRouter)
