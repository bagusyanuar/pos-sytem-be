package handler

import "github.com/bagusyanuar/pos-sytem-be/config"

type Handlers struct {
	Welcome WelcomHandler
}

func NewHandlers(config *config.Config) *Handlers {
	return &Handlers{
		Welcome: NewWelcomeHandler(config),
	}
}
