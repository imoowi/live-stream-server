/*
generated by comer,https://github.com/imoowi/comer
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package models

import (
	"github.com/imoowi/comer/components"
)

type SrsHookBase struct{
	Name  string     `json:"name" form:"name" gorm:"column:name;type:varchar(30);not null;comment:名" `
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
