package server

import (
	"context"
	"fmt"
	"github.com/aloxc/mybingo/config"
	"github.com/siddontang/go-log/log"
	"github.com/siddontang/go-mysql/client"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var configRWLock = new(sync.RWMutex)

type MybingoServer struct {
	conn         *client.Conn
	config       *config.Config
	syncerConfig *config.SyncerConfig
}

func (this *MybingoServer) StartSync(group *sync.WaitGroup) (err error) {
	this.config, _ = config.LoadConfig()

	cfg := replication.BinlogSyncerConfig{
		ServerID: 100,
		Flavor:   "mysql",
		Host:     this.config.Master.Host,
		Port:     this.config.Master.Port,
		User:     this.config.Master.User,
		Password: this.config.Master.Password,
	}
	configRWLock.Lock()
	this.initSyncerConifg()
	log.Infof(this.syncerConfig.String())
	configRWLock.Unlock()
	group.Done()
	syncer := replication.NewBinlogSyncer(cfg)
	streamer, _ := syncer.StartSync(*this.syncerConfig.Position)
	for {
		ev, _ := streamer.GetEvent(context.Background())
		// Dump event

		ev.Dump(os.Stdout)
		switch ev.Header.EventType {
		case replication.XID_EVENT:
			log.Infof("position = %d", ev.Header.LogPos)

		}
	}
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
			this.conn, err = client.Connect(this.config.Master.Host+":"+strconv.Itoa(int(this.config.Master.Port)),
				this.config.Master.User, this.config.Master.Password,
				"mysql")
			if err != nil {
				fmt.Errorf("mybingoServer第[%d]次连接服务器[host:%s,user:%s,password:%s]异常%s",
					i+1, this.config.Master.Host, this.config.Master.User, this.config.Master.Password, err)
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
func (this *MybingoServer) StopSync() error {
	fmt.Println("准备关闭mybingo服务")
	return this.conn.Close()
}
func (this *MybingoServer) StartHttp() {
	log.Infof("准备启动http管理接口")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != this.config.Manage.Url {
			w.WriteHeader(404)
			w.Write([]byte("<h1>404，找不到你要的内容，要不你来实现一个怎样</h1>"))
		} else {
			w.Write([]byte("http服务器已经启动起来了"))
		}
	})

	http.ListenAndServe(":"+strconv.Itoa(int(this.config.Manage.Port)), nil)
	log.Infof("http管理接口服务启动起来了，端口%d", this.config.Manage.Port)
}
func (this *MybingoServer) StopHttp() {

}
