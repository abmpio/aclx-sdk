package middleware

import (
	"fmt"
	"net/http"

	"github.com/abmpio/abmp/pkg/log"
	"github.com/abmpio/aclx/sdk"
	"github.com/abmpio/aclx/sdk/options"
	"github.com/abmpio/entity/tenancy"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/kataras/iris/v12"
	irisContext "github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
)

func UseAcl(apiBuilder *router.APIBuilder) {
	if !options.GetOptions().Disabled {
		apiBuilder.Use(serveHTTP)
		log.Logger.Info("已启用acl中间件")
	}
}

func serveHTTP(ctx *irisContext.Context) {
	subOwner, subName := getSubject(ctx)
	method := ctx.Method()
	urlPath := getUrlPath(ctx.RequestPath(false))

	objOwner, objName := "", ""
	if urlPath != "/login" {
		var err error
		objOwner, objName, err = getObject(ctx)
		if err != nil {
			log.Logger.Warn(fmt.Sprintf("acl error in acl middleware,err:%s", err))
			ctx.StopExecution()
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
			return
		}
	}
	isAllowed, err := getServiceGroup().aclAuthz.IsApiAllowed(subOwner, subName, method, urlPath, objOwner, objName)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("在对api请求进行授权时出现异常,url:%s,err:%s", urlPath, err.Error()))
		denyRequest(ctx)
		return
	}
	// result := "deny"
	// if isAllowed {
	// 	result = "allow"
	// }
	if !isAllowed {
		denyRequest(ctx)
		return
	}
	ctx.Next()
}

func denyRequest(ctx *irisContext.Context) {
	m := "auth:Unauthorized operation"
	log.Logger.Warn(m)
	ctx.StopExecution()
	ctx.StatusCode(iris.StatusForbidden)
	ctx.WriteString(m)
}

// return subjectOwner,subject
func getSubject(ctx *irisContext.Context) (string, string) {
	subOwner, subName := "", ""
	subOwner, _ = tenancy.TenantIdFromContext(ctx)
	if subOwner == "" {
		subOwner = appFromHeader(ctx)
	}
	userInfo := getUserInfo(ctx)
	if userInfo == nil {
		subName = sdk.Name_Anonymous
	} else {
		subName = userInfo.Name
	}
	return subOwner, subName
}

// get app from context
func appFromHeader(ctx *irisContext.Context) string {
	return ctx.GetHeader("X-App")
}

func getObject(ctx *irisContext.Context) (string, string, error) {
	method := ctx.Method()

	tenantId, _ := tenancy.TenantIdFromContext(ctx)
	if method == http.MethodGet {
		tenantId := ctx.Request().URL.Query().Get("tenantId")
		if tenantId != "" {
			return tenantId, "", nil
		}
		return tenantId, "", nil
	}
	return tenantId, "", nil
}

func getUserInfo(ctx *irisContext.Context) *casdoorsdk.Claims {
	userInfo := sdk.MayGetCurrentUserInfo(ctx)
	return userInfo
}

func getUrlPath(urlPath string) string {
	// some map
	return urlPath
}
