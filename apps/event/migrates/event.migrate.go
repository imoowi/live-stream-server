/*
generated by comer,https://github.com/imoowi/comer
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package migrate

import (
	"github.com/imoowi/live-stream-server/apps/event/models"
	"github.com/imoowi/comer/components"
	"github.com/imoowi/live-stream-server/global"
)

type EventMigrate struct {
	db *components.MysqlODM
}

func newEventMigrate() *EventMigrate {
	return &EventMigrate{
		db: global.MysqlDb,
	}
}
func init() {
	global.RegisterMigrateContainerProviders(doEventMigrate)
}
func doEventMigrate() {
	r := newEventMigrate()
	r.db.Client.Set("gorm:table_options", "ENGINE=InnoDB,COMMENT='Event表'").AutoMigrate(&models.Event{})
}
