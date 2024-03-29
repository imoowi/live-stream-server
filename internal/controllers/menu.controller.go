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
//	@Tags		menu
//	@Accept		application/json
//	@Produce	application/json
//	@Param		Authorization	header		string							true	"Bearer 用户令牌"
//	@Param		{object}		query		models.MenuFilter				false	"query参数"
//	@Success	200				{object}	response.PageListT[models.Menu]	"成功"
//	@Failure	400				"请求错误"
//	@Failure	401				"token验证失败"
//	@Failure	500				"内部错误"
//	@Router		/api/menus [get]
func MenuPageList(c *gin.Context) {
	var filter interfaces.IFilter = &models.MenuFilter{}
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
	result, err := services.Menu.PageList(c, &filter)
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
//	@Tags		menu
//	@Accept		application/json
//	@Produce	application/json
//	@Param		Authorization	header		string		true	"Bearer 用户令牌"
//	@Param		id				path		int			true	"id"
//	@Success	200				{object}	models.Menu	"成功"
//	@Failure	400				"请求错误"
//	@Failure	401				"token验证失败"
//	@Failure	500				"内部错误"
//	@Router		/api/menus/{id} [get]
func MenuOne(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`pls input id`, http.StatusBadRequest, c)
		return
	}

	one, err := services.Menu.One(c, cast.ToUint(id))
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(one, c)
}

//	@Summary	新增(add)
//	@Tags		menu
//	@Accept		application/json
//	@Produce	application/json
//	@Param		Authorization	header	string		true	"Bearer 用户令牌"
//	@Param		{object}		body	models.Menu	true	"body"
//	@Success	200
//	@Failure	400	"请求错误"
//	@Failure	401	"token验证失败"
//	@Failure	500	"内部错误"
//	@Router		/api/menus [post]
func MenuAdd(c *gin.Context) {
	menu := &models.Menu{}
	err := c.ShouldBindBodyWith(&menu, binding.JSON)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	newId, err := services.Menu.Add(c, menu)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(newId, c)
}

//	@Summary	更新(update)
//	@Tags		menu
//	@Accept		application/json
//	@Produce	application/json
//	@Param		Authorization	header	string		true	"Bearer 用户令牌"
//	@Param		id				path	int			true	"id"
//	@Param		{object}		body	models.Menu	true	"body"
//	@Success	200
//	@Failure	400	"请求错误"
//	@Failure	401	"token验证失败"
//	@Failure	500	"内部错误"
//	@Router		/api/menus/{id} [put]
func MenuUpdate(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`pls input id`, http.StatusBadRequest, c)
		return
	}
	menu := make(map[string]any)
	err := c.ShouldBindBodyWith(&menu, binding.JSON)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	updated, err := services.Menu.Update(c, menu, cast.ToUint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(err.Error(), http.StatusNotFound, c)
			return
		}
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(updated, c)
}

//	@Summary	删除(delete)
//	@Tags		menu
//	@Accept		application/json
//	@Produce	application/json
//	@Param		Authorization	header	string	true	"Bearer 用户令牌"
//	@Param		id				path	int		true	"id"
//	@Success	200
//	@Failure	400	"请求错误"
//	@Failure	401	"token验证失败"
//	@Failure	500	"内部错误"
//	@Router		/api/menus/{id} [delete]
func MenuDel(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`pls input id`, http.StatusBadRequest, c)
		return
	}
	deleted, err := services.Menu.Delete(c, cast.ToUint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(err.Error(), http.StatusNotFound, c)
			return
		}
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(deleted, c)
}
