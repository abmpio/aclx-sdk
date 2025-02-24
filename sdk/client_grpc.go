package sdk

import (
	"context"
	"fmt"

	"github.com/abmpio/abmp/pkg/log"
	"github.com/abmpio/aclx-sdk/options"
	pb "github.com/abmpio/aclx-sdk/proto"
	"github.com/abmpio/entity/tenancy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IClient interface {
	pb.AclxClient
	pb.UserServiceClient
	pb.RoleServiceClient
	pb.TenantServiceClient

	IAclAuthz
	tenancy.ITenantStore
}

type Client struct {
	option *Options

	conn *grpc.ClientConn

	pb.AclxClient
	pb.UserServiceClient
	pb.RoleServiceClient
	pb.TenantServiceClient
}

var _ IClient = (*Client)(nil)
var _ IAclAuthz = (*Client)(nil)
var _ tenancy.ITenantStore = (*Client)(nil)

func NewClient(opts ...Option) *Client {
	client := &Client{
		option: newDefaultOptions(),
	}
	for _, eachOpt := range opts {
		eachOpt(client.option)
	}
	return client
}

func (c *Client) GetOption() *Options {
	return c.option
}

// 初始化client
func (c *Client) InitConnnection(opts ...grpc.DialOption) error {
	mergedOpts := make([]grpc.DialOption, 0)
	mergedOpts = append(mergedOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	mergedOpts = append(mergedOpts, grpc.WithKeepaliveParams(*c.option.keepaliveParam))
	mergedOpts = append(mergedOpts, opts...)
	conn, err := grpc.NewClient(c.option.getHostTarget(), mergedOpts...)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("occur error when create grpc server connection , host:%s,error: %s",
			c.option.getHostTarget(), err.Error()))
		return err
	}
	log.Logger.Info(fmt.Sprintf("initialize grpc connection finished,host:%s", c.option.getHostTarget()))
	c.conn = conn
	//保存客户端
	c.AclxClient = pb.NewAclxClient(conn)
	c.UserServiceClient = pb.NewUserServiceClient(conn)
	c.RoleServiceClient = pb.NewRoleServiceClient(conn)
	c.TenantServiceClient = pb.NewTenantServiceClient(conn)

	return nil
}

// #region IAclAuthz Members

func (c *Client) IsApiAllowed(subOwner string, subName string, method string, urlPath string, objOwner string, objName string) (bool, error) {
	request := &pb.AclxIsApiAllowedRequest{
		App:      options.GetOptions().DefaultApp,
		SubOwner: subOwner,
		SubName:  subName,
		Method:   method,
		UrlPath:  urlPath,
		ObjOwner: objOwner,
		ObjName:  objName,
	}
	response, err := c.AclxIsApiAllowed(context.TODO(), request)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("在对api请求进行授权时出现异常,url:%s,err:%s", urlPath, err.Error()))
		return false, err
	}
	if response == nil {
		log.Logger.Error(fmt.Sprintf("在对api请求进行授权时acl授权中间件没有回应,url:%s", urlPath))
		return false, nil
	}

	return response.IsAllowed, nil
}

// 检测用户是否有登录的权限
func (c *Client) CheckLoginPermission(tenantId, userId string) (bool, error) {
	request := &pb.AclxCheckLoginPermissionRequest{
		App:      options.GetOptions().DefaultApp,
		TenantId: tenantId,
		UserId:   userId,
	}
	response, err := c.AclxCheckLoginPermission(context.TODO(), request)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("在对用户请求是否有登录权限时出现异常,tenantId:%s,userId:%s,err:%s",
			tenantId,
			userId,
			err.Error()))
		return false, err
	}
	if response == nil {
		log.Logger.Error(fmt.Sprintf("在对用户请求是否有登录权限时acl授权中间件没有回应,tenantId:%s,userId:%s",
			tenantId,
			userId))
		return false, nil
	}

	return response.IsAllowed, nil
}

// #endregion

// #region ITenantStore Members

func (c *Client) GetById(ctx context.Context, id string) (*tenancy.TenantConfig, error) {
	response, err := c.FindOneByTenantId(ctx, &pb.FindByTenantIdRequest{
		TenantId: id,
	})
	if err != nil {
		return nil, err
	}
	if response.Tenant == nil {
		return nil, tenancy.ErrTenantNotFound
	}
	tenantConfig := tenancy.NewTenantConfig(response.Tenant.TenantId,
		response.Tenant.Name,
		response.Tenant.AdminUserId,
	)
	return tenantConfig, nil
}

// #endregion
