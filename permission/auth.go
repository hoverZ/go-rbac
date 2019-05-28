package permission

const (
	NotManaged    = "Not managed"   // 不在权限树内的接口，无法使用
	FreeToUse     = "Free"          // 在权限树内，且 status = 0，表示无须授权，可免费试用
	Authorization = "Authorization" // 在权限树内，且 status = 1，且用户拥有该权限，表示授权使用
	NoPermission  = "No permission" // 在权限树内，且 status = 1，且用户没有改权限，表示没有权限
)

//func
