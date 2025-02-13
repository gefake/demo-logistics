package service

import "logistic_api/pkg/database"

func (s *Service) Login(login, pass string) (bool, error) {
	usr, err := s.UserRepository.GetByLogin(login)
	if err != nil {
		return false, err
	}
	return usr.Password == pass, nil
}

func (s *Service) Register(user database.User) (uint, error) {
	role, err := s.RoleRepository.GetRoleByName("Пользователь")
	if err != nil {
		return 0, err
	}

	user.PositionID = 0
	user.RoleID = role.ID

	return s.UserRepository.CreateUser(user)
}
