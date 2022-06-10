package biz

import (
	"context"
	"fmt"

	"github.com/hololee2cn/captcha/internal/biz/rpc/server"
	"github.com/hololee2cn/captcha/internal/biz/service"
	"github.com/hololee2cn/captcha/internal/biz/store"
	"github.com/hololee2cn/captcha/internal/consts"
	me "github.com/hololee2cn/captcha/internal/errors"
	"github.com/hololee2cn/captcha/internal/pkg/redisClient"

	grpcMW "github.com/grpc-ecosystem/go-grpc-middleware"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	captchaSvc = service.NewDefaultCaptchaSvc(store.NewRedisStore(consts.CaptchaDefaultMaxAge))
)

func Start() {
	redisClient.Init(consts.RedisAddrs)

	captchaSvcServer := server.NewCaptchaSvcServer(captchaSvc)

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpcMW.ChainUnaryServer(
			getFullMethodName,
			logRequest,
		)),
	}
	rs := server.NewRpcServer(consts.RPCAddr, captchaSvcServer, opts...)
	defer rs.Stop()
	rs.Start()
}

func getFullMethodName(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fullMethodName := info.FullMethod
	if md, ok := metadata.FromIncomingContext(ctx); !ok {
		log.Errorf("no metadata in context")
	} else {
		md.Append("full_method_name", fullMethodName)
	}
	return handler(ctx, req)
}

func logRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if md, ok := metadata.FromIncomingContext(ctx); !ok {
		log.Errorf("no metadata in context. info: %+v", *info)
	} else {
		log.Infof("metadata: %+v", md)
		if slice := md.Get(consts.KeyTraceID); len(slice) == 0 {
			err := me.NoTraceID(fmt.Sprintf("in metadata: %+v", md))
			log.Error(err)
			return nil, err
		}
	}
	return handler(ctx, req)
}
