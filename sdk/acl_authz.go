package sdk

type IAclAuthz interface {
	// 检测是否有指定api的访问权限
	// subOwner:通常是请求的租户id
	// subName: 请求的用户id
	// method: http api方法，如Get,Post等
	// urlPath: 请求的api路由
	// objOwner: 请求访问的租户id
	// objName: 请求访问的用户Id
	IsApiAllowed(subOwner string, subName string, method string, urlPath string, objOwner string, objName string) (bool, error)
	// 检测用户是否有登录的权限
	CheckLoginPermission(tenantId, userId string) (bool, error)
}
