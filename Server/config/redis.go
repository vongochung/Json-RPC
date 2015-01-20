package bin

import (
	"github.com/xuyu/goredis"
)

var redis *goredis.Redis

func Redis() *goredis.Redis {
	var err error

	if redis == nil {
		redis, err = goredis.Dial(&goredis.DialConfig{
			Address:  Config.String("REDIS_SERVER"),
			Database: Config.Int("REDIS_DATABASE"),
		})
		if err != nil {
			panic(err)
		}
	}

	return redis
}

/*
v, err := r.HGet(Config.String("BANDWIDTH_CACHE_NAME"), u.String())
if err != nil {
	panic(err)
}
 r.HMSet(bin.Config.String("BANDWIDTH_CACHE_NAME"), rmap)

*/
