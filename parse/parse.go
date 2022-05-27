package parse

import (
	"fmt"
	"strings"
)

type Node struct {
	parent   *Node
	children []*Node
	rawStr   string
	deep     int
}
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
	res := map[string]interface{}{}
	err := y.parse(yaml.String(), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (y *YamlParse) parse(config string, curMap map[string]interface{}) error {
	deep := 0
	lines := strings.Split(config, "\n")
}

func (y *YamlParse) identation(line string) int {
	j := 0
	for i := 0; i < len(line); i++ {
		if line[i] == ' ' {
			j++
		} else {
			break
		}
	}
	return j
}

//真正的解析逻辑，取第一个':' 做分隔，前者为key,后者为value
func (y *YamlParse) realParse(line string) (k, v string, err error) {
	kv := strings.SplitN(line, ":", 2)
	if len(kv) < 2 {
		return "", "", fmt.Errorf("yaml syntax error: %v", kv)
	}
	return strings.Trim(kv[0], " "), strings.Trim(kv[1], " "), nil
}
