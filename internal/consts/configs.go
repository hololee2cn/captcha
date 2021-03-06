package consts

import (
	"strings"

	"github.com/hololee2cn/captcha/internal/pkg/config"
)

var (
	// 验证码有效时间, 单位秒
	// CaptchaMaxAge = AppConfig.DefaultInt64("captcha_max_age", 180)

	RPCAddr = config.DefaultString("rpc_addr", ":50051")

	LogLevel = config.DefaultInt("log_level", 4)

	RedisAddrs []string
)

func Init() {
	redisAddrs := config.MustString("redis_addrs")
	if len(redisAddrs) > 0 {
		RedisAddrs = strings.Split(redisAddrs, ",")
	}
}
