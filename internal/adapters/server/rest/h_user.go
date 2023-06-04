package rest

import (
	"birthday-bot/internal/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router   /users [post]
// @Tags     users
// @Param    body  body  entities.UserCUSt false  "body"
// @Success  200  {object}
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hUserCreate(c *gin.Context) {
	reqObj := &entities.UserCUSt{}
	if !BindJSON(c, reqObj) {
		return
	}

	result, err := o.ucs.UserCreate(o.getRequestContext(c), reqObj)
	if Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": result})
}

// @Router   /users/:id [get]
// @Tags     users
// @Param    id path int64 true "id"
// @Produce  json
// @Success  200  {object}  entities.UserSt
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hUserGet(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if Error(c, err) {
		return
	}

	result, err := o.ucs.UserGet(o.getRequestContext(c), id)
	if Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router   /users/:id [put]
// @Tags     users
// @Param    id path string true "id"
// @Param    body  body  entities.UserCUSt false  "body"
// @Produce  json
// @Success  200
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hUserUpdate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if Error(c, err) {
		return
	}

	reqObj := &entities.UserCUSt{}
	if !BindJSON(c, reqObj) {
		return
	}

	Error(c, o.ucs.UserUpdate(o.getRequestContext(c), id, reqObj))
}

// @Router   /users/:id [delete]
// @Tags     users
// @Param    id path int64 true "id"
// @Success  200
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hUserDelete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if Error(c, err) {
		return
	}

	Error(c, o.ucs.UserDelete(o.getRequestContext(c), id))
}
