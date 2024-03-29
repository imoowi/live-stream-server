/*
generated by comer,https://github.com/imoowi/comer
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package models

import (
	"github.com/imoowi/comer/components"
)

type RetweetStreamBase struct{
	Name  string     `json:"name" form:"name" gorm:"column:name;type:varchar(30);not null;comment:名" `
}

// RetweetStream表
type RetweetStream struct {
	components.GormModel
	RetweetStreamBase
}

// IModel.GetID实现
func (m *RetweetStream) GetID() uint {
	return m.ID
}
func (m *RetweetStream) SetId(id uint) {
	m.ID = id
}

// IModel.TableName实现
func (m *RetweetStream) TableName() string {
	return `retweet_stream` + `s`
}
