package zapzhkratos

import (
	"testing"

	"github.com/yyle88/zaplog"
)

func TestNew创建匝普日志(t *testing.T) {
	v匝普日志 := New匝普日志(zaplog.LOGGER, New日志配置())

	v模块匝普 := v匝普日志.Sub模块匝普()

	v模块匝普.LOG.Info("abc")
	v模块匝普.SUG.Info("xyz")
}

func TestNew创建匝普日志_With模块位置键名(t *testing.T) {
	v匝普日志 := New匝普日志(zaplog.LOGGER, New日志配置().With模块位置键名("module"))

	v模块匝普 := v匝普日志.Sub模块匝普()

	v模块匝普.LOG.Info("abc")
	v模块匝普.SUG.Info("xyz")
}
