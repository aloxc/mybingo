package server

import (
	"sync/atomic"
)
type BinlogSlave struct {
	started *atomic.Bool
	
}
