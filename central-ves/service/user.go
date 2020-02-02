package service

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/control"
	userservice "github.com/Myriad-Dreamin/go-ves/central-ves/service/user"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type UserService = control.UserService

func NewUserService(m module.Module) (UserService, error) {
	return userservice.NewService(m)
}

func (s *Provider) UserService() UserService {
	return s.userService
}
