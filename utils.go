package sdk

import "fmt"

// 构建 owner/name 格式的id
func SetupNormalizedId(owner, name string) string {
	return fmt.Sprintf("%s/%s", owner, name)
}
