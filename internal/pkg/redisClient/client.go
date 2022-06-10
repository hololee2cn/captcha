package redisClient

import (
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
)

var (
	rc redis.UniversalClient
)

func Init(addrs []string) {
	if rc == nil {
		rc = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs: addrs,
		})
	}
	if err := rc.Ping().Err(); err != nil {
		panic(err)
	}
	log.Infof("ping redis-cluster success")
}

func Get() redis.UniversalClient {
	return rc
}

func Close() (err error) {
	if rc != nil {
		return rc.Close()
	}
	log.Warnf("nil redis client, no need to close")
	return nil
}
