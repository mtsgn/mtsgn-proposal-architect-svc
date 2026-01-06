package repositories

import (
	"boilerplate-api/internal/common/context"
	"boilerplate-api/internal/models"

	"gorm.io/gorm"
)

func NewUserRepository(ctx *context.AppContext) *UserRepository {
	return &UserRepository{
		DB: ctx.DB,
	}
}

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	err := r.DB.Create(user).Error
	if err != nil {
		return nil, err
	}

	var createdUser models.User
	err = r.DB.First(
		&createdUser,
		"phone_number = ? AND country_dial_code_id = ?",
		user.PhoneNumber,
		user.CountryDialCodeID,
	).Error
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}

func (r *UserRepository) Update(user *models.User) (*models.User, error) {
	err := r.DB.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByPhoneNumber(dialCode, phoneNumber string) (*models.User, error) {
	var user models.User
	if err := r.DB.
		Select("users.*").
		Preload("Profile").
		Joins("JOIN country_dial_codes ON users.country_dial_code_id = country_dial_codes.id").
		Where("country_dial_codes.dial_code = ? AND users.phone_number = ?", dialCode, phoneNumber).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindById(id string) (*models.User, error) {
	var user models.User
	err := r.DB.Preload("Country").Preload("CountryDialCode").Preload("Profile").First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
