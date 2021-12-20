package main

import (
	"database/sql"
	"gin-casbin/global"
	"gin-casbin/initiate"
	"gin-casbin/routes"
)

func main() {
	initiate.InitGlobal()
	db, _ := global.GVA_DB.DB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)
	routes.InitRouter()
}
