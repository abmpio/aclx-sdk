package middleware

import (
	"sync"

	"github.com/abmpio/aclx-sdk/sdk"
	"github.com/abmpio/app"
)

var (
	_serviceFactory             *serviceGroup
	_serviceFactoryInstanceOnce sync.Once
)

type serviceGroup struct {
	aclAuthz sdk.IAclAuthz
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
		aclAuthz: app.Context.GetInstance(new(sdk.IAclAuthz)).(sdk.IAclAuthz),
	}
	return serviceFactory
}
