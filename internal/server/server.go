package server

import (
	"mock/internal/configs"

	"github.com/gin-gonic/gin"
)

type Server struct {
	server  *gin.Engine
	addr    string
	Methods map[string]func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
}

func NewServer(addr string) *Server {
	r := gin.Default()
	return &Server{
		server: r,
		addr:   addr,
		Methods: map[string]func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes{
			"get":    r.GET,
			"post":   r.POST,
			"patch":  r.PATCH,
			"delete": r.DELETE,
			"put":    r.PUT,
		},
	}
}

func (s *Server) Init(apis []configs.Api) {
	for _, v := range apis {
		s.Methods[v.Method](v.Path, func(c *gin.Context) {
			c.JSON(200, gin.H{"ok": "ok"})
		})
	}
	s.server.Run(s.addr)
}
