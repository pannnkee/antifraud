package consterrors

// 自定义错误码
const (
	Success int = 0         // 成功
	Failed  int = 100050001 // 失败
)

var ResponseMap = map[int]string{
	Success: "success",
	Failed:  "failed",
}
