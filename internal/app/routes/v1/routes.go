package v1

import (
	"boilerplate-api/internal/api/handlers"
	"boilerplate-api/internal/common/context"
	"boilerplate-api/internal/repositories"
	"boilerplate-api/internal/services"
	"boilerplate-api/pkg/easemob"

	"github.com/gin-gonic/gin"
)

func SetupV1APIRoutes(router *gin.RouterGroup, ctx *context.AppContext) {

	//--------------------- init repositories ---------------------//
	userRepo := repositories.NewUserRepository(ctx)
	commonRepo := repositories.NewCommonRepository(ctx)
	profileRepo := repositories.NewProfileRepository(ctx)

	//----------------------- init services --------------------- //
	userService := services.NewUserService(ctx.Redis, userRepo)
	profileService := services.NewProfileService(profileRepo)
	authService := services.NewAuthService(commonRepo, userRepo, profileRepo, ctx)

	// ------------------ init clients --------------------- //
	_ = easemob.NewClient(ctx.Cfg.Easemob)

	authHandler := handlers.NewAuthHandler(
		authService,
		profileService,
		userService,
		ctx,
	)

	//----------------- routes -----------------//
	authGroupAPI := router.Group("/auth")
	{
		authGroupAPI.POST("/register", authHandler.Register)
		authGroupAPI.POST("/sign_in", authHandler.SignIn)
		authGroupAPI.POST("/refresh_token", authHandler.RefreshToken)
	}

}
