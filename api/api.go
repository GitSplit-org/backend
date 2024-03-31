package api

import (
	"github.com/GitSplit-org/backend/api/projects"
	"github.com/GitSplit-org/backend/api/user"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("")
	{
		projects.ApplyRoutes(api)
		user.ApplyRoutes(api)
	}
}
