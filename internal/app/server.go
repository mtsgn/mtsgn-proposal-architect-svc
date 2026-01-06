package app

import (
	"boilerplate-api/internal/app/routes"
	"boilerplate-api/internal/common/context"
	"fmt"
)

// Server represents the server
type Server struct {
	ctx *context.AppContext
}

// NewServer returns a new Server instance
func NewServer(ctx *context.AppContext) *Server {
	return &Server{
		ctx: ctx,
	}
}

// Run starts the server
func (s *Server) Run() {
	// Initialize router
	router := &routes.Router{}
	err := router.NewRouter(s.ctx)
	if err != nil {
		panic(fmt.Errorf("failed to initialize router: %w", err))
	}

	// Setup router
	router.SetupRouter(s.ctx)

	// Start the server
	router.Run(fmt.Sprintf(":%d", s.ctx.Cfg.Server.Port))
}

//func (s *Server) MigrateDB() {
//	err := s.ctx.DB.AutoMigrate()
//	if err != nil {
//		glog.Fatal("Database migration failed:", err)
//	}
//
//	glog.Println("Database migration completed successfully")
//}
