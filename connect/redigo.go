package connect

import (
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
	"time"
)

func RedisPool() *redis.Pool {
	//最大空闲数
	MaxIdle := viper.GetInt("redis.MaxIdle")
	//最大活跃数
	MaxActive := viper.GetInt("redis.MaxActive")
	//最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
	Timeout := viper.GetDuration("redis.Timeout")
	//network
	network := viper.GetString("redis.network")
	//address
	address := viper.GetString("redis.host") + ":" + viper.GetString("redis.port")
	//密码
	password := viper.GetString("redis.password")
	return &redis.Pool{
		MaxIdle:     MaxIdle,
		MaxActive:   MaxActive,
		IdleTimeout: Timeout * time.Second,
		Dial: func() (redis.Conn, error) {
			//此处对应redis ip及端口号
			conn, err := redis.Dial(network, address)
			if err != nil {
				return nil, err
			}

			if _, err := conn.Do("AUTH", password); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},
	}
}
