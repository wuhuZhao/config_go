package config

import "path"

type SourceConfig struct {
	local    bool
	ip       string
	port     string
	path     string
	fileName string
	fileType string
}

// 获取真正的路径
func (sourceConfig *SourceConfig) GetRealPath() string {
	if sourceConfig.local {
		return path.Join(sourceConfig.path, sourceConfig.fileName, ".", sourceConfig.fileType)
	}
	return path.Join(sourceConfig.ip, ":", sourceConfig.port, sourceConfig.path, sourceConfig.fileName, ".", sourceConfig.fileType)
}

// 返回是否为本地路径
func (sourceConfig *SourceConfig) Local() bool {
	return sourceConfig.local
}

func NewSourceConfig(local bool, ip, port, path, fileName, fileType string) *SourceConfig {
	return &SourceConfig{
		local:    local,
		ip:       ip,
		port:     port,
		path:     path,
		fileName: fileName,
		fileType: fileType,
	}
}
