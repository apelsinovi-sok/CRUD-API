package http

import (
	"github.com/gin-gonic/gin"
	"test/interanl/core/interfaces"
)

type httpServer struct {
	router  *gin.Engine
	usecase interfaces.UseCase
}

func New(usecase interfaces.UseCase) *httpServer {
	s := &httpServer{
		router:  gin.Default(),
		usecase: usecase,
	}
	s.setExternalParamInHandlers()
	s.SetRoutes()
	return s
}

func (s *httpServer) Run() {
	err := s.router.Run()
	if err != nil {
		panic(err.Error())
	}
}

func (s *httpServer) setExternalParamInHandlers() {
	s.router.Use(func(c *gin.Context) {
		c.Set("usecase", &s.usecase)
		c.Next()
	})
}
