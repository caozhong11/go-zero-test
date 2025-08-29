package userServiceImpl

import (
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

import "github.com/caozhong11/go-zero-test/client/user"

type UserServiceServerImpl struct {
	user.UnimplementedUserServiceServer
}

func (u UserServiceServerImpl) GetUserList(context context.Context, request *user.UserRequest) (*user.UserResponse, error) {
	id := request.Id
	response := &user.UserResponse{}
	response.Id = id
	response.EventTime = timestamppb.Now()
	return response, nil
}
