package system

import (
	"catering/config"
)

// 配置文件结构体
type System struct {
	Config config.Config `json:"config"`
}
