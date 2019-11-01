package stringtool

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

// GetRandomString 生成随机字符串
func GetRandomString(size int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GetMapByURL 通过传入的URL返回map[string]string 的格式
func GetMapByURL(value url.Values) (map[string]string, error) {
	m := make(map[string]string, 0)
	for k, v := range value {
		if len(v) >= 2 {
			return nil, fmt.Errorf("参数%v存在不确定性", k)
		}
		m[k] = v[0]
	}
	return m, nil
}


