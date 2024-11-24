package authservice

import "ecommers/model"

func (u *AuthService) Register(user *model.Register) error {

	if err := u.Repo.AuthRepo.Register(user); err != nil {
		u.Logger.Error("Error from Register Service: " + err.Error())
		return err
	}

	return nil
}
