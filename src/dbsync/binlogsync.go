package dbsync

import (
	"github.com/siddontang/go-mysql/replication"
)
func Dbsync() {}
	cfg := replication.BinlogSyncerConfig {
	ServerID: 100,
	Flavor:   "mysql",
	Host:     "127.0.0.1",
	Port:     3306,
	User:     "root",
	Password: "",
	}
	syncer := replication.NewBinlogSyncer(cfg)
}
