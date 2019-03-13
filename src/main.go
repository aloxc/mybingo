package main

import (
	"fmt"
	"log"
	"mybingo1/config"
	"mybingo1/dbsync"
	_ "os"
)

func main() {
	fmt.Println("mybingo项目启动")
	mysqlConfig, configError := config.LoadConfig()
	if configError != nil {
		log.Fatalf("读取mybingo配置错误", mysqlConfig, configError)
	}
	mybingoConfig := config.ReadParam()
	config.PrintConfig(mybingoConfig)
	log.Println(mysqlConfig)
	dbsync.DbSync(mysqlConfig)
}
