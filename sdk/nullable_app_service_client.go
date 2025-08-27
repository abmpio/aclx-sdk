package sdk

import (
	"context"

	"github.com/abmpio/abmp/pkg/log"
	pb "github.com/abmpio/aclx-sdk/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type nullableAppServiceClient struct {
}

var _ pb.AppServiceClient = (*nullableAppServiceClient)(nil)

// #region

// #region ApplicationServiceClient Members

func (*nullableAppServiceClient) EnsureAppBuiltinRoleListExist(ctx context.Context, in *pb.EnsureAppBuiltinRoleListExistRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	log.Logger.Warn("nullableAppServiceClient.EnsureAppBuiltinRoleListExist method")
	return &emptypb.Empty{}, nil

}

// find list by app
func (*nullableAppServiceClient) FindAppBuiltinRoleListByApp(ctx context.Context, in *pb.FindAppBuiltinRoleListByAppRequest, opts ...grpc.CallOption) (*pb.FindAppBuiltinRoleListByAppResponse, error) {
	log.Logger.Warn("nullableAppServiceClient.FindAppBuiltinRoleListByApp method")
	return &pb.FindAppBuiltinRoleListByAppResponse{}, nil
}

// #endregion
