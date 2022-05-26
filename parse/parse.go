package parse

import (
	"fmt"
	"log"
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
	res := map[string]interface{}{}
	err := y.parse(yaml.String(), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (y *YamlParse) parse(config string, curMap map[string]interface{}) error {
	// 判断递归深度，使用前置缩进来判断
	deep := 0
	lines := strings.Split(config, "\n")
	valueLines := &strings.Builder{}
	for i := 0; i < len(lines); i++ {
		if y.identation(lines[i], deep) {
			if valueLines.Len() != 0 {
				k, v, err := y.realParse(valueLines.String())
				if err != nil {
					return err
				}
				err = y.dfsParse(k, v, deep, curMap)
				if err != nil {
					return err
				}
			}
			valueLines.Reset()
		}
		valueLines.WriteString(lines[i])
		valueLines.WriteByte('\n')
	}
	if valueLines.Len() != 0 {
		k, v, err := y.realParse(valueLines.String())
		if err != nil {
			return err
		}
		err = y.dfsParse(k, v, deep, curMap)
		if err != nil {
			return err
		}
	}
	return nil
}

//用于判断前置缩进与深度是否一致
func (y *YamlParse) identation(line string, deep int) bool {
	j := 0
	for i := 0; i < len(line); i++ {
		if line[i] == ' ' {
			j++
		} else {
			break
		}
	}
	return j == deep
}

//真正的解析逻辑，取第一个':' 做分隔，前者为key,后者为value
func (y *YamlParse) realParse(line string) (k, v string, err error) {
	kv := strings.SplitN(line, ":", 2)
	if len(kv) < 2 {
		return "", "", fmt.Errorf("yaml syntax error: %v", kv)
	}
	return strings.Trim(kv[0], " "), strings.Trim(kv[1], " "), nil
}

//深度搜索解析，用深度和cur的map一直递归
func (y *YamlParse) dfsParse(k string, v string, deep int, cur map[string]interface{}) error {
	log.Printf("k: %s, v: %s deep: %d\n", k, v, deep)
	if !strings.HasPrefix(v, "\n") {
		cur[k] = v
		return nil
	}
	cur[k] = map[string]interface{}{}
	// 处理当前嵌套的字符串, 开头的\n 给去掉然后继续判断
	nextKey, nextVal, err := y.realParse(v[1:])
	if !y.identation(nextKey, deep+2) {
		return fmt.Errorf("yaml syntax error, identation error, key: %s, value: %s", nextKey, nextVal)
	}
	if err != nil {
		return err
	}
	return y.dfsParse(nextKey, nextVal, deep+2, cur[k].(map[string]interface{}))
}
