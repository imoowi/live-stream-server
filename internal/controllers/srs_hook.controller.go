/*
generated by comer,https://github.com/imoowi/comer
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/imoowi/comer/interfaces"
	"github.com/imoowi/comer/utils/response"
	"github.com/imoowi/live-stream-server/internal/global"
	"github.com/imoowi/live-stream-server/internal/models"
	"github.com/imoowi/live-stream-server/internal/services"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// @Summary	分页列表(pagelist)
// @Tags		SrsHook(srs钩子)
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header		string								true	"Bearer 用户令牌"
// @Param		{object}		query		models.SrsHookFilter				false	"query参数"
// @Success	200				{object}	response.PageListT[models.SrsHook]	"成功"
// @Failure	400				"请求错误"
// @Failure	401				"token验证失败"
// @Failure	500				"内部错误"
// @Router		/api/srs-hooks [get]
func SrsHookPageList(c *gin.Context) {
	var filter interfaces.IFilter = &models.SrsHookFilter{}
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
	result, err := services.SrsHook.PageList(c, &filter)
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
// @Tags		SrsHook(srs钩子)
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header	string	true	"Bearer 用户令牌"
// @Param		id				path	int		true	"id"
// @Success	200
// @Failure	400	"请求错误"
// @Failure	401	"token验证失败"
// @Failure	500	"内部错误"
// @Router		/api/srs-hooks/{id} [get]
func SrsHookOne(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`pls input id`, http.StatusBadRequest, c)
		return
	}

	one, err := services.SrsHook.One(c, cast.ToUint(id))
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(one, c)
}

// @Summary	新增(add)
// @Tags		SrsHook(srs钩子)
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header	string			true	"Bearer 用户令牌"
// @Param		{object}		body	models.SrsHook	true	"body"
// @Success	200
// @Failure	400	"请求错误"
// @Failure	401	"token验证失败"
// @Failure	500	"内部错误"
// @Router		/api/srs-hooks [post]
func SrsHookAdd(c *gin.Context) {
	srshook := &models.SrsHook{}
	err := c.ShouldBindBodyWith(&srshook, binding.JSON)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	newId, err := services.SrsHook.Add(c, srshook)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(newId, c)
}

// @Summary	更新(update)
// @Tags		SrsHook(srs钩子)
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header	string			true	"Bearer 用户令牌"
// @Param		id				path	int				true	"id"
// @Param		{object}		body	models.SrsHook	true	"body"
// @Success	200
// @Failure	400	"请求错误"
// @Failure	401	"token验证失败"
// @Failure	500	"内部错误"
// @Router		/api/srs-hooks/{id} [put]
func SrsHookUpdate(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`pls input id`, http.StatusBadRequest, c)
		return
	}
	srshook := make(map[string]any)
	err := c.ShouldBindBodyWith(&srshook, binding.JSON)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	updated, err := services.SrsHook.Update(c, srshook, cast.ToUint(id))
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(updated, c)
}

// @Summary	删除(delete)
// @Tags		SrsHook(srs钩子)
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header	string	true	"Bearer 用户令牌"
// @Param		id				path	int		true	"id"
// @Success	200
// @Failure	400	"请求错误"
// @Failure	401	"token验证失败"
// @Failure	500	"内部错误"
// @Router		/api/srs-hooks/{id} [delete]
func SrsHookDel(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`pls input id`, http.StatusBadRequest, c)
		return
	}
	deleted, err := services.SrsHook.Delete(c, cast.ToUint(id))
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(deleted, c)
}

// @Summary	streams/publish
// @Tags		srs-hooks
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header	string			true	"Bearer 用户令牌"
// @Param		{object}		body	models.SrsHook	true	"body"
// @Success	200
// @Failure	400	{object}	string	"请求错误"
// @Failure	401	{object}	string	"token验证失败"
// @Failure	500	{object}	string	"内部错误"
// @Router		/api/srshooks/streams/publish [post]
func StreamPublish(c *gin.Context) {
	model := &models.SrsHook{}
	err := c.ShouldBindBodyWith(&model, binding.JSON)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	URL, _ := url.Parse(model.Param)
	args := URL.Query()
	if token, ok := args[`token`]; ok {
		event, err := services.Event.GetOneByStream(c, model.Stream)
		if err != nil {
			return
		}
		if token[0] != event.Token {
			//鉴权没通过
			return
		}
		services.Event.ChangeStatus(c, models.EventStatusLiving, event.ID)
		services.SrsHook.Log(c, model)
		if event.ThirdRtmpPush != `` {
			go func() {
				if runtime.GOOS == `linux` {

					// ffmpegPath = `/usr/local/ffmpeg/bin/ffmpeg -re -i  ` + event.RtmpFrom + `/` + event.RtmpFromCode + `  -c copy -f flv ` + event.RtmpTo + `/` + event.RtmpToCode
					// cd := exec.Command("bash", "/c", " start "+ffmpegPath)
					// err := cd.Run()
					// if err != nil {
					// 	fmt.Println(`err===`, err.Error())
					// }

					args := []string{"-re", "-i", event.RtmpPull + `/` + event.RtmpPullCode, "-c", "copy", "-f", "flv", event.ThirdRtmpPush + `/` + event.ThirdRtmpPushCode}

					// 创建 *exec.Cmd
					ffmpegPath := global.Config.GetString(`live.ffmpegpath`)
					cmd := exec.Command(ffmpegPath, args...)

					// 运行 ffmpeg 命令
					if err := cmd.Run(); err != nil {
						fmt.Println(`err===`, err.Error())
					}

				} else {
					ffmpegPath := `C:/Users/simpl/Application/ffmpeg-2023-03-05-git-912ac82a3c-full_build/bin/ffmpeg.exe -re -i  ` + event.RtmpPull + `/` + event.RtmpPullCode + `  -c copy -f flv ` + event.ThirdRtmpPush + `/` + event.ThirdRtmpPushCode
					cd := exec.Command("cmd.exe", "/c", " start "+ffmpegPath)
					err := cd.Run()
					if err != nil {
						fmt.Println(`err===`, err.Error())
					}
				}
			}()
			responseMap := map[string]any{
				`code`: 0,
				`msg`:  `OK`,
			}
			response.OK(responseMap, c)
		}
	}
}

// @Summary	streams/unpublish
// @Tags		srs-hooks
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header		string					true	"Bearer 用户令牌"
// @Param		{object}		body	models.SrsHook	true	"body"
// @Success	200	 "成功"
// @Failure	400				{object}	string					"请求错误"
// @Failure	401				{object}	string					"token验证失败"
// @Failure	500				{object}	string					"内部错误"
// @Router		/api/srshooks/streams/unpublish [post]
func StreamsUnPublish(c *gin.Context) {

	model := &models.SrsHook{}
	err := c.ShouldBindBodyWith(&model, binding.JSON)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	URL, _ := url.Parse(model.Param)
	args := URL.Query()
	if token, ok := args[`token`]; ok {
		event, err := services.Event.GetOneByStream(c, model.Stream)
		if err != nil {
			return
		}
		if token[0] != event.Token {
			//鉴权没通过
			return
		}
		services.Event.ChangeStatus(c, models.EventStatusEnd, event.ID)
		services.SrsHook.Log(c, model)
		responseMap := map[string]any{
			`code`: 0,
			`msg`:  `OK`,
		}
		response.OK(responseMap, c)
	}
}

// @Summary	sessions/play
// @Tags		srs-hooks
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header		string					true	"Bearer 用户令牌"
// @Param		{object}		body	models.SrsHook	true	"body"
// @Success	200 "成功"
// @Failure	400				{object}	string					"请求错误"
// @Failure	401				{object}	string					"token验证失败"
// @Failure	500				{object}	string					"内部错误"
// @Router		/api/srshooks/sessions/play [post]
func SessionsPlay(c *gin.Context) {

	model := &models.SrsHook{}
	err := c.ShouldBindBodyWith(&model, binding.JSON)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	URL, _ := url.Parse(model.Param)
	args := URL.Query()
	if token, ok := args[`token`]; ok {
		event, err := services.Event.GetOneByStream(c, model.Stream)
		if err != nil {
			return
		}
		if token[0] != event.Token {
			//鉴权没通过
			return
		}
		services.SrsHook.Log(c, model)
		responseMap := map[string]any{
			`code`: 0,
			`msg`:  `OK`,
		}
		response.OK(responseMap, c)
	}
}

// @Summary	sessions/stop
// @Tags		srs-hooks
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header		string					true	"Bearer 用户令牌"
// @Param		body			body		models.SrsHook	true	"models.SrsHook"
// @Success	200 "成功"
// @Failure	400				{object}	string					"请求错误"
// @Failure	401				{object}	string					"token验证失败"
// @Failure	500				{object}	string					"内部错误"
// @Router		/api/srshooks/sessions/stop [post]
func SessionsStop(c *gin.Context) {
	model := &models.SrsHook{}
	err := c.ShouldBindBodyWith(&model, binding.JSON)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	URL, _ := url.Parse(model.Param)
	args := URL.Query()
	if token, ok := args[`token`]; ok {
		event, err := services.Event.GetOneByStream(c, model.Stream)
		if err != nil {
			return
		}
		if token[0] != event.Token {
			//鉴权没通过
			return
		}
		services.SrsHook.Log(c, model)
		responseMap := map[string]any{
			`code`: 0,
			`msg`:  `OK`,
		}
		response.OK(responseMap, c)
	}

}

// @Summary	dvrs
// @Tags		srs-hooks
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header		string					true	"Bearer 用户令牌"
// @Param		id				query		string					true	"id"
// @Param		body			body		models.SrsHook	true	"models.SrsHook"
// @Success	200 "成功"
// @Failure	400				{object}	string					"请求错误"
// @Failure	401				{object}	string					"token验证失败"
// @Failure	500				{object}	string					"内部错误"
// @Router		/api/srshooks/dvrs [post]
func Dvrs(c *gin.Context) {
	model := &models.SrsHook{}
	err := c.ShouldBindBodyWith(&model, binding.JSON)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	URL, _ := url.Parse(model.Param)
	args := URL.Query()
	if token, ok := args[`token`]; ok {
		event, err := services.Event.GetOneByStream(c, model.Stream)
		if err != nil {
			return
		}
		if token[0] != event.Token {
			//鉴权没通过
			return
		}
		services.SrsHook.Log(c, model)
		responseMap := map[string]any{
			`code`: 0,
			`msg`:  `OK`,
		}
		response.OK(responseMap, c)
	}

}

// @Summary	hls
// @Tags		srs-hooks
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header		string					true	"Bearer 用户令牌"
// @Param		id				query		string					true	"id"
// @Success	200 "成功"
// @Failure	400				{object}	string					"请求错误"
// @Failure	401				{object}	string					"token验证失败"
// @Failure	500				{object}	string					"内部错误"
// @Router		/api/srshooks/hls [post]
func Hls(c *gin.Context) {
	model := &models.SrsHook{}
	err := c.ShouldBindBodyWith(&model, binding.JSON)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	URL, _ := url.Parse(model.Param)
	args := URL.Query()
	if token, ok := args[`token`]; ok {
		event, err := services.Event.GetOneByStream(c, model.Stream)
		if err != nil {
			return
		}
		if token[0] != event.Token {
			//鉴权没通过
			return
		}
		services.SrsHook.Log(c, model)
		responseMap := map[string]any{
			`code`: 0,
			`msg`:  `OK`,
		}
		response.OK(responseMap, c)
	}

}

// @Summary	hls/:app/:stream/:ts_url:param
// @Tags		srs-hooks
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header		string					true	"Bearer 用户令牌"
// @Param		id				query		string					true	"id"
// @Success	200 "成功"
// @Failure	400				{object}	string					"请求错误"
// @Failure	401				{object}	string					"token验证失败"
// @Failure	500				{object}	string					"内部错误"
// @Router		/api/srshooks/hls/:app/:stream/:ts_url:param [get]
func HlsMore(c *gin.Context) {
	model := &models.SrsHook{}
	err := c.ShouldBindBodyWith(&model, binding.JSON)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	URL, _ := url.Parse(model.Param)
	args := URL.Query()
	if token, ok := args[`token`]; ok {
		event, err := services.Event.GetOneByStream(c, model.Stream)
		if err != nil {
			return
		}
		if token[0] != event.Token {
			//鉴权没通过
			return
		}
		services.SrsHook.Log(c, model)
		responseMap := map[string]any{
			`code`: 0,
			`msg`:  `OK`,
		}
		response.OK(responseMap, c)
	}

}
