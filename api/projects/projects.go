package projects

import (
	"fmt"
	"net/http"

	"github.com/GitSplit-org/backend/config/dbconfig"
	"github.com/GitSplit-org/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ApplyRoutes(r *gin.RouterGroup) {
	api := r.Group("/projects")
	{
		api.POST("", AddProject)
		api.GET("", GetAllProjects)
	}
	apiSingle := r.Group("/project")
	{
		apiSingle.GET("", getProjectById)
	}
}

func AddProject(c *gin.Context) {
	var req AddProjectRequest
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := dbconfig.GetDb()
	instance, _ := db.DB()
	defer instance.Close()

	var project models.Project
	id := uuid.New()
	project.Id = id.String()
	project.Name = req.Name
	project.Url = req.Url
	project.Description = req.Description
	project.Instagram = req.Instagram
	project.Twitter = req.Twitter
	project.Linkedin = req.Linkedin
	project.Owner = req.Owner
	project.Image = req.Image
	err := db.Create(&project).Error
	if err != nil {
		fmt.Println("error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "project added successfully",
		"data":    project,
	})
}

func GetAllProjects(c *gin.Context) {
	db := dbconfig.GetDb()
	instance, _ := db.DB()
	defer instance.Close()
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

func getProjectById(c *gin.Context) {
	db := dbconfig.GetDb()
	instance, _ := db.DB()
	defer instance.Close()

	id := c.Query("id")

	var project models.Project
	err := db.Model(&models.Project{}).Where("id = ?", id).First(&project).Error
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "project fetched successfully",
		"data":    project,
	})

}
