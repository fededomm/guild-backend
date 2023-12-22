package routes

import (
	"context"
	"guild-be/src/custom"
	"guild-be/src/database"
	"guild-be/src/models"
	"guild-be/src/rest/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

type Rest struct {
	DB  *database.DBService
	Val *validator.Validate
}

var arrToValid utils.ArrToValid

// @Summary	Get all users
// @Description
// @Produce	json
// @Success	200	{array}		models.User
// @Failure	400	{object}	custom.BadRequestError
// @Failure	404	{object}	custom.NotFoundError
// @Failure	500	{object}	custom.InternalServerError
// @Tags		Guild
// @Router		/guild/ [get]
func (r *Rest) GetAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	list, err := r.DB.GetAll(ctx)
	if err != nil {
		log.Err(err).Msg(err.Error())
		c.JSON(500, custom.InternalServerError{Code: 500, Message: err.Error()})
		return
	}
	c.JSON(200, list)
}

// @Summary		Get all pg for user
// @Description	Get all pg for user
// @Produce		json
// @Param			name	path		string			true	"name"
// @Success		200		{object}	custom.Success
// @Failure		400		{object}	custom.BadRequestError
// @Failure		404		{object}	custom.NotFoundError
// @Failure		500		{object}	custom.InternalServerError
// @Tags			Guild
// @Router			/guild/{name} [get]
func (r *Rest) GetAllPgByUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	param := c.Param("name")
	list, err := r.DB.GetAllPgForUser(ctx, param)
	if err != nil {
		log.Err(err).Msg(err.Error())
		c.JSON(500, custom.InternalServerError{Code: 500, Message: err.Error()})
		return
	}
	c.JSON(200, list)
}

// @Summary	Insert one user
// @Description
// @Accept		json
// @Produce	json
// @Param		user	body		custom.ExampleBodyUser	true	"User"
// @Success	201		{object}	custom.Created
// @Failure	400		{object}	custom.BadRequestError
// @Failure	404		{object}	custom.NotFoundError
// @Failure	500		{object}	custom.InternalServerError
// @Tags		Guild
// @Router		/guild/usr [post]
func (r *Rest) PostUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := new(models.User)
	if err := c.BindJSON(user); err != nil {
		log.Err(err).Msg(err.Error())
		c.JSON(400, custom.BadRequestError{Code: 400, Message: err.Error()})
		return
	}

	if err := r.DB.InsertUser(ctx, *user); err != nil {
		log.Err(err).Msg(err.Error())
		c.JSON(500, custom.InternalServerError{Code: 500, Message: err.Error()})
		return
	}
	c.JSON(201, custom.Created{Code: 201, Message: "Created"})
}

// @Summary Insert one pg
// @Description  Insert one pg
// @Produce json
// @Param		pg	body		custom.ExampleBodyPg	true	"User"
// @Success	201		{object}	custom.Created
// @Failure	400		{object}	custom.BadRequestError
// @Failure	404		{object}	custom.NotFoundError
// @Failure	500		{object}	custom.InternalServerError
// @Tags		Guild
// @Router /guild/pg [post]
func (r *Rest) PostPg(c *gin.Context) {
	var rank, class []string
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pg := new(models.Personaggio)
	if err := c.BindJSON(pg); err != nil {
		log.Err(err).Msg(err.Error())
		c.JSON(400, custom.BadRequestError{Code: 400, Message: err.Error()})
		return
	}
	fetchRank, err := utils.FetchArrayByName(r.DB.DB, rank, "rank")
	if err != nil {
		log.Err(err).Msg(err.Error())
		c.JSON(500, custom.InternalServerError{Code: 500, Message: err.Error()})
		return
	}
	fetchClass, err := utils.FetchArrayByName(r.DB.DB, class, "class")
	if err != nil {
		log.Err(err).Msg(err.Error())
		c.JSON(500, custom.InternalServerError{Code: 500, Message: err.Error()})
		return
	}
	arrToValid = utils.ArrToValid{
		Rank:  fetchRank,
		Class: fetchClass,
	}
	if err := arrToValid.CustomArrayRankClassValidatorGin(pg, r.Val); err != nil {
		log.Err(err).Msg(err.Error())
		c.JSON(400, custom.BadRequestError{Code: 400, Message: err.Error()})
		return
	}
	if err := r.DB.InsertPg(ctx, *pg); err != nil {
		log.Err(err).Msg(err.Error())
		c.JSON(500, custom.InternalServerError{Code: 500, Message: err.Error()})
		return
	}
	c.JSON(201, custom.Created{Code: 201, Message: "Created"})
}
