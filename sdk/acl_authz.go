package sdk

type IAclAuthz interface {
	IsApiAllowed(subOwner string, subName string, method string, urlPath string, objOwner string, objName string) (bool, error)
	// 检测用户是否有登录的权限
	CheckLoginPermission(tenantId, userName string) (bool, error)
}
