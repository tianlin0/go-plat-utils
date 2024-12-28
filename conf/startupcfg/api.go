package startupcfg

import (
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"sync"
	"time"
)

// ConfigAPI 配置访问实例
type ConfigAPI struct {
	runConf      *StartupConfig
	lock         sync.Mutex
	fromFileName string
	rowBytes     []byte //用来进行比较配置是否变动
	jsonBytes    []byte
}

// NewByYaml 通过配置内容创建实例
func NewByYaml(rawConfig []byte) (*ConfigAPI, error) {
	conf := &StartupConfig{}
	if err := yaml.Unmarshal(rawConfig, conf); err != nil {
		return nil, err
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsString, err := json.Marshal(conf)
	if err != nil {
		return nil, err
	}

	return &ConfigAPI{
		runConf:   conf,
		rowBytes:  rawConfig,
		jsonBytes: jsString,
	}, nil
}

// UpdateByYaml update config from file
func (c *ConfigAPI) UpdateByYaml(rawConfig []byte) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	newCfg, err := NewByYaml(rawConfig)
	if err != nil {
		return err
	}
	c.runConf = newCfg.runConf
	c.jsonBytes = newCfg.jsonBytes
	c.rowBytes = newCfg.rowBytes
	return nil
}

// NewByYamlFile 通过文件创建实例
func NewByYamlFile(fileName string) (*ConfigAPI, error) {
	configFile, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	c, err := NewByYaml(configFile)
	if err != nil {
		return nil, err
	}
	c.fromFileName = fileName
	return c, nil
}

// StartAutoUpdate poll for updates, default update per 60s
func (c *ConfigAPI) StartAutoUpdate(callback func(api *ConfigAPI) error, duration ...time.Duration) {
	dur := 60 * time.Second
	if len(duration) > 0 {
		dur = duration[0]
	}
	if c.fromFileName == "" {
		return
	}
	go func() {
		for {
			t := time.NewTimer(dur)
			<-t.C
			confTemp, err := NewByYamlFile(c.fromFileName)
			if err != nil {
				log.Println("Failed to startAutoUpdate for update config: ", err)
				continue
			}
			if string(confTemp.rowBytes) == string(c.rowBytes) {
				continue // 表示配置没有变化，不用自动更新
			}
			err = c.UpdateByYaml(confTemp.rowBytes)
			if err == nil {
				if err = callback(c); err != nil {
					log.Println("StartAutoUpdate Callback failed: ", err)
				}
			}
		}
	}()
}
