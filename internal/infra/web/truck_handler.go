package web

import "github.com/sandronister/crud-truck/internal/infra/handler"

func (s *Server) AddTruckHandler(t *handler.TruckHandler) {
	public := s.router.Group("/api/v1")

	public.POST("/trucks", t.Save)
	public.GET("/trucks", t.FindAll)
	public.GET("/trucks/:id", t.FindByID)
	public.PUT("/trucks/:id", t.Update)
	public.DELETE("/trucks/:id", t.Delete)
}
