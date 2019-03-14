package main

import (
	"fmt"
	"github.com/aloxc/mybingo/server"
	"github.com/siddontang/go-log/log"
	"os"
	"os/signal"
)

func main() {
	log.Infof("mybingo项目启动")
	//mysqlConfig, configError := config.LoadConfig()
	//if configError != nil {
	//	log.Fatalf("读取mybingo配置错误", mysqlConfig, configError)
	//}
	//mybingoConfig := config.ReadParam()
	//config.PrintConfig(mybingoConfig)
	//log.Println(mysqlConfig)
	mybingoServer := new(server.MybingoServer)
	mybingoServer.Start()

	ctrlC := make(chan os.Signal, 1)
	signal.Notify(ctrlC, os.Interrupt, os.Kill)

	quit := <-ctrlC
	go func() {
		for _ = range ctrlC {
			fmt.Println("系统即将退出", quit)
			//停止mybingo服务，要回收资源
			mybingoServer.Stop()
		}
	}()
}
