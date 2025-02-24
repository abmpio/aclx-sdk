package sdk

import (
	"errors"
	"fmt"
	"strings"

	"github.com/abmpio/irisx/controllerx"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/kataras/iris/v12"
)

// setup owner/name format string
func SetupNormalizedId(owner, name string) string {
	return fmt.Sprintf("%s/%s", owner, name)
}

// split two Field from string ,separator value is /
func GetOwnerAndNameFromId(id string) (string, string, error) {
	tokens := strings.Split(id, "/")
	if len(tokens) != 2 {
		return "", "", errors.New("GetOwnerAndNameFromId() error, wrong token count for ID: " + id)
	}

	return tokens[0], tokens[1], nil
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
