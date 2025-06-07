package usecase

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"myapp/app/errors"
	"myapp/app/model"
)

type UserUseCase struct {
	db *gorm.DB
}

func NewUserUseCase(db *gorm.DB) *UserUseCase {
	return &UserUseCase{
		db: db,
	}
}

func (uc *UserUseCase) Create(input model.User) (*model.User, error) {
	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.NewInternalError("Failed to hash password", err)
	}
	input.Password = string(hashedPassword)

	if err := uc.db.Create(&input).Error; err != nil {
		return nil, errors.NewInternalError("Failed to create user", err)
	}
	return &input, nil
}

func (uc *UserUseCase) Update(id uint, input model.User) (*model.User, error) {
	var user model.User
	if err := uc.db.First(&user, id).Error; err != nil {
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

	if err := uc.db.Model(&user).Updates(input).Error; err != nil {
		return nil, errors.NewInternalError("Failed to update user", err)
	}

	return &user, nil
}

func (uc *UserUseCase) Delete(id uint) error {
	if err := uc.db.Delete(&model.User{}, id).Error; err != nil {
		return errors.NewInternalError("Failed to delete user", err)
	}
	return nil
}

func (uc *UserUseCase) GetByID(id uint) (*model.User, error) {
	var user model.User
	if err := uc.db.First(&user, id).Error; err != nil {
		return nil, errors.NewNotFoundError("User not found")
	}
	return &user, nil
}

func (uc *UserUseCase) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := uc.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.NewNotFoundError("User not found")
	}
	return &user, nil
}

func (uc *UserUseCase) List() ([]model.User, error) {
	var users []model.User
	if err := uc.db.Find(&users).Error; err != nil {
		return nil, errors.NewInternalError("Failed to list users", err)
	}
	return users, nil
}

func (uc *UserUseCase) Authenticate(email, password string) (*model.User, error) {
	user, err := uc.GetByEmail(email)
	if err != nil {
		return nil, errors.NewValidationError("Invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.NewValidationError("Invalid email or password")
	}

	return user, nil
}
