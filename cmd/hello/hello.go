package hello

import (
	"yaoplugin/utils"
)

func Echo(args ...interface{}) map[string]interface{} {
	// 使用日志实例记录消息
	utils.Log("这是一条信息级别的日志", args)
	utils.Logf("带有格式的日志: %d", 123)

	if len(args) > 0 {
		return map[string]interface{}{"code": 400, "message": "该函数不接受参数"}
	}
	return map[string]interface{}{"code": 0, "message": "hello world"}
}
