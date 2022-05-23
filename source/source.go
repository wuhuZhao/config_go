package source

import (
	"config_go/config"
	"config_go/parse"
	"io/ioutil"
	"net/http"
	"sync"
)

type Source interface {
	// 从数据源中读取所有的配置文件
	ReadFromSource() func() ([]byte, error)
}

type YamlSource struct {
	// 读写锁,读不加锁，写需要加锁
	mu *sync.RWMutex
	// 保存从源读取的数据配置
	data map[string]interface{}
	// 数据源对应的parse
	parses map[string]*parse.Parse
	config *config.SourceConfig
}

// 实现读取操作
func (yaml *YamlSource) ReadFromSource() ([]byte, error) {
	// 判断是否在本地，如果在本地则读取本地的yaml
	if yaml.config.Local() {
		bytes, err := ioutil.ReadFile(yaml.config.GetRealPath())
		if err != nil {
			return nil, err
		}
		return bytes, nil
	}
	// 不在本地则调用接口返回应有的格式
	resp, err := http.Get(yaml.config.GetRealPath())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// 读取接口中所有的Yaml
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
