package usersservice

import (
	"ecommers/model"
)

func (u *UsersService) GetListAddress(token string) (*[]model.Addresses, error) {

	Addresses, err := u.Repo.UsersRepo.GetListAddress(token)
	if err != nil {
		u.Logger.Error("Error from GetListAddress service: " + err.Error())
		return nil, err
	}

	return Addresses, nil
}

func (u *UsersService) GetDetailUser(token string) (*model.Users, error) {

	user, err := u.Repo.UsersRepo.GetDetailUser(token)
	if err != nil {
		u.Logger.Error("Error from service GetDetailUser: " + err.Error())
		return nil, err
	}

	return user, nil
}

func (u *UsersService) UpdateUserData(token string, user *model.Users) error {

	err := u.Repo.UsersRepo.UpdateUserData(token, user)
	if err != nil {
		u.Logger.Error("Error from service GetDetailUser: " + err.Error())
		return err
	}

	return nil
}
