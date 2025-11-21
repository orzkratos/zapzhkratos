// Package zapzhkratos integrates Zap logging with Kratos microservice framework using Chinese names
// Wraps zapkratos package to provide intuitive Chinese function names and method names
// Enables structured logging with module tracking and flexible configuration options
// Designed to make Chinese developers more productive with intuitive naming conventions
//
// zapzhkratos 使用中文命名将匝普日志集成到奎沱微服务框架
// 包装 zapkratos 包以提供直观的中文函数名和方法名
// 支持结构化日志记录、模块追踪和灵活的配置选项
// 旨在通过熟悉的命名约定让中文开发者更高效
package zapzhkratos

import (
	"github.com/orzkratos/zapkratos"
	"github.com/yyle88/zaplog"
)

// T日志配置 wraps zapkratos.Options to provide Chinese-named configuration interface
// Controls module field naming and various logging settings
// Type alias pattern enables seamless interoperation with underlying zapkratos package
//
// T日志配置 包装 zapkratos.Options 以提供中文命名的配置接口
// 控制模块字段命名和各种日志行为设置
// 类型别名模式实现与底层 zapkratos 包的无缝互操作
type T日志配置 zapkratos.Options

// New日志配置 creates configuration instance with optimized Chinese-oriented default settings
// Sets module field name to "模块位置" (Chinese version of "module") and returns config instance
// Default value chosen to align with Chinese developers' expectations and conventions
//
// New日志配置 创建具有优化的中文友好默认设置的配置实例
// 设置模块键名为 "模块位置"（"module" 的中文版本）并返回配置实例
// 选择的默认值符合中文开发者的期望和约定
func New日志配置() *T日志配置 {
	p := zapkratos.NewOptions()
	p.WithModuleKeyName("模块位置")
	return (*T日志配置)(p)
}

// With模块位置键名 sets custom module field name using fluent chaining method
// Returns same configuration instance to enable method chaining in fluent style
// Allows customization of module field name in log output to match team conventions
//
// With模块位置键名 使用流式构建器模式设置自定义模块字段键名
// 返回相同的配置实例以支持流式风格的方法链式调用
// 允许自定义日志输出中的模块字段名以匹配团队约定
func (T *T日志配置) With模块位置键名(v模块位置键名 string) *T日志配置 {
	p := (*zapkratos.Options)(T)
	p.WithModuleKeyName(v模块位置键名)
	return (*T日志配置)(p)
}

// T匝普日志 wraps zapkratos.ZapKratos to provide Kratos-compatible logging with Chinese naming
// Holds underlying Zap instance and configuration to create various logging instances on demand
// Type alias enables seamless integration while providing intuitive Chinese method names
//
// T匝普日志 包装 zapkratos.ZapKratos 以提供带中文命名的奎沱兼容日志
// 持有底层匝普实例和配置以按需创建各种日志实例
// 类型别名实现无缝集成同时提供直观的中文方法名
type T匝普日志 zapkratos.ZapKratos

// New匝普日志 creates instance with given Zap logging engine and custom configuration
// Returns new instance capable of providing multiple types of Kratos logging interfaces
// Initializes own state and prepares to create log.Logger and log.Helper on demand
//
// New匝普日志 使用给定的匝普日志引擎和自定义配置创建实例
// 返回能够提供多种类型奎沱日志接口的新实例
// 初始化内部状态并准备按需创建 log.Logger 和 log.Helper
func New匝普日志(zap *zaplog.Zap, options *T日志配置) *T匝普日志 {
	return (*T匝普日志)(zapkratos.NewZapKratos(zap, (*zapkratos.Options)(options)))
}

// Get基本匝普 returns the underlying Zap logging instance unwrapped from Kratos adaptation
// Provides access to raw Zap when advanced features beyond Kratos interface are needed
// Enables direct usage of Zap's complete API surface when needed
//
// Get基本匝普 返回从奎沱适配中解包的底层匝普日志器实例
// 在需要奎沱接口之外的高级功能时提供对原始匝普的访问
// 在必要时允许直接使用匝普的完整 API 表面
func (A *T匝普日志) Get基本匝普() *zaplog.Zap {
	p := (*zapkratos.ZapKratos)(A)
	return p.GetZap()
}

// Sub模块匝普 creates child Zap instance with automatic module information injection
// Auto adds module field containing source file basename obtained from calling context
// Returns new Zap with enriched context that includes module identification data
// Uses runpath package to smart extract calling file information at runtime
//
// Sub模块匝普 创建带自动模块信息注入的子匝普实例
// 自动添加包含从调用上下文获取的源文件基本名的模块字段
// 返回包含丰富上下文的新匝普实例，其中包括模块识别数据
// 使用 runpath 包在运行时智能提取调用文件信息
func (A *T匝普日志) Sub模块匝普() *zaplog.Zap {
	p := (*zapkratos.ZapKratos)(A)
	return p.SubZap()
}
