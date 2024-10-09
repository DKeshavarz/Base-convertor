package api

import ("github.com/labstack/echo/v4"
"github.com/labstack/echo/v4/middleware"
)

type Server struct{
	Server    *echo.Echo
	ListenAdr  string
}
func NewServer() (s *Server){
	s = &Server{
		echo.New(),
		":8000",
	}
	return s
}

func (s *Server) StartServer()error{
	//here we set handles

	s.Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
    }))

	s.Server.Static("/" , "Ui")

	s.Server.GET("/convert-base", convertBaseHandler)
	err := s.Server.Start(s.ListenAdr)
	return err
}