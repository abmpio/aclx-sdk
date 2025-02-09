package middleware

import (
	"github.com/abmpio/entity/tenancy/iris"
	"github.com/kataras/iris/v12/core/router"
)

func UseMultiTenancy(apiBuilder *router.APIBuilder) {
	//注册中间件
	apiBuilder.Use(iris.MultiTenancy(getServiceGroup().tenantStore))
}
