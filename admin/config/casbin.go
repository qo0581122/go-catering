package config

type Casbin struct {
	ModelPath string `json:"modelPath" yaml:"model-path"` // 存放casbin模型的相对路径
}
