package atom

import (
    "crypto/hmac"
    "crypto/md5"
    "crypto/rand"
    "crypto/sha256"
    "encoding/base64"
    "encoding/hex"
    "fmt"
    logger "github.com/sotyou/go_logger"
)

// Hmac256
func Hmac256(base string, key string) string {
    h := hmac.New(sha256.New, []byte(key))
    h.Write([]byte(base))
    return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// MD5
func Md5(base string) string {
    w := md5.Sum([]byte(base))
    return fmt.Sprintf("%x", w)
}

func HmacHex(base string, key string) string {
    h := hmac.New(sha256.New, []byte(key))
    h.Write([]byte(base))
    return hex.EncodeToString(h.Sum(nil))
}

// RandomBytes
func RandomBytes(n int64) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    if err != nil {
        logger.Error("random bytes error", err)
        return nil, err
    }
    return b, nil
}

// 生成一个随机的字符串
func RandomString(n int64) string {
    const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
    bytes, _ := RandomBytes(n)
    for i, b := range bytes {
        bytes[i] = letters[b%byte(len(letters))]
    }
    return string(bytes)
}
