package dbsync

import (
	_ "fmt"
	_ "log"
	"mybingo1/config"

	"github.com/siddontang/go-mysql/replication"
)

func DbSync(mysqlConfig *config.MySqlConfig) {

	cfg := replication.BinlogSyncerConfig{
		ServerID: 100,
		Flavor:   "mysql",
		Host:     mysqlConfig.Host,
		Port:     mysqlConfig.Port,
		User:     mysqlConfig.User,
		Password: mysqlConfig.Password,
	}
	replication.NewBinlogSyncer(cfg)
}
