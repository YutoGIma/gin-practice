package usecase

import (
	"myapp/app/errors"
	"myapp/app/model"
	"myapp/app/service"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userService service.UserService
}

func NewUserUseCase(userService service.UserService) UserUseCase {
	return UserUseCase{
		userService: userService,
	}
}

func (uc *UserUseCase) Create(input model.User) (*model.User, error) {
	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.NewInternalError("Failed to hash password", err)
	}
	input.Password = string(hashedPassword)

	if err := uc.userService.CreateUser(input); err != nil {
		return nil, errors.NewInternalError("Failed to create user", err)
	}
	return &input, nil
}

func (uc *UserUseCase) Update(id uint, input model.User) (*model.User, error) {
	user, err := uc.userService.GetUserDetail(id)
	if err != nil {
		return nil, errors.NewNotFoundError("User not found")
	}

	// パスワードが変更されている場合のみハッシュ化
	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.NewInternalError("Failed to hash password", err)
		}
		input.Password = string(hashedPassword)
	}

	if err := uc.userService.UpdateUser(input); err != nil {
		return nil, errors.NewInternalError("Failed to update user", err)
	}

	return &user, nil
}

func (uc *UserUseCase) Delete(id uint) error {
	if err := uc.userService.DeleteUser(id); err != nil {
		return errors.NewInternalError("Failed to delete user", err)
	}
	return nil
}

func (uc *UserUseCase) GetByID(id uint) (*model.User, error) {
	user, err := uc.userService.GetUserDetail(id)
	if err != nil {
		return nil, errors.NewNotFoundError("User not found")
	}
	return &user, nil
}

func (uc *UserUseCase) List() ([]model.User, error) {
	users, err := uc.userService.GetUsers()
	if err != nil {
		return nil, errors.NewInternalError("Failed to list users", err)
	}
	return users, nil
}

// func (uc *UserUseCase) Authenticate(email, password string) (*model.User, error) {
// 	user, err := uc.GetByEmail(email)
// 	if err != nil {
// 		return nil, errors.NewValidationError("Invalid email or password")
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
// 		return nil, errors.NewValidationError("Invalid email or password")
// 	}

// 	return user, nil
// }
