package sdk

import (
	"context"

	"github.com/abmpio/abmp/pkg/log"
	pb "github.com/abmpio/aclx-sdk/proto"
	"google.golang.org/grpc"
)

type nullableApplicationServiceClient struct {
}

var _ pb.ApplicationServiceClient = (*nullableApplicationServiceClient)(nil)

// #region

// #region ApplicationServiceClient Members

func (*nullableApplicationServiceClient) FindOneApplication(ctx context.Context, in *pb.FindOneApplicationRequest, opts ...grpc.CallOption) (*pb.FindOneApplicationResponse, error) {
	log.Logger.Warn("nullableApplicationServiceClient.FindOneApplication method")
	return &pb.FindOneApplicationResponse{}, nil
}

// #endregion
