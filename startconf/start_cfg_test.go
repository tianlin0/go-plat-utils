package startconf_test

import (
	"fmt"
	"github.com/tianlin0/go-plat-utils/conf/startupcfg"
	"github.com/tianlin0/go-plat-utils/conn"
	"github.com/tianlin0/go-plat-utils/conv"
	"github.com/tianlin0/go-plat-utils/crypto"
	"github.com/tianlin0/go-plat-utils/startconf"
	"testing"
)

// TemplateURL 模板访问地址
type TemplateURL struct {
	TemplateGetCdBiz            string //获取CD业务信息
	TemplateBindCdBizTag        string //绑定部署标签
	TemplateAppserverSubmitById string //gdp_appserver_go 提交template执行的方法
	GetPodListByCdId            string
	InsertCI                    string
}

// GdpConfig gdp全局配置
type GdpConfig struct {
	HostAndPort              *conn.Connect
	MysqlConnect             *startupcfg.MysqlConfig
	MysqlConnectODP          *conn.Connect
	RedisConnect             *conn.Connect
	GdpExternalOrigin        string
	ClientSecret             string
	TemplateIdBatchDeleteCd  string //批量删除部署的模板ID
	TemplateIdCopyCdWithCdId string
	TemplateIdCopyCd         string

	DefaultRTXLoginToken string //rtxLoginToken

	DefaultSystemRoleNameMap map[string][]string
}

func TestGetAllApiUrlMap(t *testing.T) {

	keyStr := "jasonsjiang29121"

	startupcfg.SetDecryptHandler(func(e startupcfg.Encrypted) (string, error) {
		str, err := crypto.CBCDecrypt(string(e), keyStr)
		if err != nil {
			return "", err
		}
		return str, nil
	})

	one, _ := startconf.NewStartupForYamlFile("dev.yaml")
	mapTemp := one.GetAllApiUrlMap()
	tempUrl := new(TemplateURL)
	conv.Unmarshal(mapTemp, tempUrl)

	fmt.Println(conv.String(tempUrl))

	cMap, _ := one.GetAllCustomMap()

	tempCMap := new(GdpConfig)
	conv.Unmarshal(cMap, tempCMap)

	fmt.Println(conv.String(tempCMap))

	ccTemp, err := one.GetAllMysqlMap()

	conv.Unmarshal(ccTemp, tempCMap)

	fmt.Println(tempCMap, err)

}
