/*
generated by comer,https://github.com/imoowi/comer
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package models

import (
	"github.com/imoowi/comer/components"
)

type HooksAction string

const (
	HooksActionPublish   HooksAction = `on_publish`   //推流
	HooksActionUnPublish HooksAction = `on_unpublish` //拉流
	HooksActionPlay      HooksAction = `on_play`      //播放
	HooksActionStop      HooksAction = `on_stop`      //停止播放
	HooksActionHls       HooksAction = `on_hls`       // hls
	HooksActionHlsNotify HooksAction = `on_hls_notify`
	HooksActionDvr       HooksAction = `on_dvr` //录像
)

type SrsHookBase struct {
	Name     string      `json:"name" form:"name" gorm:"column:name;type:varchar(30);not null;comment:名" `
	Action   HooksAction `json:"action" gorm:"column:action;type:varchar(30)"`
	ServerId string      `json:"server_id" gorm:"column:server_id;type:varchar(255)"`
	Duration float32     `json:"duration" gorm:"column:duration;type:int(6)"`
	Stream   string      `json:"stream" gorm:"column:stream;type:varchar(255)"`
	Vhost    string      `json:"vhost" gorm:"column:vhost;type:varchar(255)"`
	M3u8     string      `json:"m3u8" gorm:"column:m3u8;type:varchar(255)"`
	M3u8Url  string      `json:"m3u8_url" gorm:"column:m3u8_url;type:varchar(255)"`
	App      string      `json:"app" gorm:"column:app;type:varchar(255)"`
	Param    string      `json:"param" gorm:"column:param;type:varchar(255)"`
	Cwd      string      `json:"cwd" gorm:"column:cwd;type:varchar(255)"`
	File     string      `json:"file" gorm:"column:file;type:varchar(255)"`
	ClientId string      `json:"client_id" gorm:"column:client_id;type:varchar(255)"`
	Ip       string      `json:"ip" gorm:"column:ip;type:varchar(255)"`
	Url      string      `json:"url" gorm:"column:url;type:varchar(255)"`
}

// SrsHook表
type SrsHook struct {
	components.GormModel
	SrsHookBase
}

// IModel.GetID实现
func (m *SrsHook) GetID() uint {
	return m.ID
}
func (m *SrsHook) SetId(id uint) {
	m.ID = id
}

// IModel.TableName实现
func (m *SrsHook) TableName() string {
	return `srs_hook` + `s`
}
