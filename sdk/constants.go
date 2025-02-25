package sdk

import "fmt"

const (
	Name_BuiltIn string = "built-in"
)

const (
	RoleName_Users string = "users"
	RoleName_Admin string = "admin"

	RoleDisplayName_Admin string = "系统管理员"
	RoleDisplayName_Users string = "普通用户"

	RoleDescription_Admin string = "系统管理员"
	RoleDescription_Users string = "普通用户"
)

const (
	// api model
	Name_ApiModel    string = "api-model"
	Name_ApiAdapter  string = "api-adapter"
	Name_ApiEnforcer string = "api-enforcer"

	// user model
	Name_UserModel    string = "user-model"
	Name_UserAdapter  string = "user-adapter"
	Name_UserEnforcer string = "user-enforcer"

	ActionName_Read  string = "Read"
	ActionName_Write string = "Write"
	ActionName_Admin string = "Admin"

	EffectName_Allow string = "Allow"
	EffectName_Deny  string = "Deny"

	Name_Anonymous string = "anonymous"
)

const (
	PolicyType_P string = "p"

	SubOwner_Host string = ""
	SubOwner_Any  string = "*"

	ApiRule_Any string = "*"
)

var (
	ErrorUserUnauthorized = fmt.Errorf("unauthorized")
)
