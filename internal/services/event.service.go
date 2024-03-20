/*
generated by comer,https://github.com/imoowi/comer
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package services

import (
	"github.com/imoowi/comer/interfaces/impl"
	"github.com/imoowi/comer/interfaces"
	"github.com/imoowi/live-stream-server/internal/models"
	"github.com/imoowi/live-stream-server/internal/repos"
)

var Event *EventService

type EventService struct {
	impl.Service
}

func NewEventService(r *repos.EventRepo) *EventService {
	return &EventService{
		Service: *impl.NewService(r),
	}
}

func init() {
	RegisterServices(func() {
		Event = NewEventService(repos.Event)
		var mt interfaces.IModel = &models.Event{}
		Event.MT = &mt
	})
}
