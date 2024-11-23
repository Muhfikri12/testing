package usersservice

import "ecommers/model"

func (u *UsersService) AddAddress(token string, address *model.Addresses) error {

	err := u.Repo.UsersRepo.AddAddress(token, address)
	if err != nil {
		u.Logger.Error("Error from service AddAddress: " + err.Error())
		return err
	}

	return nil
}

func (u *UsersService) UpdateAddress(token string, address *model.Addresses) error {

	err := u.Repo.UsersRepo.UpdateAddress(token, address)
	if err != nil {
		u.Logger.Error("Error from service UpdateAddress: " + err.Error())
		return err
	}

	return nil
}
