package handlers

import (
	"boilerplate-api/internal/common/utils"
	"net/http"

	"github.com/gin-gonic/gin"

	"boilerplate-api/internal/common/context"
	e "boilerplate-api/internal/common/errors"
	"boilerplate-api/internal/schemas"
	"boilerplate-api/internal/services"
	"boilerplate-api/pkg/logger"
	"errors"

	"gorm.io/gorm"
)

type AuthHandler struct {
	service        *services.AuthService
	profileService *services.ProfileService
	userService    services.User
	appCtx         *context.AppContext
}

func NewAuthHandler(
	service *services.AuthService,
	profileService *services.ProfileService,
	userService services.User,
	appCtx *context.AppContext,
) *AuthHandler {
	return &AuthHandler{
		service:        service,
		profileService: profileService,
		userService:    userService,
		appCtx:         appCtx,
	}
}

// @Summary     Register new user
// @Description Register a new user with phone number and password. Also registers the user with Easemob chat service.
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       request body schemas.RegisterPayload true "Register credentials"
// @Success     200 {object} schemas.RegisterResponse
// @Failure     400 {object} schemas.Response
// @Failure     409 {object} schemas.Response
// @Failure     422 {object} schemas.Response
// @Router      /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	logs := logger.GetLogger(c)

	var payload schemas.RegisterPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logs.Error().Err(err).Msg("bad request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload.PhoneNumber = utils.FormatPhoneNumber(payload.PhoneNumber)
	payload.DialCode = utils.FormatDialCode(payload.DialCode)

	user, err := h.service.Register(&payload)
	if err != nil {
		if errors.Is(err, e.ErrAlreadyExists) {
			logs.Warn().Err(err).Msg("user is already registered")
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		logs.Error().Err(err).Msg("failed to create user")
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	logs.Info().Str("user_id", user.ID.String()).Msg("user registered")
	c.JSON(http.StatusOK, schemas.RegisterResponse{
		ID:              user.ID,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		LastLoginAt:     user.LastLoginAt,
		PhoneNumber:     user.PhoneNumber,
		CountryDialCode: user.CountryDialCode.ToSchema(),
		Country:         user.Country.ToSchema(),
		EasemobUsername: user.EasemobUsername,
		EasemobUUID:     user.EasemobUUID,
	})
}

// @Summary     Sign in user
// @Description Authenticate a user and return JWT tokens
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       request body schemas.SignInPayload true "Login credentials"
// @Success     200 {object} schemas.PairTokensResponse
// @Failure     400 {object} schemas.Response
// @Failure     401 {object} schemas.Response
// @Failure     422 {object} schemas.Response
// @Router      /auth/sign_in [post]
func (h *AuthHandler) SignIn(c *gin.Context) {
	logs := logger.GetLogger(c)

	var payload schemas.SignInPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logs.Error().Err(err).Msg("bad request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload.PhoneNumber = utils.FormatPhoneNumber(payload.PhoneNumber)
	payload.DialCode = utils.FormatDialCode(payload.DialCode)

	user, err := h.service.SignInByCredentials(&payload)
	if err != nil {
		logs.Error().Err(err).Str("phone_number", payload.PhoneNumber).Msg("failed to sign in user")
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		logs.Error().Str("phone_number", payload.PhoneNumber).Msg("invalid credentials")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return

	}

	accessToken, refreshToken, err := h.service.GeneratePairTokens(user.ID.String(), user.Profile.Nickname, user.Profile.Avatar)
	if err != nil {
		logs.Error().Err(err).Msg("failed to generate tokens")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logs.Info().Str("phone_number", payload.PhoneNumber).Msg("user signed in")
	c.JSON(http.StatusOK, schemas.PairTokensResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

// @Summary     Get user profile
// @Description Get user profile by user ID
// @Tags        auth
// @Accept      json
// @Produce     json
// @Success     200 {object} schemas.GetMe
// @Failure     401 {object} schemas.Response
// @Failure     404 {object} schemas.Response
// @Router      /auth/me [get]
func (h *AuthHandler) GetMe(c *gin.Context) {
	logs := logger.GetLogger(c)

	userCtx := utils.GetUserContext(c)
	if userCtx == nil {
		logs.Error().Msg("user_context not found, shouldn't happened")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	//user, err := h.service.UserRepo.FindById(userCtx.UserID)
	user, err := h.userService.FindById(c.Request.Context(), userCtx.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error().Err(err).Msg("user not found. shouldn't happen")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		logs.Error().Err(err).Msg("failed to find user")
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	profile, err := h.profileService.GetProfileByUserID(userCtx.UserID)
	var profileSchema schemas.GetProfileRes
	if err != nil {
		//log.Println("Fail get profile of user")
		logs.Warn().Err(err).Msg("failed to get user's profile")
		profileSchema = schemas.GetProfileRes{}
	} else {
		profileSchema = h.profileService.ToSchema(profile)
	}

	c.JSON(http.StatusOK, schemas.GetMe{
		ID:              user.ID,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		LastLoginAt:     user.LastLoginAt,
		PhoneNumber:     user.PhoneNumber,
		CountryDialCode: user.CountryDialCode.ToSchema(),
		Country:         user.Country.ToSchema(),
		UserProfile:     profileSchema,
		EasemobUsername: user.EasemobUsername,
		EasemobUUID:     user.EasemobUUID,
	})
}

// @Summary     Refresh token
// @Description Refresh token
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       request body schemas.RefreshTokenPayload true "Refresh token"
// @Success     200 {object} schemas.PairTokensResponse
// @Failure     400 {object} schemas.Response
// @Failure     401 {object} schemas.Response
// @Router      /auth/refresh_token [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	logs := logger.GetLogger(c)

	var payload schemas.RefreshTokenPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logs.Error().Err(err).Msg("bad request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, newRefreshToken, err := h.service.RefreshToken(payload.RefreshToken)
	if err != nil {
		logs.Error().Err(err).Msg("failed to refresh token")
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	logs.Info().Msg("token got refreshed")
	c.JSON(http.StatusOK, schemas.PairTokensResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	})
}
