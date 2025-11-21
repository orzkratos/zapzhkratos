package zapzhkratos

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/orzkratos/zapkratos"
)

// Get奎沱日志 creates Kratos log.Logger with given message caption
// Wraps underlying Zap and returns the created log instance
//
// Get奎沱日志 使用给定的消息说明创建奎沱 log.Logger
// 包装底层匝普日志器并返回实例
func (A *T匝普日志) Get奎沱日志(msg说明 string) log.Logger {
	return zapkratos.NewLogImp(A.Get基本匝普().LOG, msg说明)
}

// New奎沱日志 creates Kratos log.Logger with given message caption
// Wraps underlying Zap and returns the created log instance
// Note: Same as Get奎沱日志, provided to match naming patterns
//
// New奎沱日志 使用给定的消息说明创建奎沱 log.Logger
// 包装底层匝普日志器并返回实例
// 注意：与 Get奎沱日志 相同，提供以匹配命名模式
func (A *T匝普日志) New奎沱日志(msg说明 string) log.Logger {
	return zapkratos.NewLogImp(A.Get基本匝普().LOG, msg说明)
}

// Get奎沱主簿 creates Kratos log.Helper with given message caption
// Provides convenient methods and returns the created Helper instance
//
// Get奎沱主簿 使用给定的消息说明创建奎沱 log.Helper
// 提供便捷日志方法并返回主簿实例
func (A *T匝普日志) Get奎沱主簿(msg说明 string) *log.Helper {
	return log.NewHelper(A.Get奎沱日志(msg说明))
}

// New奎沱主簿 creates Kratos log.Helper with given message caption
// Provides convenient methods and returns the created Helper instance
// Note: Same as Get奎沱主簿, provided to match naming patterns
//
// New奎沱主簿 使用给定的消息说明创建奎沱 log.Helper
// 提供便捷日志方法并返回主簿实例
// 注意：与 Get奎沱主簿 相同，提供以匹配命名模式
func (A *T匝普日志) New奎沱主簿(msg说明 string) *log.Helper {
	return log.NewHelper(A.Get奎沱日志(msg说明))
}

// Get奎沱秘书 creates Kratos log.Helper with given message caption
// Provides convenient methods and returns the created Helper instance
// Note: "秘书" is alternative name to "主簿" in Chinese
//
// Get奎沱秘书 使用给定的消息说明创建奎沱 log.Helper
// 提供便捷日志方法并返回秘书实例
// 注意："秘书"是"主簿"的备选中文名称
func (A *T匝普日志) Get奎沱秘书(msg说明 string) *log.Helper {
	return log.NewHelper(A.Get奎沱日志(msg说明))
}

// New奎沱秘书 creates Kratos log.Helper with given message caption
// Provides convenient methods and returns the created Helper instance
// Note: "秘书" is alternative name to "主簿" in Chinese
//
// New奎沱秘书 使用给定的消息说明创建奎沱 log.Helper
// 提供便捷日志方法并返回秘书实例
// 注意："秘书"是"主簿"的备选中文名称
func (A *T匝普日志) New奎沱秘书(msg说明 string) *log.Helper {
	return log.NewHelper(A.Get奎沱日志(msg说明))
}
