package middleware

import (
	"sync"

	"github.com/abmpio/aclx-sdk/sdk"
	"github.com/abmpio/app"
	"github.com/abmpio/entity/tenancy"
)

var (
	_serviceFactory             *serviceGroup
	_serviceFactoryInstanceOnce sync.Once
)

type serviceGroup struct {
	aclAuthz    sdk.IAclAuthz
	tenantStore tenancy.ITenantStore
}

// 获取ServiceGroup实例
func getServiceGroup() *serviceGroup {
	if _serviceFactory != nil {
		return _serviceFactory
	}
	_serviceFactoryInstanceOnce.Do(func() {
		_serviceFactory = newServiceGroup()
	})
	return _serviceFactory
}

// 创建一个新的实例
func newServiceGroup() *serviceGroup {
	serviceFactory := &serviceGroup{
		aclAuthz:    app.Context.GetInstance(new(sdk.IAclAuthz)).(sdk.IAclAuthz),
		tenantStore: app.Context.GetInstance(new(tenancy.ITenantStore)).(tenancy.ITenantStore),
	}
	return serviceFactory
}
