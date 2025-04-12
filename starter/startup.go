package starter

import (
	"context"
	"fmt"
	"time"

	"github.com/abmpio/abmp/pkg/log"
	"github.com/abmpio/aclx-sdk/options"
	pb "github.com/abmpio/aclx-sdk/proto"
	"github.com/abmpio/aclx-sdk/sdk"
	"github.com/abmpio/app"
	"github.com/abmpio/app/cli"
	"github.com/abmpio/entity/tenancy"
)

func init() {
	fmt.Println("abmpio.aclx_sdk.starter init")

	cli.ConfigureService(serviceConfigurator)
}

func serviceConfigurator(wa cli.CliApplication) {
	if app.HostApplication.SystemConfig().App.IsRunInCli {
		return
	}
	var _client sdk.IClient

	opt := options.GetOptions()
	if !opt.Disabled {
		aclClient := sdk.NewClient(sdk.WithHost(opt.Host), sdk.WithPort(opt.Port))
		//测试ping
		for {
			err := aclClient.InitConnnection()
			if err != nil {
				log.Logger.Warn(fmt.Sprintf("初始化aclx grpc连接时出现异常,option:%s, err:%s",
					opt.String(),
					err.Error()))

			} else {
				res, err := aclClient.AclxHealthCheck(context.TODO(), &pb.AclxHealthCheckRequest{})
				if err != nil {
					log.Logger.Warn(fmt.Sprintf("检测aclx grpc 服务健康是否正常时出现异常,option:%s, err:%s",
						opt.String(),
						err.Error()))
				} else {
					if res != nil {
						log.Logger.Info(fmt.Sprintf("aclx grpc status:%s", res.Status.String()))
					}
					// set client
					_client = aclClient
					break
				}
			}
			log.Logger.Warn("2s后重新测试...")
			time.Sleep(2 * time.Second)
		}
	} else {
		log.Logger.Warn("aclx grpc client disabled")
		_client = &sdk.NullableClient{}
	}
	sdk.SetGlobalClient(_client)
	app.Context.RegistInstanceAs(_client, new(sdk.IClient))
	app.Context.RegistInstanceAs(_client, new(pb.AclxClient))
	app.Context.RegistInstanceAs(_client, new(pb.RoleServiceClient))
	app.Context.RegistInstanceAs(_client, new(pb.UserServiceClient))
	app.Context.RegistInstanceAs(_client, new(sdk.IAclAuthz))
	app.Context.RegistInstanceAs(_client, new(tenancy.ITenantStore))
}
