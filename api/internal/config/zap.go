package config

type Zap struct {
	Level         string `yaml:"level"`          // 级别
	Format        string `yaml:"format"`         // 输出
	Prefix        string `yaml:"prefix"`         // 日志前缀
	Director      string `yaml:"director"`       // 日志文件夹
	ShowLine      bool   `yaml:"showLine"`       // 显示行
	EncodeLevel   string `yaml:"encode-level"`   // 编码级
	StacktraceKey string `yaml:"stacktrace-key"` // 栈名
	LogInConsole  bool   `yaml:"log-in-console"` // 输出控制台
}
