package sdk

import (
	"context"

	"github.com/abmpio/abmp/pkg/log"
	"github.com/abmpio/aclx-sdk/options"
	pb "github.com/abmpio/aclx-sdk/proto"
	"github.com/abmpio/entity/tenancy"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NullableClient struct {
	nullableRoleServiceClient
	nullableUserServiceClient
	nullableTenantServiceClient
	nullableApplicationServiceClient
	nullableAppServiceClient
}

var _ IClient = (*NullableClient)(nil)
var _ pb.AclxClient = (*NullableClient)(nil)

// #region

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

func (*NullableClient) AclxIsApiAllowed(ctx context.Context, in *pb.AclxIsApiAllowedRequest, opts ...grpc.CallOption) (*pb.AclxIsApiAllowedResponse, error) {
	log.Logger.Warn("NullableClient.AclxIsApiAllowed method")
	return &pb.AclxIsApiAllowedResponse{
		IsAllowed: true,
	}, nil
}

func (*NullableClient) AclxCheckLoginPermission(ctx context.Context, in *pb.AclxCheckLoginPermissionRequest, opts ...grpc.CallOption) (*pb.AclxCheckLoginPermissionResponse, error) {
	log.Logger.Warn("NullableClient.AclxCheckLoginPermission method")
	return &pb.AclxCheckLoginPermissionResponse{
		IsAllowed: true,
	}, nil
}

func (*NullableClient) AclxCheckApiPermission(ctx context.Context, in *pb.AclxCheckApiPermissionRequest, opts ...grpc.CallOption) (*pb.AclxCheckApiPermissionResponse, error) {
	log.Logger.Warn("NullableClient.AclxCheckApiPermission method")
	return &pb.AclxCheckApiPermissionResponse{
		IsAllowed: true,
	}, nil
}

func (*NullableClient) AclxInitSystemAPI(ctx context.Context, in *pb.AclxInitSystemAPIRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	log.Logger.Warn("NullableClient.AclxInitSystemAPI method")
	return &emptypb.Empty{}, nil
}

// #endregion

// #region IAclAuthz members

func (c *NullableClient) IsApiAllowed(subOwner string, subName string, method string, urlPath string, objOwner string, objName string) (bool, error) {
	r, err := c.AclxIsApiAllowed(context.TODO(), &pb.AclxIsApiAllowedRequest{
		App:      options.GetOptions().DefaultApp,
		SubOwner: subOwner,
		SubName:  subName,
		Method:   method,
		UrlPath:  urlPath,
		ObjOwner: objOwner,
		ObjName:  objName,
	})
	return r.IsAllowed, err
}

// 检测用户是否有登录的权限
func (c *NullableClient) CheckLoginPermission(tenantId, userId string) (bool, error) {
	r, err := c.AclxCheckLoginPermission(context.TODO(), &pb.AclxCheckLoginPermissionRequest{
		App:      options.GetOptions().DefaultApp,
		TenantId: tenantId,
		UserId:   userId,
	})
	return r.IsAllowed, err
}

// endregion

// #region tenancy.ITenantStore Members

func (c *NullableClient) GetById(ctx context.Context, id string) (*tenancy.TenantConfig, error) {
	log.Logger.Warn("NullableClient.GetById method")
	return tenancy.NewTenantConfig("",
		"",
		"",
	), nil
}

// #endregion
