package zapzhkratos

import (
	"path/filepath"

	"github.com/yyle88/runpath"
	"github.com/yyle88/zaplog"
)

type T日志配置 struct {
	V模块位置键名 string
}

func New日志配置() *T日志配置 {
	return &T日志配置{
		V模块位置键名: "模块位置",
	}
}

// T匝普日志 这里主要做的就是把英文的类型名改为中文类型名，再增加几个简单的逻辑
type T匝普日志 struct {
	zap *zaplog.Zap
	opt *T日志配置
}

func New匝普日志(zap *zaplog.Zap, opt *T日志配置) *T匝普日志 {
	return &T匝普日志{
		zap: zap,
		opt: opt,
	}
}

func (A *T匝普日志) Get基本匝普() *zaplog.Zap {
	return A.zap
}

func (A *T匝普日志) Sub模块匝普() *zaplog.Zap {
	return A.Get基本匝普().SubZap2(A.opt.V模块位置键名, filepath.Base(runpath.Skip(1)))
}
