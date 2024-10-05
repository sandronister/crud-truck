package web

import "github.com/sandronister/crud-truck/internal/infra/handler"

func (s *Server) AddDriverHandler(d *handler.DriverHandler) {
	public := s.router.Group("/api/v1")

	public.POST("/drivers", d.Save)
	public.GET("/drivers", d.FindAll)
	public.GET("/drivers/:id", d.FindByID)
	public.PUT("/drivers/:id", d.Update)
	public.DELETE("/drivers/:id", d.Delete)
}
