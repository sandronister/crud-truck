package web

import (
	"fmt"

	"github.com/labstack/echo"
)

type Server struct {
	router  *echo.Echo
	webPort string
}

func NewServer(webPort string) *Server {
	return &Server{
		router:  echo.New(),
		webPort: fmt.Sprintf(":%s", webPort),
	}
}

func (s *Server) Run() error {
	return s.router.Start(s.webPort)
}
