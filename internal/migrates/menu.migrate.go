/*
generated by comer,https://github.com/imoowi/comer
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package migrates

import (
	"github.com/imoowi/live-stream-server/internal/global"
	"github.com/imoowi/live-stream-server/internal/models"
)

func init() {
	RegisterMigrate(doMenuMigrate)
}
func doMenuMigrate() {
	global.MysqlDb.Client.Set("gorm:table_options", "ENGINE=InnoDB,COMMENT='menu表'").AutoMigrate(&models.Menu{})
}
