package routes

import (
	_ "boilerplate-api/docs" // This is important!
	middleware "boilerplate-api/internal/api/middlewares"
	v1 "boilerplate-api/internal/app/routes/v1"
	"boilerplate-api/internal/common/context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	*gin.Engine
}

func (r *Router) NewRouter(ctx *context.AppContext) error {
	// Set mode
	gin.SetMode(ctx.Cfg.Server.Mode)

	r.Engine = gin.New()

	// Disable automatic trailing slash redirects
	r.Engine.RedirectTrailingSlash = false
	r.Engine.RedirectFixedPath = false

	// Add recovery middleware
	//r.Use(glog.FileLoggerMiddleware())
	//r.Use(glog.FileLoggerMiddleware())

	//捕获panic
	r.Engine.Use(gin.Recovery())
	// Add request_id middleware
	r.Engine.Use(middleware.RequestID())
	// Add logger middleware
	r.Engine.Use(middleware.Logger())

	//r.Engine.Use(func(c *gin.Context) {
	//	defer func() {
	//		if err := recover(); err != nil {
	//			var buf [4096]byte
	//			n := runtime.Stack(buf[:], false)
	//			tmpStr := fmt.Sprintf("terr: %v, tpanic==> %s\n", err, string(buf[:n]))
	//			glog.Errorf(tmpStr)
	//		}
	//	}()
	//	c.Next()
	//})

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowMethods = []string{
		"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS",
	}
	config.AllowHeaders = []string{
		"Origin",
		"Content-Length",
		"Content-RoomType",
		"Accept",
		"Authorization",
		"X-Requested-With",
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Headers",
		"Access-Control-Allow-Methods",
	}
	config.ExposeHeaders = []string{
		"Content-Length",
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Headers",
		"Access-Control-Allow-Methods",
	}
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.MaxAge = 12 * 60 * 60 // 12 hours

	r.Engine.Use(cors.New(config))

	r.Engine.Use(func(c *gin.Context) {
		c.Set("db", ctx.DB)
		c.Set("cfg", ctx.Cfg)
		c.Next()
	})
	return nil
}

func (r *Router) Run(addr string) {
	r.Engine.Run(addr)
}

func (r *Router) SetupRouter(ctx *context.AppContext) {
	// Swagger route should be outside of any group to be accessible
	r.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health Check Route
	r.Engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	v1Group := r.Engine.Group("/api/v1")
	{
		v1.SetupV1APIRoutes(v1Group, ctx)
	}

}
