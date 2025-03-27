package middleware

import "server/src/service"

type Middleware func(service service.Service) service.Service
