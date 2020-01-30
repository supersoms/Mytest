package model

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

//程序初始化数据
func init() {
	//初始化数据库
	//RegisterDB()
}

//注册数据库
func RegisterDB() {
	models := []interface{}{
		NewBlokc(),
		NewRawTransaction(),
		NewLottery(),
	}
	///映射Models到数据库
	orm.RegisterModelWithPrefix(beego.AppConfig.DefaultString("db::prefix", "cxc_"), models...)

	//掉用数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	//数据库连接参数
	dbUser := beego.AppConfig.String("db::user")
	dbPassword := beego.AppConfig.String("db::password")
	dbDatabases := beego.AppConfig.String("db::database")
	dbCharset := beego.AppConfig.String("db::charset")
	dbHost := beego.AppConfig.String("db::host")
	dbPort := beego.AppConfig.String("db::port")

	dblink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", dbUser, dbPassword, dbHost, dbPort, dbDatabases, dbCharset)
	//连接数据 ( 默认参数 ，mysql数据库 ，"数据库的用户名 ：数据库密码@tcp("+数据库地址+":"+数据库端口+")/库名？格式",默认参数）
	// create database if not exists bit default charset utf8 collate utf8_general_ci;
	err := orm.RegisterDataBase("default", "mysql", dblink, 60)
	if err != nil {
		log.Panic("数据库连接失败！", err.Error())
	}

	//创建数据库表
	//第一个是别名
	//第二个是是否强制替换模块   如果表变更就将false 换成true 之后再换回来表就便更好来了
	//第三个参数是如果没有则同步或创建
	orm.RunSyncdb("default", false, false)
}

//获取带表前缀的数据表
func getTable(table string) string {
	prefix := beego.AppConfig.DefaultString("db::prefix", "bt_")
	if !strings.HasPrefix(table, prefix) {
		table = prefix + table
	}
	return table
}