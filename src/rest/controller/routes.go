package controller

import (
	"context"
	"guild-be/src/config"
	"guild-be/src/database"
	"guild-be/src/observability"
	"guild-be/src/rest/middleware"
	"guild-be/src/rest/routes"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(db *database.DBService, conf *config.GlobalConfig) {

	var rest routes.IRest = &routes.Rest{
		DB:  db,
		Val: validator.New(),
	}

	ctx := context.Background()
	if conf.Observability.Enable == false {
		// init tracer
		trace, err := observability.InitTracer(ctx, conf.Observability.Endpoint, conf.Observability.ServiceName)
		if err != nil {
			panic("trace error" + err.Error())
		}
		defer trace(ctx)

		// init metric
		metric, err := observability.InitMetric(ctx, conf.Observability.Endpoint, conf.Observability.ServiceName)
		if err != nil {
			panic("metric error" + err.Error())
		}

		defer metric(ctx)
	}

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	v1 := router.Group("/api/v1")
	{
		guild := v1.Group("/guild")
		{
			guild.GET("", rest.GetAllUsers)
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
