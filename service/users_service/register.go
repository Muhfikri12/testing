package usersservice

import "ecommers/model"

func (u *UsersService) Register(user *model.Users) error {

	if err := u.Repo.UsersRepo.Register(user); err != nil {
		u.Logger.Error("Error from Register Service: " + err.Error())
		return err
	}

	return nil
}
