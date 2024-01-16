package server

import (
	"fmt"

	"github.com/VADIMREX/GoMock/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	actions map[string](func(gin.H)gin.H)
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	var srv Server
	srv.config = config
	srv.actions = map[string](func(gin.H)gin.H){
		"reload-config": srv.reloadConfig,
	}
	return &srv
}

func (s *Server) getInterface(ctx *gin.Context) {
	ctx.SetCookie("interface-path", s.config.InterfacePath, 0, s.config.InterfacePath, "localhost", false, false)
	ctx.File("index.html")
	ctx.Status(200)
}

func get[TRes any](obj *gin.H, field string) (TRes, error) {
	var val TRes
	rawVal, ok := (*obj)[field]
	if !ok {
		return val, fmt.Errorf("no field")
	}
	val, ok = rawVal.(TRes)
	if !ok {
		return val, fmt.Errorf("failed cast")
	}
	return val, nil
}

func (s *Server) reloadConfig(msg gin.H) gin.H {
	s.config.Reload()
	return gin.H{
		"interface-path": s.config.InterfacePath,
	}
}

func (s *Server) interfaceHandler(ctx *gin.Context) {
	var msg gin.H
	ctx.BindJSON(&msg)
	action, err := get[string](&msg, "action")
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 1,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, s.actions[action](msg))
}

func (s *Server) MockHandler(ctx *gin.Context) {
	if ctx.Request.RequestURI == s.config.InterfacePath {
		if ctx.Request.Method == "GET" {
			s.getInterface(ctx)
			return
		}
		if ctx.Request.Method == "POST" {
			s.interfaceHandler(ctx)
			return
		}
	}
}