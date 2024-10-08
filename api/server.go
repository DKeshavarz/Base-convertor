package api

import "github.com/labstack/echo/v4"

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
	s.Server.GET("/convert-base", convertBaseHandler)
	err := s.Server.Start(s.ListenAdr)
	return err
}