/*
generated by comer,https://github.com/imoowi/comer
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package models

import "github.com/imoowi/comer/components"

// 用户角色表
type UserRole struct {
	components.GormModel
	UserID uint `json:"user_id" form:"user_id" gorm:"column:user_id;not null;comment:用户id;uniqueIndex:idx_user_role_rel" binding:"required"`
	RoleId uint `json:"role_id" form:"role_id" gorm:"column:role_id;not null;comment:角色id;uniqueIndex:idx_user_role_rel" binding:"required"`
}

type UserRoleAdd struct {
	UserID uint `json:"user_id" form:"user_id"  binding:"required"`
	RoleId uint `json:"role_id" form:"role_id"  binding:"required"`
}

// IModel.GetID实现
func (m *UserRole) GetID() uint {
	return m.ID
}
func (m *UserRole) SetId(id uint) {
	m.ID = id
}

// IModel.TableName实现
func (m *UserRole) TableName() string {
	return `user_role` + `s`
}
