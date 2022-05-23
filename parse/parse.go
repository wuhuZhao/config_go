package parse

type Parse interface {
	// 解析不同格式需要对应不同的实现
	Parse([]byte) map[string]interface{}
}

// yaml的解析版本
type YamlParse struct {
}

func (y *YamlParse) Parse([]byte) map[string]interface{} {
	return map[string]interface{}{}
}
