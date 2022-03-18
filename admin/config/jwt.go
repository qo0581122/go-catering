package config

type JWT struct {
	SigningKey  string `json:"signingKey" yaml:"signing-key"`   // jwt签名
	ExpiresTime int64  `json:"expiresTime" yaml:"expires-time"` // 过期时间
	BufferTime  int64  `json:"bufferTime" yaml:"buffer-time"`   // 缓冲时间
	Issuer      string `json:"issuer" yaml:"issuer"`            // 签发者
}
