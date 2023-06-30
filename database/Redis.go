package database

import (
	"dj-lets-go/settings"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Host     string
	Port     string
	Database uint64
	Password string
}

var redisIns *redis.Client

func NewRedis(database uint64) *redis.Client {
	return (&Redis{}).getConn(database)
}

func (receiver Redis) getConn(database uint64) *redis.Client {
	config := settings.NewSetting()

	receiver.Host = config.DB.Section("redis").Key("host").MustString("127.0.0.1")
	receiver.Port = config.DB.Section("redis").Key("port").MustString("6379")
	receiver.Database = config.DB.Section("redis").Key("database").MustUint64(0)
	receiver.Password = config.DB.Section("redis").Key("password").String()
	if database == 0 {
		database = receiver.Database
	}

	if redisIns == nil {
		redisIns = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", receiver.Host, receiver.Port),
			Password: receiver.Password, // no password set
			DB:       int(database),     // use default DB
		})
	}

	return redisIns
}
