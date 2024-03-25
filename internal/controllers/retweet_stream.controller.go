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

//	@Summary	分页列表(pagelist)
//	@Tags		RetweetStream(转推流)
//	@Accept		application/json
//	@Produce	application/json
//	@Param		Authorization	header		string										true	"Bearer 用户令牌"
//	@Param		{object}		query		models.RetweetStreamFilter					false	"query参数"
//	@Success	200				{object}	response.PageListT[models.RetweetStream]	"成功"
//	@Failure	400				"请求错误"
//	@Failure	401				"token验证失败"
//	@Failure	500				"内部错误"
//	@Router		/api/retweet-streams [get]
func RetweetStreamPageList(c *gin.Context) {
	var filter interfaces.IFilter = &models.RetweetStreamFilter{}
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
	result, err := services.RetweetStream.PageList(c, &filter)
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

//	@Summary	详情(one)
//	@Tags		RetweetStream(转推流)
//	@Accept		application/json
//	@Produce	application/json
//	@Param		Authorization	header	string	true	"Bearer 用户令牌"
//	@Param		id				path	int		true	"id"
//	@Success	200
//	@Failure	400	"请求错误"
//	@Failure	401	"token验证失败"
//	@Failure	500	"内部错误"
//	@Router		/api/retweet-streams/{id} [get]
func RetweetStreamOne(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`pls input id`, http.StatusBadRequest, c)
		return
	}

	one, err := services.RetweetStream.One(c, cast.ToUint(id))
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(one, c)
}

//	@Summary	新增(add)
//	@Tags		RetweetStream(转推流)
//	@Accept		application/json
//	@Produce	application/json
//	@Param		Authorization	header	string					true	"Bearer 用户令牌"
//	@Param		{object}		body	models.RetweetStream	true	"body"
//	@Success	200
//	@Failure	400	"请求错误"
//	@Failure	401	"token验证失败"
//	@Failure	500	"内部错误"
//	@Router		/api/retweet-streams [post]
func RetweetStreamAdd(c *gin.Context) {
	var model *models.RetweetStream
	err := c.ShouldBindBodyWith(&model, binding.JSON)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	newId, err := services.RetweetStream.Add(c, model)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(newId, c)
}

//	@Summary	更新(update)
//	@Tags		RetweetStream(转推流)
//	@Accept		application/json
//	@Produce	application/json
//	@Param		Authorization	header	string					true	"Bearer 用户令牌"
//	@Param		id				path	int						true	"id"
//	@Param		{object}		body	models.RetweetStream	true	"body"
//	@Success	200
//	@Failure	400	"请求错误"
//	@Failure	401	"token验证失败"
//	@Failure	500	"内部错误"
//	@Router		/api/retweet-streams/{id} [put]
func RetweetStreamUpdate(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`pls input id`, http.StatusBadRequest, c)
		return
	}
	model := make(map[string]any)
	err := c.ShouldBindBodyWith(&model, binding.JSON)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	updated, err := services.RetweetStream.Update(c, model, cast.ToUint(id))
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(updated, c)
}

//	@Summary	删除(delete)
//	@Tags		RetweetStream(转推流)
//	@Accept		application/json
//	@Produce	application/json
//	@Param		Authorization	header	string	true	"Bearer 用户令牌"
//	@Param		id				path	int		true	"id"
//	@Success	200
//	@Failure	400	"请求错误"
//	@Failure	401	"token验证失败"
//	@Failure	500	"内部错误"
//	@Router		/api/retweet-streams/{id} [delete]
func RetweetStreamDel(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`pls input id`, http.StatusBadRequest, c)
		return
	}
	deleted, err := services.RetweetStream.Delete(c, cast.ToUint(id))
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(deleted, c)
}
