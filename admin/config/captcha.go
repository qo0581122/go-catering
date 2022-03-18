package config

type Captcha struct {
	KeyLong   int `json:"keyLong" yaml:"key-long"`     // 验证码长度
	ImgWidth  int `json:"imgWidth" yaml:"img-width"`   // 验证码宽度
	ImgHeight int `json:"imgHeight" yaml:"img-height"` // 验证码高度
}
