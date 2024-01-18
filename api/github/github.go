package github

import (
	"fmt"
	"net/http"

	"github.com/GitSplit-org/backend/config/dbconfig"
	"github.com/GitSplit-org/backend/models"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	api := r.Group("/github")
	{
		api.POST("", AddProject)
		api.GET("", GetAllProjects)
	}
}

func AddProject(c *gin.Context) {
	var req AddProjectRequest
	if err := c.BindJSON(&req); err != nil {
		fmt.Println("error", err)
		return
	}
	db := dbconfig.GetDb()
	var project models.Project
	project.Name = req.Name
	project.Url = req.Url
	project.Description = req.Description
	err := db.Create(&project).Error
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "project added successfully",
		"data":    project,
	})
}

func GetAllProjects(c *gin.Context) {
	db := dbconfig.GetDb()
	var projects []models.Project
	err := db.Model(&models.Project{}).Find(&projects).Error
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "projects fetched successfully",
		"data":    projects,
	})
}
