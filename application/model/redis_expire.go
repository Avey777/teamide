package model

// 设置一个key在当前时间"seconds"(秒)之后过期
type RedisExpire struct {
	Key     string `json:"key,omitempty" yaml:"key,omitempty"`         // 建
	Timeout string `json:"timeout,omitempty" yaml:"timeout,omitempty"` // 过期时间
}

type ActionStepRedisExpire struct {
	Base *ActionStepBase

	RedisExpire      *RedisExpire `json:"redisExpire,omitempty" yaml:"redisExpire,omitempty"`           // 执行 SQL DELETE 操作
	VariableName     string       `json:"variableName,omitempty" yaml:"variableName,omitempty"`         // 变量名称
	VariableDataType string       `json:"variableDataType,omitempty" yaml:"variableDataType,omitempty"` // 变量数据类型
}

func (this_ *ActionStepRedisExpire) GetBase() *ActionStepBase {
	return this_.Base
}

func (this_ *ActionStepRedisExpire) SetBase(v *ActionStepBase) {
	this_.Base = v
}
