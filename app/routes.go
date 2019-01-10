package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) routes() {
	s.router.GET("/admin", s.adminOnly(s.handleAdminIndex()))

	api := s.router.Group("/api")
	api.GET("healthcheck", s.handleHealthcheck())
	api.POST("items", s.handleAddItem())
}

// adminOnly will return HTTP 404 if not isAdmin returns false.
// If isAdmin returns true, execution is passed to the handler that was passed in.
func (s *Server) adminOnly(h gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !isAdmin(ctx.Request) {
			ctx.String(http.StatusNotFound, "not allowed")
			return
		}
		h(ctx)
	}
}

func isAdmin(req *http.Request) bool {
	return req.Header.Get("admin") == "true"
}

func (s *Server) handleAdminIndex() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "admin page")
	}
}

func (s *Server) handleHealthcheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	}
}

func (s *Server) handleAddItem() gin.HandlerFunc {
	type request struct {
		Name string
	}

	type response struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	return func(ctx *gin.Context) {
		req := request{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}

		if len(req.Name) == 0 {
			ctx.String(http.StatusBadRequest, "name is empty")
			return
		}

		res := response{ID: 1, Name: req.Name}
		ctx.JSON(http.StatusCreated, res)
	}
}
