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

var UserLog *UserLogService

type UserLogService struct {
	impl.Service
}

func NewUserLogService(r *repos.UserLogRepo) *UserLogService {
	return &UserLogService{
		Service: *impl.NewService(r),
	}
}

func init() {
	RegisterServices(func() {
		UserLog = NewUserLogService(repos.UserLog)
		var mt interfaces.IModel = &models.UserLog{}
		UserLog.MT = &mt
	})
}
