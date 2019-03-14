package main

import (
	"fmt"
	"github.com/siddontang/go-mysql/client"
)

func main() {
	fmt.Println("mybingo项目启动")
	conn, err := client.Connect("localhost:33066", "root", "123456", "mysql")
	if err == nil {
		fmt.Println("连接上了")
		fmt.Println(conn)
	} else {
		fmt.Println("没连接上")
		fmt.Println(err)
	}
	rs, err := conn.Execute("show master status;")
	fmt.Println(rs)
	name, _ := rs.GetString(0, 0)
	position, _ := rs.GetInt(0, 1)
	fmt.Printf("binlog name [  %s ]\n", name)
	fmt.Printf("position [  %d ]\n", position)
	fmt.Println("aaaaaaaaaaaa")
}
