/*
generated by comer,https://github.com/imoowi/comer
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package services

import (
	"github.com/imoowi/comer/interfaces/impl"
	"github.com/imoowi/live-stream-server/internal/models"
	"github.com/imoowi/live-stream-server/internal/repos"
)

var Role *RoleService

type RoleService struct {
	impl.Service[*models.Role]
}

func init() {
	RegisterServices(func() {
		Role = NewRoleService(repos.Role)
	})
}
func NewRoleService(r *repos.RoleRepo) *RoleService {
	return &RoleService{
		Service: *impl.NewService[*models.Role](r),
	}
}
