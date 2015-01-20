package bin

import (
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
)

var mem *memcache.Client

func Memcache() *memcache.Client {
	if mem == nil {
		mem = memcache.New(strings.Split(Config.String("MEMCACHE_SERVERS"), ";")...)
	}

	return mem
}


/*
item := memcache.Item{
		Key:        Config.MemcachePrefix + "_filetoken_" + token,
		Value:      td,
		Expiration: Config.TokenTTL,
	}

	err = bin.Memcache().Add(&item)
	if err != nil {
		panic(err)
	}

	

*/