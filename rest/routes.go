package rest

import (
	"apocalypse/middleware"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Router(db *sql.DB) {

	var rest IRest = &Rest{
		DB: db,
	}

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	
	router.GET("/ping", rest.GetAll)

	router.Run(":8000")
}
