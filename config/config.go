package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go/build"
	"os"
	"strings"
)

const (
	True  = "true"
	False = "false"
)

type (
	MysqlDB struct {
		Host     string
		Port     int
		User     string
		Password string
		DataBase string
	}
)

var (
	VeigarRoot string
	DevMode    bool
	Mysql      MysqlDB
)

func init() {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		goPath = build.Default.GOPATH
	}
	VeigarRoot = goPath + "/src/github.com/wubingwei/veigar"
	// develop mode for local debug
	DevMode = os.Getenv("DEV_MODE") == True

	conf := viper.New()
	conf.SetConfigName("veigar")
	if DevMode {
		conf.SetConfigName("veigar.local")
	}
	conf.SetConfigType("yml")
	conf.AddConfigPath("/etc/veigar")
	conf.AddConfigPath(VeigarRoot)
	if err := conf.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.Error("veigar config file not found: ", err.Error())
		} else {
			logrus.Error(err.Error())
		}
	}

	conf.SetEnvPrefix("VEIGAR")
	conf.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	conf.AutomaticEnv()

	conf.SetDefault("PROJECT", "veigar")
	conf.SetDefault("CODER", "wubingwei")
	conf.SetDefault("LOVER", "sunya")

	Mysql.Host = conf.GetString("mysql.host")
	Mysql.Port = conf.GetInt("mysql.port")
	Mysql.User = conf.GetString("mysql.user")
	Mysql.Password = conf.GetString("mysql.password")
	Mysql.DataBase = conf.GetString("mysql.database")
}
