package param

import (
	"context"
	"fmt"
	"github.com/tianlin0/go-plat-utils/conv"
	"net/http"
	"net/url"
	"testing"
)

func TestToString(t *testing.T) {
	a := "a=1&b=2&a=7&a=9"
	q, err := url.ParseQuery(a)

	fmt.Println(conv.String(q), err)
}
func TestConvString(t *testing.T) {
	II := getIns()
	if oneParamMapTemp, ok := II.(map[string]interface{}); ok {
		for key, val := range oneParamMapTemp {
			fmt.Println(key, val)
		}
	}

	//na := conv.ChangeVariableName("aaa_bbb_ccc", "camel")
	//fmt.Println(na)
}

func getIns() interface{} {
	return map[string]interface{}{
		"aaa": 1,
	}
}

func TestContext(t *testing.T) {
	rawQuery := "pp=1&mm=2&pp=2"
	ret, err := url.ParseQuery(rawQuery)

	fmt.Println(ret, err)

	//ctx := context.Background()
	//setContext(ctx)
	//
	//fmt.Println(ctx.Value("aaaa"))

	//na := conv.ChangeVariableName("aaa_bbb_ccc", "camel")
	//fmt.Println(na)
}
func TestParam(t *testing.T) {
	req := new(http.Request)
	req.Method = http.MethodGet
	req.URL = new(url.URL)
	req.URL.RawQuery = "/v1/auth/auth-check?gpid=&exCluster=&paas_name=gdp-appserver-go"
	data := NewParam().GetAll(req)
	fmt.Println(data)
}
func TestPath(t *testing.T) {
	req := new(http.Request)
	req.Method = http.MethodGet
	req.URL = new(url.URL)
	req.URL.RawQuery = "/v1/auth/auth-check?gpid=&exCluster=&paas_name=gdp-appserver-go"
	data := NewParam().SetParsePathFunc(func(r *http.Request) map[string]string {
		return map[string]string{
			"name": "zhangsan",
		}
	}).GetAllQuery(req)
	fmt.Println(data)
}

type AAA struct {
	Name   string `json:"name" validate:"len=10"`
	Name2  string `json:"name2" validate:"required"`
	Number int    `json:"number" validate:"required"`
}

func TestParse(t *testing.T) {
	req := new(http.Request)
	req.Method = http.MethodGet
	req.URL = new(url.URL)
	req.URL.RawQuery = "/v1/auth/auth-check?gpid=&exCluster=&paas_name=gdp-appserver-go"
	aaa := new(AAA)
	err := NewParam().SetParsePathFunc(func(r *http.Request) map[string]string {
		return map[string]string{
			"name":   "1111111111",
			"name2":  "33333",
			"number": "0",
		}
	}).SetValidatorCustomErrorMessages(map[string]string{
		"aaaa":               "dddd",
		"AAA.Name2.required": "aaaaaajkjkjk",
	}).Parse(req, aaa)

	fmt.Println(aaa, err)
}

func setContext(ctx context.Context) context.Context {
	newCtx := context.WithValue(ctx, "aaaa", "bbbb")
	ctx = newCtx
	return newCtx
}
