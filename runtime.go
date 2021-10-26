package atom

import "runtime"

// RuntimeFunc 取得执行中的函数名
func RuntimeFunc(skip int) string {
	if pc, _, _, ok := runtime.Caller(skip); ok {
		return runtime.FuncForPC(pc).Name()
	}
	return ""
}

