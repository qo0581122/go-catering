package config

type Zap struct {
	Level         string `json:"level" yaml:"level"`                  // 级别
	Format        string `json:"format" yaml:"format"`                // 输出
	Prefix        string `json:"prefix" yaml:"prefix"`                // 日志前缀
	Director      string `json:"director"  yaml:"director"`           // 日志文件夹
	ShowLine      bool   `json:"showLine" yaml:"showLine"`            // 显示行
	EncodeLevel   string `json:"encodeLevel" yaml:"encode-level"`     // 编码级
	StacktraceKey string `json:"stacktraceKey" yaml:"stacktrace-key"` // 栈名
	LogInConsole  bool   `json:"logInConsole" yaml:"log-in-console"`  // 输出控制台
}
