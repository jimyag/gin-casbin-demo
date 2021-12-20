package model

import (
	"gin-casbin/global"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)
import "github.com/casbin/casbin/v2"
import "sync"

type Casbin struct {
}

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func (casbin1 *Casbin) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.GVA_DB)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer("config/model.conf", a)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
