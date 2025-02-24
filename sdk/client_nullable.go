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
}

type nullableRoleServiceClient struct {
}

type nullableUserServiceClient struct {
}

type nullableTenantServiceClient struct {
}

var _ IClient = (*NullableClient)(nil)
var _ pb.AclxClient = (*NullableClient)(nil)
var _ pb.RoleServiceClient = (*NullableClient)(nil)
var _ pb.UserServiceClient = (*NullableClient)(nil)
var _ pb.TenantServiceClient = (*NullableClient)(nil)

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

// #region nullableRoleServiceClient Members

func (*nullableRoleServiceClient) EnsureRoleExist(ctx context.Context, in *pb.EnsureRoleExistRequest, opts ...grpc.CallOption) (*pb.EnsureRoleExistResponse, error) {
	log.Logger.Warn("nullableRoleServiceClient.EnsureRoleExist method")
	return &pb.EnsureRoleExistResponse{
		Successful: false,
	}, nil
}

func (*nullableRoleServiceClient) AddUserListToRole(ctx context.Context, in *pb.AddUserListToRoleRequest, opts ...grpc.CallOption) (*pb.AddUserListToRoleResponse, error) {
	log.Logger.Warn("nullableRoleServiceClient.AddUserListToRole method")
	return &pb.AddUserListToRoleResponse{
		Successful: false,
	}, nil
}

// #endregion

// #region nullableUserServiceClient Members

func (*nullableUserServiceClient) UserLogin(ctx context.Context, in *pb.UserServiceLoginRequest, opts ...grpc.CallOption) (*pb.UserServiceLoginResponse, error) {
	log.Logger.Warn("nullableUserServiceClient.UserLogin method")
	return &pb.UserServiceLoginResponse{
		Result: &pb.UserLoginResult{},
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

// #endregion

// #region nullableTenantServiceClient Members

func (*nullableTenantServiceClient) FindOneByTenantId(ctx context.Context, in *pb.FindByTenantIdRequest, opts ...grpc.CallOption) (*pb.TenantServiceFindOneByTenantIdResponse, error) {
	log.Logger.Warn("nullableTenantServiceClient.FindOneByTenantId method")
	return &pb.TenantServiceFindOneByTenantIdResponse{}, nil
}

func (*nullableTenantServiceClient) RemoveUserFromTenant(ctx context.Context, in *pb.RemoveUserFromTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	log.Logger.Warn("nullableTenantServiceClient.RemoveUserIdFromTenant method")
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

func (c *NullableClient) GetById(ctx context.Context, id string) (*tenancy.TenantConfig, error) {
	log.Logger.Warn("NullableClient.GetById method")
	return tenancy.NewTenantConfig("",
		"",
		"",
	), nil
}
