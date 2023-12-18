package rest

import (
	"guild-be/middleware"
	"database/sql"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(db *sql.DB, rank []string) {

	var rest IRest = &Rest{
		DB:   db,
		Rank: rank,
	}

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	v1 := router.Group("/api/v1")
	{
		guild := v1.Group("/guild")
		{
			guild.GET("/getall", rest.GetAll)
			guild.POST("/insert", rest.PostOne)
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
