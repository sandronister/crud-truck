package web

import "github.com/sandronister/crud-truck/internal/infra/handler"

func (s *Server) AddLinkHandler(t *handler.LinkHandler) {
	public := s.router.Group("/api/v1")

	public.POST("/links", t.Save)
	public.DELETE("/links/:driver_id/:truck_id", t.Delete)
	public.GET("/links/:driver_id", t.ListByDriver)
}
