package startupcfg

import (
	"fmt"
	"github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)

// ConvertTo convert part config to target interface by json path
// @receiver api
// @param path json path to convert
// @param to target interface to convert
// @return error
func (c *ConfigAPI) ConvertTo(path string, to interface{}) error {
	if c.jsonBytes == nil {
		return fmt.Errorf("startup config nil")
	}
	from := gjson.GetBytes(c.jsonBytes, path).Raw
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.Unmarshal([]byte(from), &to); err != nil {
		return err
	}
	return nil
}

// ConvertFromCustomTo convert part of custom config to target interface by relative(relate of 'custom') json path
// @receiver api
// @param relativePath relative(relate of 'custom') json path to convert
// @param to target interface to convert
// @return error
//func (api *ConfigAPI) ConvertFromCustomTo(relativePath string, to interface{}) error {
//	if api.configBytes == nil {
//		return fmt.Errorf("startup config nil")
//	}
//	from := gjson.GetBytes(api.configBytes, fmt.Sprintf("custom.%s", relativePath)).Raw
//	var json = jsoniter.ConfigCompatibleWithStandardLibrary
//	if err := json.Unmarshal([]byte(from), &to); err != nil {
//		return err
//	}
//	return nil
//}

// ConvertFromCustomNormalTo convert part of custom normal config to target interface by relative(relate of 'custom.normal') json path
// @receiver api
// @param relativePath relative(relate of 'custom.normal') json path to convert
// @param to target interface to convert
// @return error
func (api *ConfigAPI) ConvertFromCustomNormalTo(relativePath string, to interface{}) error {
	if api.jsonBytes == nil {
		return fmt.Errorf("startup config nil")
	}
	from := gjson.GetBytes(api.jsonBytes, fmt.Sprintf("custom.normal.%s", relativePath)).Raw
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.Unmarshal([]byte(from), &to); err != nil {
		return err
	}
	return nil
}
