package service

import "accounting/repository"

type IUserService interface {
	Login()
}

type UserService struct {
	userRepo *repository.UserRepository
}


