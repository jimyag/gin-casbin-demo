package global

import (
	"gin-casbin/utils/load"
	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

var (
	GVA_Set    *load.Setting
	GVA_DB     *gorm.DB
	GVA_CASBIN *casbin.SyncedEnforcer
)
