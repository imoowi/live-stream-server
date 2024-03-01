/*
generated by comer,https://github.com/imoowi/comer
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package migrate

import (
	"github.com/imoowi/live-stream-server/apps/user/models" 
	"github.com/imoowi/comer/components"
	"github.com/imoowi/live-stream-server/global"
)

type UserLogMigrate struct {
	db *components.MysqlODM
}

func newUserLogMigrate() *UserLogMigrate {
	return &UserLogMigrate{
		db: global.MysqlDb,
	}
}
func init() {
	global.RegisterMigrateContainerProviders(doUserLogMigrate)
}
func doUserLogMigrate() {
	r := newUserLogMigrate()
	r.db.Client.Set("gorm:table_options", "ENGINE=InnoDB,COMMENT='用户行为记录表'").AutoMigrate(&models.UserLog{})
}
