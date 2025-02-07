package middleware

import (
	"fmt"

	"github.com/abmpio/irisx/controllerx"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/kataras/iris/v12"
)

var (
	ErrorUserUnauthorized = fmt.Errorf("unauthorized")
)

func MustGetCurrentUserInfo(ctx iris.Context) (*casdoorsdk.Claims, error) {
	userInfo, err := MayGetCurrentUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	if userInfo == nil {
		return nil, ErrorUserUnauthorized
	}
	return userInfo, nil
}

func MayGetCurrentUserInfo(ctx iris.Context) (*casdoorsdk.Claims, error) {
	claims := controllerx.GetCasdoorMiddleware().GetUserClaims(ctx)
	if claims == nil {
		return nil, nil
	}
	return claims, nil
}
