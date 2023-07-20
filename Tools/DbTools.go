package Tools

import (
	"github.com/glebarez/sqlite"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _db *gorm.DB //全局的db对象，我们执行数据库操作主要通过他实现。

func init() { //包初始化函数，golang特性，每个包初始化的时候会自动执行init函数，这里用来初始化gorm。
	/*
	*  dsn配置
	 */
	/*
	 *  注意：声明err变量，下面不能使用:=赋值运算符，否则_db变量会当成局部变量，导致外部无法访问_db变量
	 */
	var err error
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	//dsn := "host=192.168.10.200 user=postgres password=lcl930108 dbname=MyService port=15432 sslmode=disable TimeZone=Asia/Shanghai"
	//postgres.Open(dsn)
	_db, err = gorm.Open(sqlite.Open("wol.db"), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	sqlDB, _ := _db.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

func GetDB() *gorm.DB {
	return _db
}
