package parse

import (
	"errors"
	"strings"
)

type Parse interface {
	// 解析不同格式需要对应不同的实现
	Parse([]byte) (map[string]interface{}, error)
}

// yaml的解析版本
type YamlParse struct {
}

func (y *YamlParse) Parse(raw []byte) (map[string]interface{}, error) {
	// 目前只支持简单的yaml文件解析
	// 按缩进实现解析的逻辑 直接解析int float map slice
	yaml := &strings.Builder{}
	yaml.Write(raw)
	// 目前先实现kv的格式， 后续再实现递归的格式
	lines := strings.Split(yaml.String(), "\r")
	res := map[string]interface{}{}
	for i := range lines {
		kv := strings.Split(lines[i], ":")
		if len(kv) != 2 {
			return nil, errors.New("yaml config file error syntax")
		}
		res[kv[0]] = kv[1]
	}
	return res, nil
}
