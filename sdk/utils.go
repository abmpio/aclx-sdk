package sdk

import (
	"fmt"

	"github.com/abmpio/irisx/controllerx"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/kataras/iris/v12"
)

// 构建 owner/name 格式的id
func SetupNormalizedId(owner, name string) string {
	return fmt.Sprintf("%s/%s", owner, name)
}

func MustGetCurrentUserInfo(ctx iris.Context) (*casdoorsdk.Claims, error) {
	userInfo := MayGetCurrentUserInfo(ctx)
	if userInfo == nil {
		return nil, ErrorUserUnauthorized
	}
	return userInfo, nil
}

func MayGetCurrentUserInfo(ctx iris.Context) *casdoorsdk.Claims {
	return controllerx.GetCasdoorMiddleware().GetUserClaims(ctx)
}
