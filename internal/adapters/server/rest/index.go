package rest

import (
	"birthday-bot/internal/adapters/logger"
	"birthday-bot/internal/domain/usecases"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type St struct {
	lg  logger.Lite
	ucs *usecases.St
}

func GetHandler(lg logger.Lite, ucs *usecases.St, withCors bool) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// middlewares

	r.Use(MwRecovery(lg, nil))
	if withCors {
		r.Use(MwCors())
	}

	// handlers

	s := &St{lg: lg, ucs: ucs}

	// healthcheck
	r.GET("/healthcheck", func(c *gin.Context) { c.Status(http.StatusOK) })

	// users
	r.GET("/users/:id", s.hUserGet)
	r.POST("/users", s.hUserCreate)
	r.PUT("/users/:id", s.hUserUpdate)
	r.DELETE("/users/:id ", s.hUserDelete)
	return r
}

func (o *St) getRequestContext(c *gin.Context) context.Context {
	return context.Background()
}
