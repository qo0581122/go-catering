package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// oss
	AliyunOSS AliyunOSS `mapstructure:"aliyun-oss" json:"aliyunOSS" yaml:"aliyun-oss"`

	Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`
	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}

type Config struct {
	JWT     JWT     `json:"jwt" yaml:"jwt"`
	Zap     Zap     `json:"zap" yaml:"zap"`
	Redis   Redis   `json:"redis" yaml:"redis"`
	Email   Email   `json:"email" yaml:"email"`
	Casbin  Casbin  `json:"casbin" yaml:"casbin"`
	System  System  `json:"system" yaml:"system"`
	Captcha Captcha `json:"captcha" yaml:"captcha"`
	// gorm
	Mysql Mysql `json:"mysql" yaml:"mysql"`
	// oss
	Local     Local     `mapstructure:"local" json:"local" yaml:"local"`
	AliyunOSS AliyunOSS `json:"aliyunOSS" yaml:"aliyun-oss"`

	Excel Excel `json:"excel" yaml:"excel"`
	// 跨域配置
	Cors CORS `json:"cors" yaml:"cors"`
}
