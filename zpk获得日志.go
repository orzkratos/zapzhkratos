package zapzhkratos

import "github.com/go-kratos/kratos/v2/log"

func (A *T匝普日志) Get奎沱日志(msg模块 string) log.Logger {
	return New奎沱日志(A.Get基本匝普().LOG, msg模块)
}

func (A *T匝普日志) New奎沱日志(msg模块 string) log.Logger {
	return New奎沱日志(A.Get基本匝普().LOG, msg模块)
}

func (A *T匝普日志) Get奎沱主簿(msg模块 string) *log.Helper {
	return log.NewHelper(A.Get奎沱日志(msg模块))
}

func (A *T匝普日志) New奎沱秘书(msg模块 string) *log.Helper {
	return log.NewHelper(A.Get奎沱日志(msg模块))
}
