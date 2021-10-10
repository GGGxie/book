package svc

import (
	"book/service/search/cmd/api/internal/config"
	"book/service/search/cmd/api/internal/middleware"
	"github.com/tal-tech/go-zero/rest"
)

type ServiceContext struct {
	Config  config.Config
	Example rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Example: middleware.NewExampleMiddleware().Handle,
	}
}
