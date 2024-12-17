package handler

import (
	"myChat-API2/internal/service"
)

type Handlers struct {
	AuthService service.IAuthService
	ChatService service.IChatService
}

func NewHandlers(
	as service.IAuthService,
	cs service.IChatService,
) *Handlers {
	return &Handlers{
		AuthService: as,
		ChatService: cs,
	}
}
