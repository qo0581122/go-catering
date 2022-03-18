package config

type System struct {
	Env          string `json:"env" yaml:"env"`          // 环境值
	Addr         int    `json:"addr" yaml:"addr"`        // 端口值
	DbType       string `json:"dbType" yaml:"db-type"`   // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType      string `json:"ossType" yaml:"oss-type"` // Oss类型
	LimitCountIP int    `json:"iplimitCount" yaml:"iplimit-count"`
	LimitTimeIP  int    `json:"iplimitTime" yaml:"iplimit-time"`
}
