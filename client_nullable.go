package sdk

import (
	"context"

	"github.com/abmpio/abmp/pkg/log"
	pb "github.com/abmpio/aclx/sdk/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NullableClient struct {
}

var _ IClient = (*NullableClient)(nil)
var _ pb.AclxClient = (*NullableClient)(nil)
var _ pb.RoleServiceClient = (*NullableClient)(nil)
var _ pb.UserServiceClient = (*NullableClient)(nil)

func (*NullableClient) AclxHealthCheck(ctx context.Context, in *pb.AclxHealthCheckRequest, opts ...grpc.CallOption) (*pb.AclxHealthCheckResponse, error) {
	return &pb.AclxHealthCheckResponse{
		Status: pb.AclxHealthCheckResponse_NOT_SERVING,
	}, nil
}

// 注册acl参数
func (*NullableClient) RegistAppAcl(ctx context.Context, in *pb.RegistAppAclRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	log.Logger.Warn("NullableClient.RegistAppAcl method")
	return &emptypb.Empty{}, nil
}

func (*NullableClient) RegistAclApiRule(ctx context.Context, in *pb.RegistAclApiRuleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	log.Logger.Warn("NullableClient.RegistAclApiRule method")
	return &emptypb.Empty{}, nil
}

func (*NullableClient) RegistAclPermissionRule(ctx context.Context, in *pb.RegistAclPermissionRuleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	log.Logger.Warn("NullableClient.RegistAclPermissionRule method")
	return &emptypb.Empty{}, nil
}

func (*NullableClient) ReloadAppAcl(ctx context.Context, in *pb.ReloadAppAclRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	log.Logger.Warn("NullableClient.ReloadAppAcl method")
	return &emptypb.Empty{}, nil
}

func (*NullableClient) EnsureRoleExist(ctx context.Context, in *pb.EnsureRoleExistRequest, opts ...grpc.CallOption) (*pb.EnsureRoleExistResponse, error) {
	log.Logger.Warn("NullableClient.EnsureRoleExist method")
	return &pb.EnsureRoleExistResponse{
		Successful: false,
	}, nil
}

func (*NullableClient) UserLogin(ctx context.Context, in *pb.UserLoginRequest, opts ...grpc.CallOption) (*pb.UserLoginResponse, error) {
	log.Logger.Warn("NullableClient.UserLogin method")
	return &pb.UserLoginResponse{
		Result: &pb.UserLoginResult{},
	}, nil
}
