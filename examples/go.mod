module github.com/hololee2cn/captcha/examples

go 1.16

replace (
	github.com/hololee2cn/captcha => ../
	github.com/hololee2cn/captcha/pkg/grpcIFace => ./../pkg/grpcIFace
)

require (
	github.com/hololee2cn/captcha v0.0.0-00010101000000-000000000000
	github.com/hololee2cn/captcha/pkg/grpcIFace v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.36.1
)
