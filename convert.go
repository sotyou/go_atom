package atom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/FireCoinJp/quant-core/core/logger"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// DecodeFromByte []byte -> struct
func DecodeFromByte(b []byte, v interface{}) error {
	buff := bytes.NewReader(b)
	return json.NewDecoder(buff).Decode(v)
}

// DecodeFromFile jsonFile -> struct
func DecodeFromFile(filename string, v interface{}) error {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.WithMessage(err, filename+" open error")
	}

	return json.NewDecoder(ByteToReader(f)).Decode(&v)
}

// DecodeFromReader json -> struct
func DecodeFromReader(reader io.Reader, v interface{}) error {
	if reader == nil {
		return errors.Errorf("DecodeFromReader() reader=nil")
	}
	return json.NewDecoder(reader).Decode(&v)
}

// 测试用的函数，输出内容
func ReaderDump(reader io.Reader) error {
	s := ReaderToString(reader)
	logger.Debug("dump", s)
	return nil
}

// DecodeFromReader string -> json
func DecodeFromString(str string, v interface{}) error {
	return json.NewDecoder(strings.NewReader(str)).Decode(&v)
}

// Encode struct -> json
func Encode(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", nil
	}
	return string(b), nil
}

// EncodeToBytes struct -> []byte
func EncodeToBytes(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func EncodeToReader(v interface{}) (io.Reader, error) {
	s, err := Encode(v)
	if err != nil {
		return nil, err
	}
	return strings.NewReader(s), nil
}

// io.Reader -> string
func ReaderToString(reader io.Reader) string {
	if reader == nil {
		return ""
	}
	buff := new(bytes.Buffer)
	_, _ = buff.ReadFrom(reader)
	return buff.String()
}

func ByteToReader(data []byte) io.Reader {
	return bytes.NewReader(data)
}

// url.Values -> json
func ValuesToJson(u url.Values) map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range u {
		if len(v) == 0 {
			continue
		}
		m[k] = v[0]
	}
	return m
}

// ======================= time =======================

func Time() string { return time.Now().Format("15:04:05") }
func DateTime() string { return time.Now().Format("2006-01-02 15:04:05") }
func Date() string { return time.Now().Format("2006-01-02") }
func HHMMSS2Time(s string) time.Time {
	t, err := time.Parse("15:04:05", s)
	if err != nil {
		panic(err)
	}
	return t
}
func Str2Time(s string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		panic(err)
	}
	return t
}

// ========= math =============
// utils
func StrToFloat(s string) float64 {
	if len(s) == 0 {
		return 0
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		logger.Error(fmt.Sprintf("can not convert float [%s]", s), err)
	}
	return f
}
func ObjToString(v interface{}) string {
	s, err := json.Marshal(v)
	if err != nil {
		logger.ErrorV("convert object to json error, "+err.Error(), v)
	}
	return string(s)
}
func AnyToString(v interface{}) string { return fmt.Sprintf("%#v", v) }
