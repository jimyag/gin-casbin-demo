package initiate

import (
	"fmt"
	"gin-casbin/global"
	"gin-casbin/model"
	"gin-casbin/utils/load"
)

func InitGlobal() {
	global.GVA_Set = load.InitSet()
	fmt.Println(global.GVA_Set)
	global.GVA_DB = model.InitDB(global.GVA_Set)
	casbin := &model.Casbin{}
	global.GVA_CASBIN = casbin.Casbin()

	if global.GVA_DB == nil {
		fmt.Println("数据库错误")
	}

}
