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

type UserRoleMigrate struct {
	db *components.MysqlODM
}

func newUserRoleMigrate() *UserRoleMigrate {
	return &UserRoleMigrate{
		db: global.MysqlDb,
	}
}
func init() {
	global.RegisterMigrateContainerProviders(doUserRoleMigrate)
}
func doUserRoleMigrate() {
	r := newUserRoleMigrate()
	r.db.Client.Set("gorm:table_options", "ENGINE=InnoDB,COMMENT='用户角色关系表'").AutoMigrate(&models.UserRole{})
}
