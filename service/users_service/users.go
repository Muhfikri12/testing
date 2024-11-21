package usersservice

import (
	"ecommers/model"
)

func (u *UsersService) GetListAddress(token string) (*[]model.Users, error) {

	users, err := u.Repo.UsersRepo.GetListAddress(token)
	if err != nil {
		u.Logger.Error("Error from GetListAddress service: " + err.Error())
		return nil, err
	}

	return users, nil
}
