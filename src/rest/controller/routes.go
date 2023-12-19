package controller

import (
	"guild-be/src/rest/routes"
	"guild-be/src/database"
	"guild-be/src/rest/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


func Router(db *database.DBService, rank []string, class []string) {

	var rest routes.IRest = &routes.Rest{
		DB:    db,
		Rank:  rank,
		Class: class,
		Val:   validator.New(),
	}

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	v1 := router.Group("/api/v1")
	{
		guild := v1.Group("/guild")
		{
			guild.GET("", rest.GetAll)
			guild.POST("", rest.PostOne)
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
