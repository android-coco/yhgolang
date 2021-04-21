package grpc

import "context"

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) mustEmbedUnimplementedHelloServiceServer() {

}
func (p *HelloServiceImpl) Hello(ctx context.Context, args *String, ) (*String, error) {
	reply := &String{Value: "hello:" + args.GetValue()}
	return reply, nil
}
