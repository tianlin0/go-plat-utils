package startconf

import (
	startupCfg "github.com/tianlin0/go-plat-utils/conf/startupcfg"
	"github.com/tianlin0/go-plat-utils/conv"
	"github.com/tianlin0/go-plat-utils/templates"
)

type urlStruct struct {
	ServiceName   string
	ApiName       string
	serviceAPICfg *startupCfg.ServiceApiConfig
	ServiceAPIUrl string
}

// ServiceApiCfg 直接获取url全地址
func (u *urlStruct) ServiceApiCfg() *startupCfg.ServiceApiConfig {
	return u.serviceAPICfg
}

// StartCfg 初始化一个自定义配置
type StartCfg struct {
	Mysql           map[string]*startupCfg.MysqlConfig
	Redis           map[string]*startupCfg.RedisConfig
	Api             map[string]*startupCfg.ServiceApiConfig
	Custom          map[string]interface{}
	CustomSensitive map[string]string
}

func getInstanceFromYaml(isFilePathName bool, configFile string) (*startupCfg.ConfigAPI, error) {
	var conf *startupCfg.ConfigAPI
	var err error

	if isFilePathName {
		conf, err = startupCfg.NewByYamlFile(configFile)
	} else {
		conf, err = startupCfg.NewByYaml([]byte(configFile))
	}

	if err != nil {
		return nil, err
	}
	return conf, nil
}

// NewStartupForYamlFile 初始化一个yaml文件的Startup配置
func NewStartupForYamlFile(configFilePath string) (*StartCfg, error) {
	confTemp, err := getInstanceFromYaml(true, configFilePath)
	if err != nil {
		return nil, err
	}
	return commGetStartup(confTemp)
}

// NewStartupForYamlContent 初始化一个yaml内容的Startup配置
func NewStartupForYamlContent(configContent string) (*StartCfg, error) {
	confTemp, err := getInstanceFromYaml(false, configContent)
	if err != nil {
		return nil, err
	}
	return commGetStartup(confTemp)
}

func commGetStartup(sConf *startupCfg.ConfigAPI) (*StartCfg, error) {
	startTemp := new(StartCfg)
	startTemp.Mysql = sConf.MysqlAll()
	startTemp.Redis = sConf.RedisAll()
	startTemp.Api = sConf.ApiAll()
	startTemp.Custom = sConf.CustomNormalAll()
	startTemp.CustomSensitive = sConf.CustomSensitiveAll()
	return startTemp, nil
}

func (s *StartCfg) GetAllApiUrlMap() map[string]string {
	allApi := make(map[string]string)
	allUrlList := s.getAllApiUrl()
	for _, oneUrl := range allUrlList {
		allApi[oneUrl.ApiName] = oneUrl.ServiceAPIUrl
	}
	return allApi
}

func (s *StartCfg) GetAllCustomMap() (map[string]interface{}, error) {
	customMapNew := s.Custom
	if len(s.CustomSensitive) > 0 {
		customStr := conv.String(customMapNew)
		if customStr != "" {
			postData, err := templates.Template(customStr, s.CustomSensitive)
			if err != nil {
				return nil, err
			}
			customMapNew = make(map[string]interface{})
			err = conv.Unmarshal(postData, &customMapNew)
			if err != nil {
				return nil, err
			}
		}
	}
	return customMapNew, nil
}

func (s *StartCfg) GetAllMysqlMap() (map[string]*startupCfg.MysqlConfig, error) {
	customMapNew := make(map[string]*startupCfg.MysqlConfig)

	if s.Mysql != nil {
		for key, val := range s.Mysql {
			customMapNew[key] = val
		}
	}

	return customMapNew, nil
}

// getAllApiUrl 取得所有服务的地址
func (s *StartCfg) getAllApiUrl() []*urlStruct {
	allApi := s.Api
	urlApiList := make([]*urlStruct, 0)
	for sName, one := range allApi {
		for aName, url := range one.Urls {
			urlTemp := &urlStruct{
				ServiceName:   sName,
				serviceAPICfg: one,
				ApiName:       aName,
				ServiceAPIUrl: one.DomainName() + url,
			}
			urlApiList = append(urlApiList, urlTemp)
		}
	}
	return urlApiList
}