package sdk

import (
	"context"

	"github.com/abmpio/abmp/pkg/log"
	pb "github.com/abmpio/aclx-sdk/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type nullableTenantServiceClient struct {
}

var _ pb.TenantServiceClient = (*nullableTenantServiceClient)(nil)

// #region

// #region TenantServiceClient Members

func (*nullableTenantServiceClient) FindOneTenantByTenantId(ctx context.Context, in *pb.FindOneTenantByTenantIdRequest, opts ...grpc.CallOption) (*pb.FindOneTenantByTenantIdResponse, error) {
	log.Logger.Warn("nullableTenantServiceClient.FindOneTenantByTenantId method")
	return &pb.FindOneTenantByTenantIdResponse{}, nil
}

// 确保用户在指定租户角色中
func (*nullableTenantServiceClient) AddUserToTenantRole(ctx context.Context, in *pb.AddUserToTenantRoleRequest, opts ...grpc.CallOption) (*pb.AddUserToTenantRoleResponse, error) {
	log.Logger.Warn("nullableTenantServiceClient.AddUserToTenantRole method")
	return &pb.AddUserToTenantRoleResponse{}, nil
}

func (*nullableTenantServiceClient) RemoveTenantUserFromTenant(ctx context.Context, in *pb.RemoveTenantUserFromTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	log.Logger.Warn("nullableTenantServiceClient.RemoveTenantUserFromTenant method")
	return &emptypb.Empty{}, nil
}

// 将用户id增加到租户中
func (*nullableTenantServiceClient) AddUserToTenant(ctx context.Context, in *pb.AddUserToTenantRequest, opts ...grpc.CallOption) (*pb.AddUserToTenantResponse, error) {
	log.Logger.Warn("nullableTenantServiceClient.AddUserToTenant method")
	return &pb.AddUserToTenantResponse{}, nil
}

// 查询指定的用户id的所有tenantId列表
func (*nullableTenantServiceClient) FindTenantIdListByUserId(ctx context.Context, in *pb.FindTenantIdListByUserIdRequest, opts ...grpc.CallOption) (*pb.FindTenantIdListByUserIdResponse, error) {
	log.Logger.Warn("nullableTenantServiceClient.FindTenantIdListByUserId method")
	return &pb.FindTenantIdListByUserIdResponse{}, nil
}

func (*nullableTenantServiceClient) CreateTenant(ctx context.Context, in *pb.CreateTenantRequest, opts ...grpc.CallOption) (*pb.CreateTenantResponse, error) {
	log.Logger.Warn("nullableTenantServiceClient.CreateTenant method")
	return &pb.CreateTenantResponse{}, nil
}

// 删除租户
func (*nullableTenantServiceClient) DeleteTenant(ctx context.Context, in *pb.DeleteTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	log.Logger.Warn("nullableTenantServiceClient.DeleteTenant method")
	return &emptypb.Empty{}, nil
}

// #endregion
