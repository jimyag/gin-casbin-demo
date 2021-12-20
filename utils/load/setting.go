package load

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type Setting struct {
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
	DbTimeZone string
}

func InitSet() *Setting {
	Set := &Setting{}
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查路径")
	}
	Set.LoadData(file)
	Set.LoadServer(file)
	return Set
}

func (this *Setting) LoadServer(file *ini.File) {
	this.AppMode = file.Section("server").Key("AppMode").String()
	this.HttpPort = file.Section("server").Key("HttpPort").String()
	this.JwtKey = file.Section("server").Key("JwtKey").String()
}

func (this *Setting) LoadData(file *ini.File) {
	this.Db = file.Section("database").Key("Db").String()
	this.DbHost = file.Section("database").Key("DbHost").String()
	this.DbPort = file.Section("database").Key("DbPort").String()
	this.DbUser = file.Section("database").Key("DbUser").String()
	this.DbPassWord = file.Section("database").Key("DbPassWord").String()
	this.DbName = file.Section("database").Key("DbName").String()
	this.DbTimeZone = file.Section("database").Key("DbTimeZone").String()
}
