package http

import "test/interanl/delivery/http/internal"

func (s *httpServer) SetRoutes() {
	s.setApi()
}
func (s *httpServer) setApi() {
	s.router.POST("/user", internal.CreateUser)
	s.router.GET("/user/:id", internal.GetUser)
	s.router.PUT("/user/:id", internal.UpdateUser)
}
