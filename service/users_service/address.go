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

func (u *UsersService) DeleteAddress(token string, id int) error {

	err := u.Repo.UsersRepo.DeleteAddress(token, id)
	if err != nil {
		u.Logger.Error("Error from service DeleteAddress: " + err.Error())
		return err
	}

	return nil
}

func (u *UsersService) SetDefault(token string, id int) error {

	err := u.Repo.UsersRepo.SetDefault(token, id)
	if err != nil {
		u.Logger.Error("Error from service SetDefault: " + err.Error())
		return err
	}

	return nil
}
