package atom

import (
    "fmt"
    logger "github.com/sotyou/go_logger"
    "time"
)

// WaitForNext 有些指标需要整点开始计算，支持(s,m,h)
func WaitForNext(delta time.Duration) {
    var layout = ""
    switch delta {
    case time.Second:
        layout = "2006-01-02 15:04:05"
    case time.Minute:
        layout = "2006-01-02 15:04:00"
    case time.Hour:
        layout = "2006-01-02 15:00:00"
    default:
        logger.ErrorV("[WaitNextFor] unsupported delta", delta)
        return
    }
    start := time.Now()
    next, _ := time.Parse(layout, start.Add(delta).Format(layout))
    time.Sleep(next.Sub(start))
}

// Measuring 一个计时器
// start := time.Now(); defer Microseconds(start)
func Microseconds(start time.Time) int64 {
    micro := time.Now().Sub(start).Microseconds()
    logger.Info(fmt.Sprintf("runtime is %d microseconds", micro), RuntimeFunc(2)) // 调用函数的文件名
    return micro
}

// FormatUTC 当前时间UTC变换, 整形输出
func FormatUTC() string {
    return time.Now().UTC().Format("2006-01-02T15:04:05")
}

// 测量处理时间
func RunningTime(fn func() error) (int64, error) {
    start := time.Now()
    err := fn()
    diff := time.Now().Sub(start).Milliseconds()
    return diff, err
}

func ZeroTime() time.Time {
    return time.Unix(0, 0)
}

func GetMs() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetMsForTime(t time.Time) int64 {
    return t.UnixNano() / int64(time.Millisecond)
}

func WaitForPeriod(period time.Duration) {
    sleepTo := time.Now().Truncate(period).Add(period)
    time.Sleep(sleepTo.Sub(time.Now()))
}

func Period(unit string, multi int64) time.Duration {
    switch unit {
    case "ms":
        return time.Duration(multi) * time.Millisecond
    case "second":
        return time.Duration(multi) * time.Second
    case "minute":
        return time.Duration(multi) * time.Minute
    case "hour":
        return time.Duration(multi) * time.Hour
    case "day":
        return time.Duration(multi*24) * time.Hour
    default:
        return 0
    }
}
