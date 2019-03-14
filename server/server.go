package server

import (
	"context"
	"fmt"
	"github.com/aloxc/mybingo/config"
	"github.com/siddontang/go-log/log"
	"github.com/siddontang/go-mysql/client"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"os"
	"strconv"
	"sync"
)

var configRWLock = new(sync.RWMutex)

type MybingoServer struct {
	conn         *client.Conn
	masterConfig *config.Master
	syncerConfig *config.SyncerConfig
}

func (this *MybingoServer) Start() (err error) {
	this.masterConfig, _ = config.LoadConfig()
	cfg := replication.BinlogSyncerConfig{
		ServerID: 100,
		Flavor:   "mysql",
		Host:     this.masterConfig.Host,
		Port:     this.masterConfig.Port,
		User:     this.masterConfig.User,
		Password: this.masterConfig.Password,
	}
	configRWLock.Lock()
	this.initSyncerConifg()
	log.Infof(this.syncerConfig.String())
	configRWLock.Unlock()
	syncer := replication.NewBinlogSyncer(cfg)
	streamer, _ := syncer.StartSync(*this.syncerConfig.Position)
	for {
		ev, _ := streamer.GetEvent(context.Background())
		// Dump event
		ev.Dump(os.Stdout)
	}

	//streamer, _ := syncer.StartSync(mysql.Position{binlogFile, binlogPos})
	return nil
}

//读取同步位置配置
func (this *MybingoServer) initSyncerConifg() (err error) {
	this.syncerConfig = new(config.SyncerConfig)
	this.syncerConfig.LoadSyncerConfig()
	if this.syncerConfig.Position == nil || len(this.syncerConfig.Position.Name) == 0 {
		//	上次同步位置没有记录，就是说第一次同步
		retryTimes := 1
		for i := 0; i < retryTimes; i++ {
			this.conn, err = client.Connect(this.masterConfig.Host+":"+strconv.Itoa(int(this.masterConfig.Port)),
				this.masterConfig.User, this.masterConfig.Password,
				"mysql")
			if err != nil {
				fmt.Errorf("mybingoServer第[%d]次连接服务器[host:%s,user:%s,password:%s]异常%s",
					i+1, this.masterConfig.Host, this.masterConfig.User, this.masterConfig.Password, err)
			}
			rs, err1 := this.conn.Execute("SHOW MASTER STATUS")
			if err1 != nil {
				fmt.Printf("从数据库读取同步进度异常\n")
				return err1
			}
			name, _ := rs.GetString(0, 0)
			pos, _ := rs.GetInt(0, 1)
			this.syncerConfig.Position = &mysql.Position{
				Name: name,
				Pos:  uint32(pos),
			}
			this.syncerConfig.SaveSyncerConfig()
			break
		}
	}
	return nil
}
func (this *MybingoServer) Stop() error {
	fmt.Println("准备关闭mybingo服务")
	return this.conn.Close()
}
