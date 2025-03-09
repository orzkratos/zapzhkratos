package zapzhkratos

import (
	"github.com/orzkratos/zapkratos"
	"github.com/yyle88/zaplog"
)

type T日志配置 zapkratos.Options

func New日志配置() *T日志配置 {
	p := zapkratos.NewOptions()
	p.WithModuleKeyName("模块位置")
	return (*T日志配置)(p)
}

func (T *T日志配置) With模块位置键名(v模块位置键名 string) *T日志配置 {
	p := (*zapkratos.Options)(T)
	p.WithModuleKeyName(v模块位置键名)
	return (*T日志配置)(p)
}

// T匝普日志 这里主要做的就是把英文的类型名改为中文类型名，再增加几个简单的逻辑
type T匝普日志 zapkratos.ZapKratos

func New匝普日志(zap *zaplog.Zap, options *T日志配置) *T匝普日志 {
	return (*T匝普日志)(zapkratos.NewZapKratos(zap, (*zapkratos.Options)(options)))
}

func (A *T匝普日志) Get基本匝普() *zaplog.Zap {
	p := (*zapkratos.ZapKratos)(A)
	return p.GetZap()
}

func (A *T匝普日志) Sub模块匝普() *zaplog.Zap {
	p := (*zapkratos.ZapKratos)(A)
	return p.SubZap()
}
