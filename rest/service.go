package rest

import (
	"guild-be/custom"
	"guild-be/database"
	"guild-be/models"
	"context"
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	DB   *sql.DB
	Rank models.Rank
}

// @Summary	Get all users
// @Description
// @Produce	json
// @Success	200	{array}		models.User
// @Failure	400	{object}	custom.BadRequestError
// @Failure	404	{object}	custom.NotFoundError
// @Failure	500	{object}	custom.InternalServerError
// @Tags		Guild
// @Router		/guild/getall [get]
func (r *Rest) GetAll(c *gin.Context) {
	list, err := database.GetAll(r.DB, c)
	if err != nil {
		c.JSON(500, custom.InternalServerError{Code: 500, Message: err.Error()})
		return
	}
	c.JSON(200, list)
}

// @Summary	Insert one user
// @Description
// @Accept		json
// @Produce	json
// @Param		user	body		custom.ExampleBody	true	"User"
// @Success	201		{object}	custom.Created
// @Failure	400		{object}	custom.BadRequestError
// @Failure	404		{object}	custom.NotFoundError
// @Failure	500		{object}	custom.InternalServerError
// @Tags		Guild
// @Router		/guild/insert [post]
func (r *Rest) PostOne(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := new(models.User)
	if err := c.BindJSON(user); err != nil {
		c.JSON(400, custom.BadRequestError{Code: 400, Message: err.Error()})
		return
	}
	if err := CustomValidatorGin(user, r.Rank); err != nil {
		c.JSON(400, custom.BadRequestError{Code: 400, Message: err.Error()})
		return
	}
	if err := database.DoTrx(r.DB, ctx, *user); err != nil {
		c.JSON(500, custom.InternalServerError{Code: 500, Message: err.Error()})
		return
	}
	c.JSON(201, custom.Created{Code: 201, Message: "Created", Body: *user})
}
