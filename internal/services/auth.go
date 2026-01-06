package services

import (
	appContext "boilerplate-api/internal/common/context"
	"boilerplate-api/internal/common/errors"
	"boilerplate-api/internal/common/utils"
	"boilerplate-api/internal/models"
	"boilerplate-api/internal/repositories"
	"boilerplate-api/internal/schemas"
	"boilerplate-api/pkg/easemob"
	"fmt"
)

type AuthService struct {
	CommonRepo    *repositories.CommonRepository
	UserRepo      *repositories.UserRepository
	ProfileRepo   *repositories.ProfileRepository
	Ctx           *appContext.AppContext
	EasemobClient *easemob.Client
}

func NewAuthService(
	CommonRepo *repositories.CommonRepository,
	UserRepo *repositories.UserRepository,
	ProfileRepo *repositories.ProfileRepository,
	Ctx *appContext.AppContext,
) *AuthService {
	return &AuthService{
		CommonRepo:    CommonRepo,
		UserRepo:      UserRepo,
		ProfileRepo:   ProfileRepo,
		Ctx:           Ctx,
		EasemobClient: easemob.NewClient(Ctx.Cfg.Easemob),
	}
}

func (s *AuthService) Register(payload *schemas.RegisterPayload) (*models.User, error) {
	countryDialCode, err := s.CommonRepo.FindCountryDialCodeByDialCode(payload.DialCode)
	if err != nil {
		return nil, errors.ErrUnprocessableEntity
	}
	_, err = s.UserRepo.FindByPhoneNumber(payload.DialCode, payload.PhoneNumber)
	if err == nil {
		return nil, errors.ErrAlreadyExists
	}

	easemobUsername := fmt.Sprintf("%s%s", payload.DialCode, payload.PhoneNumber)
	easemobReq := easemob.EasemobUserRegistrationRequest{
		Username: easemobUsername,
		Password: payload.Password,
	}

	easemobResp, err := s.EasemobClient.RegisterUser(easemobReq)
	if err != nil {
		return nil, fmt.Errorf("failed to register user in Easemob: %w", err)
	}
	easemobUUID := easemobResp.Entities[0].UUID

	user := &models.User{
		CountryDialCodeID: countryDialCode.ID,
		CountryID:         countryDialCode.CountryID,
		PhoneNumber:       payload.PhoneNumber,
		Password:          payload.Password,
		EasemobUsername:   easemobUsername,
		EasemobUUID:       easemobUUID,
		EasemobPassword:   payload.Password,
	}

	createdUser, err := s.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}

	gender, nickname, birthday := utils.GenerateDefaultProfile()

	profile := &models.Profile{
		UserID:   createdUser.ID.String(),
		Nickname: nickname,
		Gender:   *gender,
		Birthday: birthday,
	}

	_, err = s.ProfileRepo.CreateProfile(createdUser.ID.String(), profile)
	if err != nil {
		return nil, fmt.Errorf("failed to create profile: %w", err)
	}

	return createdUser, nil
}

func (s *AuthService) SignInByCredentials(payload *schemas.SignInPayload) (*models.User, error) {
	user, err := s.UserRepo.FindByPhoneNumber(payload.DialCode, payload.PhoneNumber)
	if err != nil {
		return nil, err
	}

	if !user.ComparePassword(payload.Password) {
		return nil, errors.ErrUnauthorized
	}

	return user, nil

}

func (s *AuthService) GeneratePairTokens(userId, nickname, avatar string) (string, string, error) {
	accessToken, err := utils.GenerateAccessToken(userId, nickname, avatar, s.Ctx)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.GenerateRefreshToken(userId, nickname, avatar, s.Ctx)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) RefreshToken(refreshToken string) (string, string, error) {

	claims, err := utils.ValidateToken(refreshToken, s.Ctx)
	if err != nil {
		return "", "", err
	}

	accessToken, err := utils.GenerateAccessToken(claims.UserID, claims.Nickname, claims.Avatar, s.Ctx)
	if err != nil {
		return "", "", err
	}

	refreshTokenNew, err := utils.GenerateRefreshToken(claims.UserID, claims.Nickname, claims.Avatar, s.Ctx)
	if err != nil {
		return "", "", err
	}
	//this is error
	return accessToken, refreshTokenNew, nil
}
