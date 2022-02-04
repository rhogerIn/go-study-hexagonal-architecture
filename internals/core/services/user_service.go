package services

import (
	"errors"
	"hexagonArchitecture/internals/core/ports"
)

type UserService struct {
	userRepository ports.IUserRepository
}

// check if the interface has been implemented correctly.
var _ ports.IUserService = (*UserService)(nil)

func NewUserService(repository ports.IUserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) Login(email string, password string) error {
	err := s.userRepository.Login(email, password)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Register(email string, password string, confirmPass string) error {
	if password != confirmPass {
		return errors.New("Password not matched")
	}
	err := s.userRepository.Register(email, password)

	if err != nil {
		return err
	}

	return nil
}