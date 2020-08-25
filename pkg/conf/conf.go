package conf

import (
	"flag"
	"io/ioutil"

	"rtmp-recorder/pkg/utils"
)

type Config struct {
	Port int

	DbHost string
	DbPort int
	DbUser string
	DbPwd  string
	DbName string

	MasterId string
	Token    string
}

func ParseConfig() *Config {
	c := &Config{}

	flag.IntVar(&c.Port, "port", 5000, "server port")
	// flag.StringVar(&c.DbHost, "db-host", "127.0.0.1", "database host")
	// flag.IntVar(&c.DbPort, "db-port", 3306, "database port")
	// flag.StringVar(&c.DbUser, "db-user", "root", "database user")
	// flag.StringVar(&c.DbPwd, "db-pwd", "root", "database password")
	// flag.StringVar(&c.DbName, "db-name", "artisano", "database name")
	flag.StringVar(&c.Token, "token", "12345678", "token")

	flag.Parse()

	// 获取 Master ID
	var masterId string
	masterIdFile := ".master"
	if utils.Exists(masterIdFile) {
		if buf, err := ioutil.ReadFile(masterIdFile); err == nil {
			masterId = string(buf)
		}
	} else {
		masterId = utils.NewId()
		ioutil.WriteFile(masterIdFile, []byte(masterId), 0666)
	}
	c.MasterId = masterId

	return c
}
