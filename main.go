package main

import (
	"fmt"

	"github.com/VADIMREX/GoMock/config"
	"github.com/VADIMREX/GoMock/server"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	r := gin.Default()
	s := server.NewServer(cfg)

	r.NoRoute(s.MockHandler)

	r.Run(fmt.Sprintf("localhost:%d", cfg.Port))
}