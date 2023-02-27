package main

import (
	"bubble/dao"
	"bubble/model"
	"bubble/router"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		log.Panicln(err.Error())
	}

	defer dao.Close()

	//建立表关联 模型绑定
	dao.DB.AutoMigrate(&model.Todo{})

	r := router.SetupRouter()
	r.Run(":8088")

}
