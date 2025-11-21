package zapzhkratos

import (
	"testing"

	"github.com/yyle88/zaplog"
)

func TestT匝普日志_Get奎沱主簿(t *testing.T) {
	v匝普日志 := New匝普日志(zaplog.LOGGER, New日志配置())

	slog := v匝普日志.Get奎沱主簿("测试用例")

	slog.Info("woca", "[a]", "[b]", "[c]")
	slog.Infow("k", "v", "k1", "v2")
}

func TestT匝普日志_New奎沱秘书(t *testing.T) {
	v匝普日志 := New匝普日志(zaplog.LOGGER, New日志配置())

	slog := v匝普日志.New奎沱秘书("测试秘书")

	slog.Info("我爱杨亦乐")
	slog.Infow("我爱", "杨亦乐")
}
