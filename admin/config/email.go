package config

type Email struct {
	To       string `json:"to" yaml:"to"`             // 收件人:多个以英文逗号分隔
	Port     int    `json:"port" yaml:"port"`         // 端口
	From     string `json:"from" yaml:"from"`         // 收件人
	Host     string `json:"host" yaml:"host"`         // 服务器地址
	IsSSL    bool   `json:"isSSL" yaml:"is-ssl"`      // 是否SSL
	Secret   string `json:"secret" yaml:"secret"`     // 密钥
	Nickname string `json:"nickname" yaml:"nickname"` // 昵称
}
