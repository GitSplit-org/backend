package api

import (
	"github.com/GitSplit-org/backend/api/projects"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("")
	{
		projects.ApplyRoutes(api)
	}
}
