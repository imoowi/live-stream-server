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

type SrshookMigrate struct {
	db *components.MysqlODM
}

func newSrshookMigrate() *SrshookMigrate {
	return &SrshookMigrate{
		db: global.MysqlDb,
	}
}
func init() {
	global.RegisterMigrateContainerProviders(doSrshookMigrate)
}
func doSrshookMigrate() {
	r := newSrshookMigrate()
	r.db.Client.Set("gorm:table_options", "ENGINE=InnoDB,COMMENT='Srshook表'").AutoMigrate(&models.Srshook{})
}
