package main

import (
	"path"
	"runtime"

	"github.com/urydmi/migorm"
	_ "github.com/urydmi/migorm/example/migrations"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//
// only for run  make migrations command
// uncomment lines for create dbConn and fill connection params for execute other commands
func main() {

	var dbConn *gorm.DB

	//db_user := ""
	//db_pass := ""
	//db_host := ""
	//db_port := ""
	//db_name := ""
	//
	//conStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=true&loc=Local", db_user, db_pass, db_host, db_port, db_name)
	//dbConn, err := gorm.Open("mysql", conStr)
	//
	//if err != nil{
	//	panic(err)
	//}

	migrater := migorm.NewMigrater(dbConn)

	_, file, _, _ := runtime.Caller(0)
	curDir := path.Dir(file)
	migrater.Conf().MigrationsDir = curDir + "/../migrations"

	migorm.Run(migrater)
}
