package sdk

import (
	"context"

	"github.com/abmpio/abmp/pkg/log"
	pb "github.com/abmpio/aclx-sdk/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type nullableUserServiceClient struct {
}

var _ pb.UserServiceClient = (*nullableUserServiceClient)(nil)

// #region

// #region nullableUserServiceClient Members

func (*nullableUserServiceClient) UserLogin(ctx context.Context, in *pb.UserServiceLoginRequest, opts ...grpc.CallOption) (*pb.UserServiceLoginResponse, error) {
	log.Logger.Warn("nullableUserServiceClient.UserLogin method")
	return &pb.UserServiceLoginResponse{
		Result: &pb.UserLoginResult{},
	}, nil
}

// 创建一个用户
func (*nullableUserServiceClient) CreateUser(ctx context.Context, in *pb.CreateUserRequest, opts ...grpc.CallOption) (*pb.CreateUserResponse, error) {
	log.Logger.Warn("nullableUserServiceClient.CreateUser method")
	return &pb.CreateUserResponse{
		UserInfo: nil,
	}, nil
}

func (*nullableUserServiceClient) UserRefreshToken(ctx context.Context, in *pb.UserServiceRefreshTokenRequest, opts ...grpc.CallOption) (*pb.UserServiceRefreshTokenResponse, error) {
	log.Logger.Warn("nullableUserServiceClient.UserRefreshToken method")
	return &pb.UserServiceRefreshTokenResponse{}, nil
}

func (*nullableUserServiceClient) FindUserByUserId(ctx context.Context, in *pb.FindUserByUserIdRequest, opts ...grpc.CallOption) (*pb.FindUserByUserIdResponse, error) {
	log.Logger.Warn("nullableUserServiceClient.FindUserByUserId method")
	return &pb.FindUserByUserIdResponse{}, nil
}

// 搜索用户id列表
func (*nullableUserServiceClient) FindUserIdList(ctx context.Context, in *pb.FindUserIdListRequest, opts ...grpc.CallOption) (*pb.FindUserIdListResponse, error) {
	log.Logger.Warn("nullableUserServiceClient.FindUserIdList method")
	return &pb.FindUserIdListResponse{}, nil
}

func (*nullableUserServiceClient) FindUserListByPhone(ctx context.Context, in *pb.FindUserListByPhoneRequest, opts ...grpc.CallOption) (*pb.FindUserListByPhoneResponse, error) {
	log.Logger.Warn("nullableUserServiceClient.FindUserListByPhone method")
	return &pb.FindUserListByPhoneResponse{}, nil
}

// 根据角色查找
func (*nullableUserServiceClient) FindUserListByRole(ctx context.Context, in *pb.FindUserListByRoleRequest, opts ...grpc.CallOption) (*pb.FindUserListByRoleResponse, error) {
	log.Logger.Warn("nullableUserServiceClient.FindUserListByRole method")
	return &pb.FindUserListByRoleResponse{}, nil
}

// 切换到个人帐号
func (*nullableUserServiceClient) ChangeToPrivateAccount(ctx context.Context, in *pb.ChangeToPrivateAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	log.Logger.Warn("nullableUserServiceClient.ChangeToPrivateAccount method")
	return &emptypb.Empty{}, nil
}

// 切换到个人帐号
func (*nullableUserServiceClient) ChangeUserCurrentTenantId(ctx context.Context, in *pb.ChangeUserCurrentTenantIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	log.Logger.Warn("nullableUserServiceClient.ChangeUserCurrentTenantId method")
	return &emptypb.Empty{}, nil
}

// 确保用户在指定角色中
func (*nullableUserServiceClient) AddUserToRole(ctx context.Context, in *pb.AddUserToRoleRequest, opts ...grpc.CallOption) (*pb.AddUserToRoleResponse, error) {
	log.Logger.Warn("nullableUserServiceClient.AddUserToRole method")
	return &pb.AddUserToRoleResponse{
		Successful: false,
	}, nil
}

func (*nullableUserServiceClient) FindUserIdAndDisplayNameList(ctx context.Context, in *pb.FindUserIdAndDisplayNameListRequest, opts ...grpc.CallOption) (*pb.FindUserIdAndDisplayNameListResponse, error) {
	log.Logger.Warn("nullableUserServiceClient.FindUserIdAndDisplayNameList method")
	return &pb.FindUserIdAndDisplayNameListResponse{}, nil
}

// #endregion
