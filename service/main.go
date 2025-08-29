package main

import (
	"github.com/caozhong11/go-zero-test/client/user"
	"github.com/caozhong11/go-zero-test/service/grpcimpl/userServiceImpl"
	"google.golang.org/grpc"
	"net"
)

func main() {
	//服务注册
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	userImpl := &userServiceImpl.UserServiceServerImpl{}
	user.RegisterUserServiceServer(server, userImpl)
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
