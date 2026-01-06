package services

import (
	"boilerplate-api/internal/models"
	"boilerplate-api/internal/repositories"
	"boilerplate-api/internal/schemas"
)

type ProfileService struct {
	ProfileRepo *repositories.ProfileRepository
}

func NewProfileService(profileRepo *repositories.ProfileRepository) *ProfileService {
	return &ProfileService{
		ProfileRepo: profileRepo,
	}
}

func (s *ProfileService) GetProfileByUserID(userID string) (*models.Profile, error) {
	return s.ProfileRepo.GetProfileByUserID(userID)
}

func (s *ProfileService) ToSchema(profile *models.Profile) schemas.GetProfileRes {
	if profile == nil {
		return schemas.GetProfileRes{}
	}
	return schemas.GetProfileRes{
		ID:             profile.ID,
		UserID:         profile.UserID,
		Nickname:       profile.Nickname,
		Gender:         profile.Gender,
		Birthday:       profile.Birthday,
		Avatar:         profile.Avatar,
		Voice:          profile.Voice,
		SelfIntro:      profile.SelfIntro,
		LikeCount:      profile.LikeCount,
		FollowingCount: profile.FollowingCount,
		FollowerCount:  profile.FollowerCount,
	}
}
