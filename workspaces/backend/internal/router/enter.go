package router

import (
	"github.com/go-open-auth/internal/router/admin"
	"github.com/go-open-auth/internal/router/auth"
)

type GroupRouter struct {
	Auth  auth.RouterGroup
	Admin admin.RouterGroup
}

var AppRouter = new(GroupRouter)
