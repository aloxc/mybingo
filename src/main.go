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
	

	ctrlC := make(chan os.Signal,1)
	signal.Notify(ctrlC,os.Interrupt,os.Kill)

	quit := <-ctrlC
	go func() {
		for _ = range ctrlC{
			fmt.Println("系统即将退出",s)
			//停止mybingo服务，要回收资源
		}
	}()





}

