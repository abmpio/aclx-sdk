package sdk

import (
	"context"

	"github.com/abmpio/abmp/pkg/log"
	pb "github.com/abmpio/aclx-sdk/proto"
	"google.golang.org/grpc"
)

type nullableRoleServiceClient struct {
}

var _ pb.RoleServiceClient = (*nullableRoleServiceClient)(nil)

// #region

// #region RoleServiceClient Members

func (*nullableRoleServiceClient) EnsureRoleExist(ctx context.Context, in *pb.EnsureRoleExistRequest, opts ...grpc.CallOption) (*pb.EnsureRoleExistResponse, error) {
	log.Logger.Warn("nullableRoleServiceClient.EnsureRoleExist method")
	return &pb.EnsureRoleExistResponse{
		Successful: false,
	}, nil
}

// #endregion
