/*
generated by comer,https://github.com/imoowi/comer
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package models

import (
	"github.com/imoowi/comer/components"
)

type MenuBase struct{
	Name  string     `json:"name" form:"name" gorm:"column:name;type:varchar(30);not null;comment:名" `
}

// Menu表
type Menu struct {
	components.GormModel
	MenuBase
}

// IModel.GetID实现
func (m *Menu) GetID() uint {
	return m.ID
}
func (m *Menu) SetId(id uint) {
	m.ID = id
}

// IModel.TableName实现
func (m *Menu) TableName() string {
	return `menu` + `s`
}
