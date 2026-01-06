package repositories

import (
	"boilerplate-api/internal/common/context"
	"boilerplate-api/internal/models"

	"gorm.io/gorm"
)

type ProfileRepository struct {
	DB *gorm.DB
}

func NewProfileRepository(ctx *context.AppContext) *ProfileRepository {
	return &ProfileRepository{
		DB: ctx.DB,
	}
}

func (r *ProfileRepository) CreateProfile(userID string, profile *models.Profile) (*models.Profile, error) {
	profile.UserID = userID
	if err := r.DB.Create(profile).Error; err != nil {
		return nil, err
	}
	return profile, nil
}

func (r *ProfileRepository) GetProfileByUserID(userID string) (*models.Profile, error) {
	var profile models.Profile
	if err := r.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *ProfileRepository) UpdateProfile(profile *models.Profile) error {
	return r.DB.Save(profile).Error
}
