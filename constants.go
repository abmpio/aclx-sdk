package sdk

const (
	PolicyType_P string = "p"

	SubOwner_Host string = ""
	SubOwner_Any  string = "*"

	ApiRule_Any string = "*"
)

const (
	RoleName_Users string = "users"
	RoleName_Admin string = "admin"

	RoleDescription_Admin = "系统管理员"
	RoleDescription_Users = "普通用户"
)

const (
	// api model
	Name_ApiModel    = "api-model"
	Name_ApiAdapter  = "api-adapter"
	Name_ApiEnforcer = "api-enforcer"

	// user model
	Name_UserModel    = "user-model"
	Name_UserAdapter  = "user-adapter"
	Name_UserEnforcer = "user-enforcer"

	ActionName_Read  string = "Read"
	ActionName_Write string = "Write"
	ActionName_Admin string = "Admin"

	EffectName_Allow string = "Allow"
	EffectName_Deny  string = "Deny"
)
