package main

import (
	"github.com/hololee2cn/captcha/internal/biz"
	"github.com/hololee2cn/captcha/internal/consts"

	"github.com/hololee2cn/pkg/extra"
	log "github.com/sirupsen/logrus"
)

func main() {
	consts.Init()
	extra.Default(log.Level(consts.LogLevel))
	log.Info("captcha starting...")
	biz.Start()
}
