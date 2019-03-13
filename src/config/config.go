package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	lsconf "github.com/larspensjo/config"
)

const (
	configFile = "./config/config.ini"
	mysql      = "mysql"
	serverId   = "serverId"
	host       = "host"
	port       = "port"
	user       = "user"
	password   = "password"

	print   = "print"
	version = "version"
)

type MySqlConfig struct {
	ServerId uint32
	Host     string
	Port     uint16
	User     string
	Password string
}
type MyBingoConfig struct {
	Print   bool
	Version bool
}

func LoadConfig() (*MySqlConfig, error) {

	file := flag.String("config", configFile, "mybingo配置文件")
	flag.Parse()
	cfg, err := lsconf.ReadDefault(*file)
	if err != nil {
		log.Fatalf("找不到mybingo配置文件", *file, err)
	}
	mysqlConfig := MySqlConfig{}
	if cfg.HasSection(mysql) {
		_, err := cfg.SectionOptions(mysql)
		if err == nil {
			servId, _ := cfg.Int(mysql, serverId)
			mysqlConfig.ServerId = uint32(servId)
			mysqlConfig.Host, _ = cfg.String(mysql, host)
			pot, _ := cfg.Int(mysql, port)
			mysqlConfig.Port = uint16(pot)
			mysqlConfig.User, _ = cfg.String(mysql, user)
			mysqlConfig.Password, _ = cfg.String(mysql, password)

		}
	}
	return &mysqlConfig, nil
}
func ReadParam() *MyBingoConfig {
	args := os.Args
	config := MyBingoConfig{}
	if args != nil && len(args) > 0 {
		argstr := args[1]
		fmt.Println("mybingo运行参数如下")
		params := strings.Split(argstr, "&")
		for _, v := range params {
			fmt.Println("\t", v)
			// kv := strings.Split(v,"=")
			// switch kv[0]{
			// 	case "print":
			// 	config.Print = bool(kv[1])
			// 	case
			// }
		}
	}
	return &config
}
func PrintConfig(config *MyBingoConfig) {
	if config.Print {
		fmt.Println(config.Version)
	}
}
