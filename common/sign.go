package common

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "net/url"
    "sort"
    "strings"

    "github.com/gookit/goutil/strutil"
)

var signKey = "1234567890"

func CheckSign(sign string, signTime string, values url.Values) bool {
    params := parseValues(values)
    params["time"] = signTime

    return sign == getSign(params)
}

// 解析URl Values
func parseValues(values url.Values) map[string]interface{} {
    params := make(map[string]interface{})

    for k, v := range values {
        if len(v) == 1 {
            params[k] = v[0]
        } else if len(v) > 1 {
            params[k] = v
        }
    }

    return params
}

// 签名算法
func getSign(params map[string]interface{}) string {
    // 按key排序
    keys := make([]string, 0)
    for key, _ := range params {
        keys = append(keys, key)
    }
    sort.Strings(keys)

    var strs []string
    for _, key := range keys {
        if len(key) == 0 {
            continue
        }
        str, err := strutil.String(params[key])
        if err != nil {
            continue
        }
        strs = append(strs, fmt.Sprintf("%s=%s", key, str))
    }

    h := hmac.New(sha256.New, []byte(signKey))
    h.Write([]byte(strings.Join(strs, "&")))

    return hex.EncodeToString(h.Sum(nil))
}
