package config

type Config struct {
	JWT     JWT     `json:"jwt" yaml:"jwt"`
	Zap     Zap     `json:"zap" yaml:"zap"`
	Redis   Redis   `json:"redis" yaml:"redis"`
	Casbin  Casbin  `json:"casbin" yaml:"casbin"`
	System  System  `json:"system" yaml:"system"`
	Captcha Captcha `json:"captcha" yaml:"captcha"`
	// gorm
	Mysql Mysql `json:"mysql" yaml:"mysql"`
	// oss
	Local     Local     `json:"local" yaml:"local"`
	AliyunOSS AliyunOSS `json:"aliyunOSS" yaml:"aliyun-oss"`

	Excel Excel `json:"excel" yaml:"excel"`
	// 跨域配置
	Cors CORS `json:"cors" yaml:"cors"`
}
