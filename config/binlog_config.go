package config

type BinlogServerConfig struct {
	Addr string `json:"addr"`
	// User is for MySQL user.
	User string `json:"user"`
	// Password is for MySQL password.
	Password string `json:"password"`
}
