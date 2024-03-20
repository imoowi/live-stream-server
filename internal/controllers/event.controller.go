/*
generated by comer,https://github.com/imoowi/comer
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/imoowi/comer/interfaces"
	"github.com/imoowi/comer/utils/response"
	"github.com/imoowi/live-stream-server/internal/models"
	"github.com/imoowi/live-stream-server/internal/services"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// @Summary	分页列表(pagelist)
// @Tags		Event(直播活动)
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header		string				true	"Bearer 用户令牌"
// @Param		{object}		query		models.EventFilter	false	"query参数"
// @Success	200				{object}	response.PageList	"成功"
// @Failure	400				"请求错误"
// @Failure	401				"token验证失败"
// @Failure	500				"内部错误"
// @Router		/api/events [get]
func EventPageList(c *gin.Context) {
	var filter interfaces.IFilter = &models.EventFilter{}
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}

	if 0 >= filter.GetPage() { //如果不传Page，默认为1
		filter.SetPage(1)
	}
	if 0 >= filter.GetPageSize() { //如果不传PageSize，默认取20条
		filter.SetPageSize(20)
	}
	if filter.GetPageSize() > 1000 {
		response.Error(`每一页不能超过1000条记录`, http.StatusBadRequest, c)
		return
	}
	var mt interfaces.IModel = &models.Event{}
	result, err := services.Event.PageList(c, &filter, &mt)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(err.Error(), http.StatusNotFound, c)
			return
		}
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(result, c)
}

// @Summary	详情(one)
// @Tags		Event(直播活动)
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header	string	true	"Bearer 用户令牌"
// @Param		id				path	int		true	"id"
// @Success	200
// @Failure	400	"请求错误"
// @Failure	401	"token验证失败"
// @Failure	500	"内部错误"
// @Router		/api/events/{id} [get]
func EventOne(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`pls input id`, http.StatusBadRequest, c)
		return
	}

	var f interfaces.IFilter = &models.EventFilter{}
	var mt interfaces.IModel = &models.Event{}
	one, err := services.Event.One(c, &f, cast.ToUint(id), &mt)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(one, c)
}

// @Summary	新增(add)
// @Tags		Event(直播活动)
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header	string			true	"Bearer 用户令牌"
// @Param		{object}		body	models.Event	true	"body"
// @Success	200
// @Failure	400	"请求错误"
// @Failure	401	"token验证失败"
// @Failure	500	"内部错误"
// @Router		/api/events [post]
func EventAdd(c *gin.Context) {
	var event interfaces.IModel = &models.Event{}
	err := c.ShouldBindBodyWith(&event, binding.JSON)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	var f interfaces.IFilter = &models.EventFilter{}
	var mt interfaces.IModel = &models.Event{}
	newId, err := services.Event.Add(c, &f, &event, &mt)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(newId, c)
}

// @Summary	更新(update)
// @Tags		Event(直播活动)
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header	string			true	"Bearer 用户令牌"
// @Param		id				path	int				true	"id"
// @Param		{object}		body	models.Event	true	"body"
// @Success	200
// @Failure	400	"请求错误"
// @Failure	401	"token验证失败"
// @Failure	500	"内部错误"
// @Router		/api/events/{id} [put]
func EventUpdate(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`pls input id`, http.StatusBadRequest, c)
		return
	}
	var event interfaces.IModel = &models.Event{}
	err := c.ShouldBindBodyWith(&event, binding.JSON)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	var f interfaces.IFilter = &models.EventFilter{}
	var mt interfaces.IModel = &models.Event{}
	updated, err := services.Event.Update(c, &f, &event, cast.ToUint(id), &mt)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(updated, c)
}

// @Summary	删除(delete)
// @Tags		Event(直播活动)
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header	string	true	"Bearer 用户令牌"
// @Param		id				path	int		true	"id"
// @Success	200
// @Failure	400	"请求错误"
// @Failure	401	"token验证失败"
// @Failure	500	"内部错误"
// @Router		/api/events/{id} [delete]
func EventDel(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`pls input id`, http.StatusBadRequest, c)
		return
	}
	var f interfaces.IFilter = &models.EventFilter{}
	var mt interfaces.IModel = &models.Event{}
	deleted, err := services.Event.Delete(c, &f, cast.ToUint(id), &mt)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(deleted, c)
}

// @Summary	状态Map
// @Tags		Event(直播活动)
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header	string	true	"Bearer 用户令牌"
// @Param		id				path	int		true	"id"
// @Success	200
// @Failure	400	"请求错误"
// @Failure	401	"token验证失败"
// @Failure	500	"内部错误"
// @Router		/api/event/status [get]
func EventStatus(c *gin.Context) {
	response.OK(services.Event.Status(), c)
}
