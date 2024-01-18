package api

import (
	"github.com/GitSplit-org/backend/api/github"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("")
	{
		github.ApplyRoutes(api)
	}
}
