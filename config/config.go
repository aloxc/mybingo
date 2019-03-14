package config

import (
	"encoding/json"
	"flag"
	"fmt"
	lsconf "github.com/larspensjo/config"
	"log"
	"os"
	"strings"
)

const (
	configFile   = "./config/config.ini"
	master       = "master"
	serverId     = "serverId"
	host         = "host"
	port         = "port"
	user         = "user"
	password     = "password"
	positionFile = "./data/position.log"

	print   = "print"
	version = "version"
)

type Master struct {
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

func (this *Master) String() string {
	bs, _ := json.Marshal(this)
	return "Master = " + string(bs)
}
func LoadConfig() (*Master, error) {

	file := flag.String("config", configFile, "mybingo配置文件")
	flag.Parse()
	cfg, err := lsconf.ReadDefault(*file)
	if err != nil {
		log.Fatalf("找不到mybingo配置文件", *file, err)
	}
	mysqlConfig := new(Master)
	if cfg.HasSection(master) {
		_, err := cfg.SectionOptions(master)
		if err == nil {
			servId, _ := cfg.Int(master, serverId)
			mysqlConfig.ServerId = uint32(servId)
			mysqlConfig.Host, _ = cfg.String(master, host)
			pot, _ := cfg.Int(master, port)
			mysqlConfig.Port = uint16(pot)
			mysqlConfig.User, _ = cfg.String(master, user)
			mysqlConfig.Password, _ = cfg.String(master, password)

		}
	}
	return mysqlConfig, nil
}
func ReadParam() *MyBingoConfig {
	args := os.Args
	config := MyBingoConfig{}
	//fmt.Println("参数个数",len(args))
	//for v,i := range args{
	//	fmt.Printf("参数[%d]=[%s]\n",i,v)
	//}
	if args != nil && len(args) > 1 {
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
