package atom

import (
    "fmt"
    "github.com/gobeam/stringy"
    "github.com/shopspring/decimal"
    "github.com/sotyou/logger"
    "net/url"
    "strconv"
    "time"
)

// CamelCase prefix_vvv -> PrefixVVV
func CamelCase(in string) string {
    s := stringy.New(in)
    return s.CamelCase("", "")
}

// SnakeCase AbcDef -> abc_def
func SnakeCase(in string) string {
    s := stringy.New(in)
    return s.SnakeCase("", "").ToLower()
}

// ParseUrl 解析一个URL字符串
func ParseUrl(rawUrl string) *url.URL {
    u, err := url.Parse(rawUrl)
    if err != nil {
        logger.Error("is not url, "+rawUrl, err)
    }
    return u
}

// GetHost 从一个Url中解析出Host
func GetHost(rawUrl string) string {
    u := ParseUrl(rawUrl)
    if u == nil {
        return ""
    }
    return u.Host
}

func IntToString(i int64) string {
    return fmt.Sprintf("%d", i)
}

func FloatToString(f float64) string { return strconv.FormatFloat(f, 'f', -1, 64) }

func FloatToRoundString(f float64, p int) string { return strconv.FormatFloat(f, 'f', p, 64) }

func UnixToTime(ts int64) time.Time {
    return time.Unix(ts, 0)
}

func InStrings(target string, list []string) bool {
    for _, v := range list {
        if v == target {
            return true
        }
    }
    return false
}

func DecimalToFloat64(d decimal.Decimal) float64 {
    f, _ := d.Float64()
    return f
}
