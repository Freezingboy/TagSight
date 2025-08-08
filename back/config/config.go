package config

import (
	"back/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log/slog"
	"time"
)

var (
	ServerOption *ServerSetting
	MysqlOption  *MysqlSetting
	LogOption    *logger.LogSetting
	Jwt          *JWT
)

type ServerSetting struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type MysqlSetting struct {
	UserName     string
	Password     string
	Host         string
	DbName       string
	MaxIdleConns int
	MaxOpenConns int
}

type JWT struct {
	Secret string
	Ttl    int
	Name   string
}

func (m *MysqlSetting) Dsn() string {
	return m.UserName + ":" + m.Password + "@tcp(" + m.Host + ")/" + m.DbName + "?charset=utf8&parseTime=True&loc=Local"
}

func ReadConfigFile(path string) error {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		err := ReadSection("log", &LogOption)
		if err != nil {
			slog.Error("read log section error", "err", err)
		}
		if logger.GetLogLevel(LogOption.Level) != logger.LogLevel.Level() {
			logger.SetLevel(LogOption.Level)
		}
	})
	return nil
}

// 分段读取
func ReadSection(key string, v any) error {
	return viper.UnmarshalKey(key, v)
}

func InitConfig(path string) {
	if err := ReadConfigFile(path); err != nil {
		panic(err)
	}

	err := ReadSection("server", &ServerOption)
	if err != nil {
		panic(err)
	}

	err = ReadSection("mysql", &MysqlOption)
	if err != nil {
		panic(err)
	}

	err = ReadSection("log", &LogOption)
	if err != nil {
		panic(err)
	}

	err = ReadSection("Jwt", &Jwt)
	if err != nil {
		panic(err)
	}
}
