package source

import (
	"config_go/config"
	"config_go/parse"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Source interface {
	// 从数据源中读取所有的配置文件
	ReadFromSource() func() error
	ConvertBytesToMap() func() error
}

type YamlSource struct {
	// 读写锁,读不加锁，写需要加锁
	mu *sync.RWMutex
	// 保存从源读取的数据配置
	data map[string]interface{}
	// 数据源对应的parse
	parse  *parse.YamlParse
	config *config.SourceConfig
	// 数据源的原始配置信息
	raw []byte
	// 判断是否已经读取过source的数据
	done bool
}

// 实现读取操作, 初始化的时候读取
func (yaml *YamlSource) ReadFromSource() error {
	// 判断是否在本地，如果在本地则读取本地的yaml
	if yaml.done {
		return errors.New("read from source only can call one")
	}
	if yaml.config.Local() {
		bytes, err := ioutil.ReadFile(yaml.config.GetRealPath())
		if err != nil {
			return err
		}
		yaml.raw = bytes
		return nil
	}
	// 不在本地则调用接口返回应有的格式
	resp, err := http.Get(yaml.config.GetRealPath())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// 读取接口中所有的Yaml
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	yaml.raw = bytes
	return nil
}

func (yaml *YamlSource) ConvertBytesToMap() error {
	if len(yaml.raw) == 0 {
		log.Printf("rawConfig don't have something, please check your yaml")
		return nil
	}
	data, err := yaml.parse.Parse(yaml.raw)
	if err != nil {
		return err
	}
	yaml.data = data
	return nil
}
