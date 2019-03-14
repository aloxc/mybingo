package config

import (
	"encoding/json"
	"fmt"
	"github.com/siddontang/go-log/log"
	"github.com/siddontang/go-mysql/mysql"
	"io/ioutil"
	"os"
)

//同步配置
type SyncerConfig struct {
	Position *mysql.Position
}

func (this *SyncerConfig) String() string {
	bs, _ := json.Marshal(this)
	return "SyncerConfig = " + string(bs)
}

//读取同步位置
func (this *SyncerConfig) LoadSyncerConfig() {
	log.Infof("准备读取复制进度")
	_, err := os.Stat(positionFile)
	if err == nil { //文件存在
		file, err := os.Open(positionFile)
		if err == nil {
			bs, _ := ioutil.ReadAll(file)
			if len(bs) != 0 {
				err := json.Unmarshal(bs, &this.Position)
				fmt.Println(err)
			}
		}
		defer file.Close()
	}

	if os.IsNotExist(err) { //文件不存在
		crefile, _ := os.Create(positionFile)
		defer crefile.Close()
	}
}

//保存同步位置
func (this *SyncerConfig) SaveSyncerConfig() {
	file, err := os.Open(positionFile)
	log.Infof("准备保存复制进度" + positionFile)
	if err == nil {
		bs, _ := json.Marshal(this.Position)
		ioutil.WriteFile(positionFile, bs, 0644)
	}
	defer file.Close()
}
