package routes

import (
	"context"
	"guild-be/src/database"
	"guild-be/src/models"
	"guild-be/src/custom"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Rest struct {
	DB    *database.DBService
	Rank  []string
	Class []string
	Val   *validator.Validate
}

var validArr ArrToValid

//	@Summary	Get all users
//	@Description
//	@Produce	json
//	@Success	200	{array}		models.User
//	@Failure	400	{object}	custom.BadRequestError
//	@Failure	404	{object}	custom.NotFoundError
//	@Failure	500	{object}	custom.InternalServerError
//	@Tags		Guild
//	@Router		/guild/ [get]
func (r *Rest) GetAll(c *gin.Context) {
	list, err := r.DB.GetAll()
	if err != nil {
		c.JSON(500, custom.InternalServerError{Code: 500, Message: err.Error()})
		return
	}
	c.JSON(200, list)
}

//	@Summary		Get all pg for user
//	@Description	Get all pg for user
//	@Produce		json
//	@Param			name	path		string			true	"name"
//	@Success		200		{object}	custom.Success	
//	@Failure		400		{object}	custom.BadRequestError
//	@Failure		404		{object}	custom.NotFoundError
//	@Failure		500		{object}	custom.InternalServerError
//	@Tags			Guild
//	@Router			/guild/{name} [get]
func (r *Rest) GetAllPgByUser(c *gin.Context) {
	param := c.Param("name")
	list, err := r.DB.GetAllPgForUser(param)
	if err != nil {
		c.JSON(500, custom.InternalServerError{Code: 500, Message: err.Error()})
		return
	}
	c.JSON(200, list)
}

//	@Summary	Insert one user
//	@Description
//	@Accept		json
//	@Produce	json
//	@Param		user	body		custom.ExampleBody	true	"User"
//	@Success	201		{object}	custom.Created
//	@Failure	400		{object}	custom.BadRequestError
//	@Failure	404		{object}	custom.NotFoundError
//	@Failure	500		{object}	custom.InternalServerError
//	@Tags		Guild
//	@Router		/guild/ [post]
func (r *Rest) PostOne(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := new(models.User)
	if err := c.BindJSON(user); err != nil {
		c.JSON(400, custom.BadRequestError{Code: 400, Message: err.Error()})
		return
	}
	validArr.Rank = r.Rank
	validArr.Class = r.Class
	if err := validArr.CustomArrayValidatorGin(user, r.Val); err != nil {
		c.JSON(400, custom.BadRequestError{Code: 400, Message: err.Error()})
		return
	}
	if err := r.DB.DoTrx(ctx, *user); err != nil {
		c.JSON(500, custom.InternalServerError{Code: 500, Message: err.Error()})
		return
	}
	c.JSON(201, custom.Created{Code: 201, Message: "Created", Body: *user})
}
