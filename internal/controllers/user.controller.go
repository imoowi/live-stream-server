package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/imoowi/comer/interfaces"
	"github.com/imoowi/comer/utils/response"
	"github.com/imoowi/live-stream-server/internal/models"
	"github.com/imoowi/live-stream-server/internal/services"
	"github.com/spf13/cast"
)

func UserPageList(c *gin.Context) {
	var filter interfaces.IFilter = &models.UserFilter{}
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	res, err := services.User.PageList(c, &filter)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(res, c)
}

func UserOne(c *gin.Context) {
	id := c.DefaultQuery(`id`, `0`)
	res, err := services.User.One(c, cast.ToUint(id))
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(res, c)
}
func UserAdd(c *gin.Context) {
	var addModel models.UserAdd
	err := c.ShouldBindBodyWith(&addModel, binding.JSON)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	res, err := services.User.Add(c, &addModel)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(res, c)
}

func UserUpdate(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`pls input id`, http.StatusBadRequest, c)
		return
	}
	addModel := make(map[string]any)
	err := c.ShouldBindBodyWith(&addModel, binding.JSON)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	res, err := services.User.Update(c, addModel, cast.ToUint(id))
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(res, c)
}

func UserDelete(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`pls input id`, http.StatusBadRequest, c)
		return
	}
	res, err := services.User.Delete(c, cast.ToUint(id))
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(res, c)
}
