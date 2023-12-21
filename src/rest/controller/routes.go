package controller

import (
	"guild-be/src/database"
	"guild-be/src/rest/middleware"
	"guild-be/src/rest/routes"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(db *database.DBService) {

	var rest routes.IRest = &routes.Rest{
		DB:    db,
		Val:   validator.New(),
	}

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	v1 := router.Group("/api/v1")
	{
		guild := v1.Group("/guild")
		{
			guild.GET("", rest.GetAll)
			guild.POST("/usr", rest.PostUser)
			guild.POST("/pg", rest.PostPg)
			guild.GET(":name", rest.GetAllPgByUser)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/heathcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Alive!",
		})
	})
	router.Run(":8000")
}
