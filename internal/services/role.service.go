/*
generated by comer,https://github.com/imoowi/comer
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package services

import (
	"github.com/imoowi/live-stream-server/internal/models"
	"github.com/imoowi/live-stream-server/internal/repos"
	"github.com/imoowi/comer/interfaces"
	"github.com/imoowi/comer/interfaces/impl"
)

var Role *RoleService

type RoleService struct {
	impl.Service
}

func init() {
	RegisterServices(func() {
		Role = NewRoleService(repos.Role)
		var mt interfaces.IModel = &models.Role{}
		Role.MT = &mt
	})
}
func NewRoleService(r *repos.RoleRepo) *RoleService {
	return &RoleService{
		Service: *impl.NewService(r),
	}
}